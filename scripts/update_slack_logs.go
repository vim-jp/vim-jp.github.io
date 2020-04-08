package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func main() {
	if err := doMain(); err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s\n", err)
		os.Exit(1)
	}
}

func doMain() error {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run scripts/update_slack_logs.go {config} {indir} {outdir}")
		fmt.Println("   ex: go run scripts/update_slack_logs.go scripts/update_slack_logs_config.json slacklog_data/ slacklog/")
		return nil
	}

	configFile := filepath.Clean(os.Args[1])
	inDir := filepath.Clean(os.Args[2])
	outDir := filepath.Clean(os.Args[3])

	cfg, err := readConfig(configFile)
	if err != nil {
		return fmt.Errorf("could not read config: %s", err)
	}
	userMap, err := readUsers(filepath.Join(inDir, "users.json"))
	if err != nil {
		return fmt.Errorf("could not read users.json: %s", err)
	}
	channels, err := readChannels(filepath.Join(inDir, "channels.json"))
	if err != nil {
		return fmt.Errorf("could not read channels.json: %s", err)
	}

	if err := mkdir(outDir); err != nil {
		return fmt.Errorf("could not create out directory: %s", err)
	}

	// Generate {outdir}/index.html (links to {channel})
	content, err := genIndex(channels)
	err = ioutil.WriteFile(filepath.Join(outDir, "index.html"), content, 0666)
	if err != nil {
		return fmt.Errorf("could not create %s/index.html: %s", outDir, err)
	}

	for i := range channels {
		if err := mkdir(filepath.Join(outDir, channels[i].Name)); err != nil {
			return fmt.Errorf("could not create %s/%s directory: %s", outDir, channels[i].Name, err)
		}
		msgMap, err := getMsgPerMonth(inDir, channels[i].Name)
		if err != nil {
			return err
		}
		// Generate {outdir}/{channel}/index.html (links to {channel}/{year}/{month})
		content, err := genChannelIndex(inDir, &channels[i], msgMap)
		if err != nil {
			return fmt.Errorf("could not generate %s/%s: %s", outDir, channels[i].Name, err)
		}
		err = ioutil.WriteFile(filepath.Join(outDir, channels[i].Name, "index.html"), content, 0666)
		if err != nil {
			return fmt.Errorf("could not create %s/%s/index.html: %s", outDir, channels[i].Name, err)
		}
		// Generate {outdir}/{channel}/{year}/{month}/index.html
		for _, msgPerMonth := range msgMap {
			if err := mkdir(filepath.Join(outDir, channels[i].Name, msgPerMonth.Year, msgPerMonth.Month)); err != nil {
				return fmt.Errorf("could not create %s/%s/%s/%s directory: %s", outDir, channels[i].Name, msgPerMonth.Year, msgPerMonth.Month, err)
			}
			content, err := genChannelPerMonthIndex(inDir, &channels[i], msgPerMonth, userMap, cfg)
			if err != nil {
				return fmt.Errorf("could not generate %s/%s/%s/%s/index.html: %s", outDir, channels[i].Name, msgPerMonth.Year, msgPerMonth.Month, err)
			}
			err = ioutil.WriteFile(filepath.Join(outDir, channels[i].Name, msgPerMonth.Year, msgPerMonth.Month, "index.html"), content, 0666)
			if err != nil {
				return fmt.Errorf("could not create %s/%s/index.html: %s", outDir, channels[i].Name, err)
			}
		}
	}
	return nil
}

func mkdir(path string) error {
	os.MkdirAll(path, 0777)
	if fi, err := os.Stat(path); os.IsNotExist(err) || !fi.IsDir() {
		return err
	}
	return nil
}

func genIndex(channels []channel) ([]byte, error) {
	params := make(map[string]interface{})
	params["channels"] = channels
	var out bytes.Buffer
	t, err := template.New("channelIndex").Delims("<<", ">>").Parse(`---
# vim:set ts=2 sts=2 sw=2 et:
layout: slacklog
title: vim-jp.slack.com log
---
<div>
<h2><a href='{{ post.url }}'>{{ page.title }}</a></h2>

<p>参加方法、各チャンネルの概要等は以下を参照して下さい。<br>
<a href='/docs/chat.html'>vim-jpのチャットルームについて</a></p>

<ul>
<<- range .channels >>
<li><a href='/slacklog/<< .Name >>/'>#<< .Name >></a></li>
<<- end >>
</ul>

</div>
`)
	if err != nil {
		return nil, err
	}
	err = t.Execute(&out, params)
	return out.Bytes(), err
}

func genChannelIndex(inDir string, channel *channel, msgMap map[string]*msgPerMonth) ([]byte, error) {
	params := make(map[string]interface{})
	params["channel"] = channel
	params["msgMap"] = msgMap
	var out bytes.Buffer
	t, err := template.New("channelIndex").Delims("<<", ">>").Parse(`---
# vim:set ts=2 sts=2 sw=2 et:
layout: slacklog
title: vim-jp.slack.com log - &#35<< .channel.Name >>
---
<div>
<h2><a href='/slacklog/'>vim-jp.slack.com log</a> - &#35<< .channel.Name >></h2>

<p>参加方法、各チャンネルの概要等は以下を参照して下さい。<br>
<a href='/docs/chat.html'>vim-jpのチャットルームについて</a></p>

<ul>
<<- range .msgMap >>
<li><a href='/slacklog/<< $.channel.Name >>/<< .Year >>/<< .Month >>/index.html'><< .Year >>年<< .Month >>月</a></li>
<<- end >>
</ul>

</div>
`)
	if err != nil {
		return nil, err
	}
	err = t.Execute(&out, params)
	return out.Bytes(), err
}

func genChannelPerMonthIndex(inDir string, channel *channel, msgPerMonth *msgPerMonth, userMap map[string]user, cfg *config) ([]byte, error) {
	params := make(map[string]interface{})
	params["channel"] = channel
	params["msgPerMonth"] = msgPerMonth
	var out bytes.Buffer
	// TODO check below subtypes work correctly
	// TODO support more subtypes
	t, err := template.New("channelPerMonthIndex").
		Delims("<<", ">>").
		Funcs(map[string]interface{}{
			"datetime": func(ts string) string {
				t := strings.Split(ts, ".")
				if len(t) != 2 {
					return ""
				}
				sec, err := strconv.ParseInt(t[0], 10, 64)
				if err != nil {
					return ""
				}
				nsec, err := strconv.ParseInt(t[0], 10, 64)
				if err != nil {
					return ""
				}
				japan, err := time.LoadLocation("Asia/Tokyo")
				if err != nil {
					return ""
				}
				return time.Unix(sec, nsec).In(japan).Format("2日 15:04:05")
			},
			"userIconUrl": func(userId string) string {
				user, ok := userMap[userId]
				if !ok {
					return "" // TODO show default icon
				}
				return user.Profile.Image48
			},
			"text": func() func(*message) string {
				var reLinkWithTitle = regexp.MustCompile(`&lt;(https?://[^>]+?\|(.+?))&gt;`)
				var reLink = regexp.MustCompile(`&lt;(https?://[^>]+?)&gt;`)
				var reCode = regexp.MustCompile("```([^`]+?)```")
				var reCodeShort = regexp.MustCompile("`([^`]+?)`")
				var reDel = regexp.MustCompile("~([^~]+?)~")
				var reNewline = regexp.MustCompile(`\n`)
				return func(msg *message) string {
					text := html.EscapeString(html.UnescapeString(msg.Text))
					text = reLinkWithTitle.ReplaceAllString(text, "<a href='${1}'>${2}</a>")
					text = reLink.ReplaceAllString(text, "<a href='${1}'>${1}</a>")
					text = reCode.ReplaceAllString(text, "<code>${1}</code>")
					text = reCodeShort.ReplaceAllString(text, "<code>${1}</code>")
					text = reDel.ReplaceAllString(text, "<del>${1}</del>")
					text = reNewline.ReplaceAllString(text, "<br>")
					if msg.Edited != nil {
						text += cfg.EditedSuffix
					}
					return text
				}
			}(),
		}).
		Parse(`---
# vim:set ts=2 sts=2 sw=2 et:
layout: slacklog
title: vim-jp.slack.com log - &#35<< .channel.Name >> - << .msgPerMonth.Year >>年<< .msgPerMonth.Month >>月
---
<div>
<h2><a href='/slacklog/'>vim-jp.slack.com log</a> - <a href='/slacklog/<< .channel.Name >>/'>&#35<< .channel.Name >></a> - << .msgPerMonth.Year >>年<< .msgPerMonth.Month >>月</h2>

{% raw %}
<<- range .msgPerMonth.Messages >>
<<- if eq .Subtype "" >>
  <span class='slacklog-message' id='<< .Ts >>'>
  <img class='slacklog-icon' src='<< userIconUrl .User >>'>
  <span class='slacklog-name'><< or .UserProfile.DisplayName .UserProfile.RealName >></span>
  <a class='slacklog-datetime' href='#<< .Ts >>'><< datetime .Ts >></a>
  <span class='slacklog-text'><< text . >></span>
    <<- if .Attachments >>
    <span class='slacklog-attachments'>
      <<- range .Attachments >>
        <<- if eq .ServiceName "GitHub" >>
          <span class='slacklog-attachment slacklog-attachment-github'>
            <span class='slacklog-attachment-github-serviceicon'><img src='<< .ServiceIcon >>'></span>
            <span class='slacklog-attachment-github-servicename'><< html .ServiceName >></span>
            <span class='slacklog-attachment-github-title'><a href='<< .TitleLink >>'><< html .Title >></a></span>
            <span class='slacklog-attachment-github-text'><< html .Text >></span>
          </span>
        <<- else if eq .ServiceName "twitter" >>
          <span class='slacklog-attachment slacklog-attachment-twitter'>
            <span class='slacklog-attachment-twitter-authoricon'><img src='<< .AuthorIcon >>'></span>
            <span class='slacklog-attachment-twitter-authorname'><< .AuthorName >></span>
            <span class='slacklog-attachment-twitter-authorsubname'><< .AuthorSubname >></span>
            <span class='slacklog-attachment-twitter-text'><< html .Text >></span>
            <span class='slacklog-attachment-twitter-footericon'><img src='<< .FooterIcon >>'></span>
            <span class='slacklog-attachment-twitter-footer'><< html .Footer >></span>
            <<- if .VideoHtml >>
            <span class='slacklog-attachment-twitter-video'><< .VideoHtml >></span>
            <<- end >>
          </span>
        <<- else if or .Title .Text >>
          <span class='slacklog-attachment slacklog-attachment-other'>
            <<- if and .ServiceIcon .ServiceName >>
            <div>
              <span class='slacklog-attachment-other-serviceicon'><img src='<< .ServiceIcon >>'></span>
              <span class='slacklog-attachment-other-servicename'><< html .ServiceName >></span>
            </div>
            <<- end >>
            <<- if and .Title .TitleLink >>
            <div class='slacklog-attachment-other-title'><a href='<< .TitleLink >>'><< html .Title >></a></div>
            <<- else if .Title >>
            <div class='slacklog-attachment-other-title'><< html .Title >></div>
            <<- end >>
            <<- if .Text >>
            <div class='slacklog-attachment-other-text'><< html .Text >></div>
            <<- end >>
          </span>
        <<- end >>
      <<- end >>
    </span>
    <<- end >>
  </span>
<<- else if eq .Subtype "bot_message" >>
<< .Username >> << datetime .Ts >>: << html .Text >>
<<- end >>
<<- end >>
{% endraw %}

</div>
`)
	if err != nil {
		return nil, err
	}
	err = t.Execute(&out, params)
	return out.Bytes(), err
}

type msgPerMonth struct {
	Year     string
	Month    string
	Messages []message
}

// "{year}-{month}-{day}.json"
var reMsgFilename = regexp.MustCompile(`^(\d{4})-(\d{2})-\d{2}\.json$`)

func getMsgPerMonth(inDir string, channelName string) (map[string]*msgPerMonth, error) {
	dir, err := os.Open(filepath.Join(inDir, channelName))
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	result := make(map[string]*msgPerMonth)
	for i := range names {
		m := reMsgFilename.FindStringSubmatch(names[i])
		if len(m) == 0 {
			fmt.Fprintf(os.Stderr, "[warning] skipping %s/%s/%s ...", inDir, channelName, names[i])
			continue
		}
		key := m[1] + m[2]
		if _, exists := result[key]; !exists {
			result[key] = &msgPerMonth{Year: m[1], Month: m[2]}
		}
		msgs, err := readMessages(filepath.Join(inDir, channelName, names[i]))
		if err != nil {
			return nil, err
		}
		result[key].Messages = append(result[key].Messages, msgs...)
	}
	for key := range result {
		sort.SliceStable(result[key].Messages, func(i, j int) bool {
			// must be the same digits, so no need to convert the timestamp to a number
			return result[key].Messages[i].Ts < result[key].Messages[j].Ts
		})
	}
	return result, nil
}

type message struct {
	ClientMsgId string              `json:"client_msg_id"`
	Typ         string              `json:"type"`
	Subtype     string              `json:"subtype"`
	Text        string              `json:"text"`
	User        string              `json:"user"`
	Ts          string              `json:"ts"`
	Username    string              `json:"username"`
	Team        string              `json:"team"`
	UserTeam    string              `json:"user_team"`
	SourceTeam  string              `json:"source_team"`
	UserProfile messageUserProfile  `json:"user_profile"`
	Attachments []messageAttachment `json:"attachments,omitempty"`
	// Blocks      []messageBlock     `json:"blocks,omitempty"`    // TODO
	Reactions []messageReaction `json:"reactions,omitempty"`
	Edited    *messageEdited    `json:"edited"`
}

type messageEdited struct {
	User string `json:"user"`
	Ts   string `json:"ts"`
}

type messageUserProfile struct {
	AvatarHash        string `json:"avatar_hash"`
	Image72           string `json:"image72"`
	FirstName         string `json:"first_name"`
	RealName          string `json:"real_name"`
	DisplayName       string `json:"display_name"`
	Team              string `json:"team"`
	Name              string `json:"name"`
	IsRestricted      bool   `json:"is_restricted"`
	IsUltraRestricted bool   `json:"is_ultra_restricted"`
}

type messageBlock struct {
	Typ      string                `json:"type"`
	Elements []messageBlockElement `json:"elements,omitempty"`
}

type messageBlockElement struct {
	Typ       string `json:"type"`
	Name      string `json:"name"`       // for type = "emoji"
	Text      string `json:"text"`       // for type = "text"
	ChannelId string `json:"channel_id"` // for type = "channel"
}

type messageAttachment struct {
	ServiceName     string `json:"service_name"`
	AuthorIcon      string `json:"author_icon"`
	AuthorName      string `json:"author_name"`
	AuthorSubname   string `json:"author_subname"`
	Title           string `json:"title"`
	TitleLink       string `json:"title_link"`
	Text            string `json:"text"`
	Fallback        string `json:"fallback"`
	ThumbUrl        string `json:"thumb_url"`
	FromUrl         string `json:"from_url"`
	ThumbWidth      int    `json:"thumb_width"`
	ThumbHeight     int    `json:"thumb_height"`
	ServiceIcon     string `json:"service_icon"`
	Id              int    `json:"id"`
	OriginalUrl     string `json:"original_url"`
	VideoHtml       string `json:"video_html"`
	VideoHtmlWidth  int    `json:"video_html_width"`
	VideoHtmlHeight int    `json:"video_html_height"`
	Footer          string `json:"footer"`
	FooterIcon      string `json:"footer_icon"`
}

type messageReaction struct {
	Name  string   `json:"name"`
	Users []string `json:"users,omitempty"`
	Count int      `json:"count"`
}

func readMessages(msgJsonPath string) ([]message, error) {
	content, err := ioutil.ReadFile(msgJsonPath)
	if err != nil {
		return nil, err
	}
	var msgs []message
	err = json.Unmarshal(content, &msgs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s: %s", msgJsonPath, err)
	}
	return msgs, nil
}

type config struct {
	EditedSuffix string `json:"edited_suffix"`
}

func readConfig(configPath string) (*config, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg config
	err = json.Unmarshal(content, &cfg)
	return &cfg, err
}

type user struct {
	Id                string      `json:"id"`
	TeamId            string      `json:"team_id"`
	Name              string      `json:"name"`
	Deleted           bool        `json:"deleted"`
	Color             string      `json:"color"`
	RealName          string      `json:"real_name"`
	Tz                string      `json:"tz"`
	TzLabel           string      `json:"tz_label"`
	TzOffset          int         `json:"tz_offset"` // tzOffset / 60 / 60 = [-+] hour
	Profile           userProfile `json:"profile"`
	IsAdmin           bool        `json:"is_admin"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
	IsBot             bool        `json:"is_bot"`
	IsAppUser         bool        `json:"is_app_user"`
	Updated           int64       `json:"updated"`
}

type userProfile struct {
	Title                 string      `json:"title"`
	Phone                 string      `json:"phone"`
	Skype                 string      `json:"skype"`
	RealName              string      `json:"real_name"`
	RealNameNormalized    string      `json:"real_name_normalized"`
	DisplayName           string      `json:"display_name"`
	DisplayNameNormalized string      `json:"display_name_normalized"`
	Fields                interface{} `json:"fields"` // TODO ???
	StatusText            string      `json:"status_text"`
	StatusEmoji           string      `json:"status_emoji"`
	StatusExpiration      int64       `json:"status_expiration"`
	AvatarHash            string      `json:"avatar_hash"`
	FirstName             string      `json:"first_name"`
	LastName              string      `json:"last_name"`
	Image24               string      `json:"image_24"`
	Image32               string      `json:"image_32"`
	Image48               string      `json:"image_48"`
	Image72               string      `json:"image_72"`
	Image192              string      `json:"image_192"`
	Image512              string      `json:"image_512"`
	StatusTextCanonical   string      `json:"status_text_canonical"`
	Team                  string      `json:"team"`
}

func readUsers(usersJsonPath string) (map[string]user, error) {
	content, err := ioutil.ReadFile(usersJsonPath)
	if err != nil {
		return nil, err
	}
	var users []user
	err = json.Unmarshal(content, &users)
	result := make(map[string]user, len(users))
	for i := range users {
		result[users[i].Id] = users[i]
	}
	return result, err
}

type channel struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Created    int64          `json:"created"`
	Creator    string         `json:"creator"`
	IsArchived bool           `json:"is_archived"`
	IsGeneral  bool           `json:"is_general"`
	Members    []string       `json:"members,omitempty"`
	Pins       []channelPin   `json:"pins,omitempty"`
	Topic      channelTopic   `json:"topic"`
	Purpose    channelPurpose `json:"purpose"`
}

type channelPin struct {
	Id      string `json:"id"`
	Typ     string `json:"type"`
	Created int64  `json:"created"`
	User    string `json:"user"`
	Owner   string `json:"owner"`
}

type channelTopic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int64  `json:"last_set"`
}

type channelPurpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int64  `json:"last_set"`
}

func readChannels(channelsJsonPath string) ([]channel, error) {
	content, err := ioutil.ReadFile(channelsJsonPath)
	if err != nil {
		return nil, err
	}
	var channels []channel
	err = json.Unmarshal(content, &channels)
	sort.Slice(channels, func(i, j int) bool {
		return channels[i].Name < channels[j].Name
	})
	return channels, err
}

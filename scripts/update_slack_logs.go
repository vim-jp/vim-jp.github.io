package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run update_slack_logs _data/slack_logs/")
	}
	// slackLogDir := filepath.Join(filepath.Dir(os.Args[0]), "..", "_data", "slack_logs")
	slackLogDir := os.Args[1]
	users, err := readUsers(filepath.Join(slackLogDir, "users.json"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s\n", err)
	}
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

func readUsers(usersJsonPath string) ([]user, error) {
	content, err := ioutil.ReadFile(usersJsonPath)
	if err != nil {
		return nil, err
	}
	var users []user
	err = json.Unmarshal(content, &users)
	return users, err
}

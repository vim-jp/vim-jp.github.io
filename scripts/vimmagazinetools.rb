#!/usr/bin/ruby1.9
# coding: utf-8


require 'net/https'
require 'uri'
require 'json'
require 'optparse'
require 'date'


VIMPATCH_URL = "http://ftp.vim.org/pub/vim/patches/7.3/%s"
VIMPATCH_README_URL = "http://ftp.vim.org/pub/vim/patches/7.3/README"
VIMSCRIPT_URL = "http://www.vim.org/scripts/script.php?script_id=%s"
VIMSCRIPT_LIST_URL = "http://www.vim.org/scripts/script_search_results.php?&show_me=99999"
VIMHG_COMMIT_URL = "http://code.google.com/p/vim/source/detail?r=%s"
VIMHG_TAGS_URL = "http://vim.googlecode.com/hg/.hgtags"
GITHUBAPI_ISSUE_LIST_URL = "https://api.github.com/repos/%s/%s/issues"
GITHUBAPI_ISSUE_URL = "https://api.github.com/repos/%s/%s/issues/%s"


def count_if(seq, &cond)
  i = 0
  for e in seq
    if yield e
      i = i + 1
    end
  end
  return i
end


def httpget(url)
  if url.start_with?("https")
    uri = URI.parse(url)
    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl = true
    request = Net::HTTP::Get.new(uri.request_uri)
    response = http.request(request)
  else
    response = Net::HTTP.get_response(URI.parse(url))
  end
  return response
end


def htmlentityunescape(text)
  # TODO: add entry when you found a entity in the output.
  html_entities_name2codepoint = {
    "quot" => "\"",
    "amp" => "&",
    "apos" => "'",
    "lt" => "<",
    "gt" => ">",
    "nbsp" => " ",
  }
  # TODO: web browser accept entity without semicolon.  pass it for now.
  text.gsub(/&#?\w+;/) {|text|
    if text[0...2] == "&#"
      # character reference
      if text[0...3] == "&#x"
        text[3...-1].to_i(16).chr("UTF-8")
      else
        text[2...-1].to_i().chr("UTF-8")
      end
    else
      html_entities_name2codepoint[text[1...-1]] || text
    end
  }
end


def mdescape(s)
  s = s.gsub("\\", "\\\\")
  s = s.gsub("<", "\\<")
  s = s.gsub("[", "\\[")
  s = s.gsub("]", "\\]")
  return s
end


def vimpatch_all()
  items = []
  tag2rev = vimhg_tags()
  readme = httpget(VIMPATCH_README_URL).entity.force_encoding("UTF-8")
  for line in readme.split(/\r\n|\r|\n/)
    m = line.match(/^\s*(?#size)(\d+)  (?#version)(\d\.\d\.\d{3})  (?#summary)(.*)$/)
    if !m
      next
    end
    e = {}
    e["size"] = m[1].to_i
    e["version"] = m[2]
    e["summary"] = m[3]
    # e["url"] = sprintf(VIMPATCH_URL, e["version"])
    major, minor, patchlevel = e["version"].split(".")
    tag = "v#{major}-#{minor}-#{patchlevel}"
    e["tag"] = tag
    if tag2rev.has_key?(tag)
      e["revision"] = tag2rev[tag]
      e["url"] = sprintf(VIMHG_COMMIT_URL, e["revision"])
    else
      # tag is missing
      e["revision"] = ""
      e["url"] = ""
    end
    items << e
  end
  items.sort_by!{|e| e["version"]}
  return items
end


def vimhg_tags()
  hgtags = httpget(VIMHG_TAGS_URL).entity.force_encoding("UTF-8")
  tags = {}
  for line in hgtags.split(/\r\n|\r|\n/)
    rev, tag = line.split()
    tags[tag] = rev
  end
  return tags
end


def vimscript_all()
  data = httpget(VIMSCRIPT_LIST_URL).entity.force_encoding("ISO-8859-1").encode("UTF-8")
  items = []
  i = 0
  e = {}
  for line in data.split(/\r\n|\r|\n/)
    line = line.strip()
    if line !~ /rowodd|roweven/
      next
    end
    if i == 0
      e["script_id"] = line.match(/script_id=(\d+)/)[1]
      e["url"] = sprintf(VIMSCRIPT_URL, e["script_id"])
      e["name"] = htmlentityunescape(line.gsub(/<[^>]*>/, ""))
    elsif i == 1
      e["rating"] = htmlentityunescape(line.gsub(/<[^>]*>/, ""))
    elsif i == 2
      e["rating"] = line.gsub(/<[^>]*>/, "").to_i()
    elsif i == 3
      e["downloads"] = line.gsub(/<[^>]*>/, "").to_i()
    elsif i == 4
      e["summary"] = htmlentityunescape(line.gsub(/<[^>]*>/, ""))
    end
    i = i + 1
    if i == 5
      items << e
      i = 0
      e = {}
    end
  end
  items.sort_by!{|e| e["script_id"].to_i()}
  return items
end


def parse_linkheader(s)
  res = []
  for link in s.split(',')
    e = {}
    m = link.match(/<([^>]*)>/)
    e["url"] = m[1]
    link.match(/(\w+)="([^"]*)"/) {|m|
      e[m[1].downcase()] = m[2]
    }
    res << e
  end
  return res
end


def githubissue_getallpages(nexturl)
  items = []
  while nexturl != nil
    r = httpget(nexturl)
    pageitems = JSON.load(r.entity.force_encoding("UTF-8"))
    items = items + pageitems
    nexturl = nil
    if r.key?("link")
      for link in parse_linkheader(r["link"])
        if link["rel"] == "next"
          nexturl = link["url"]
          break
        end
      end
    end
  end
  return items
end


def githubissue_all(user, repo)
  url = sprintf(GITHUBAPI_ISSUE_LIST_URL, user, repo) + "?per_page=9999"
  items = []
  items = items + githubissue_getallpages(url + "&state=open")
  items = items + githubissue_getallpages(url + "&state=closed")
  items.sort_by!{|e| e["number"]}
  return items
end


def scriptranking(oldstate, curstate)
  ranking = curstate.dup()
  olddownloads = {}
  for e in oldstate
    olddownloads[e["script_id"]] = e["downloads"]
  end
  for e in ranking
    e["downloads_diff"] = e["downloads"] - (olddownloads[e["script_id"]] || 0)
  end
  ranking.sort_by!{|e| -e["script_id"].to_i}
  ranking.sort_by!{|e| -e["downloads_diff"]}
  items = []
  i = 1
  for e in ranking
    summary = "#{e["name"]} : #{e["summary"]}"
    items << "#{i}. [#{mdescape(summary)}][#{e["url"]}] (#{e["downloads_diff"]})"
    i = i + 1
  end
  return items
end


def cmd_patchlist(_args)
  for e in vimpatch_all()
    summary = "#{e["version"]} : #{e["summary"]}"
    puts "- [#{mdescape(summary)}](#{e["url"]})"
  end
end


def cmd_scriptlist(_args)
  for e in vimscript_all()
    summary = "#{e["name"]} : #{e["summary"]}"
    puts "- [#{mdescape(summary)}](#{e["url"]})"
  end
end


def cmd_githubissuelist(args)
  for e in githubissue_all(args["user"], args["repo"])
    summary = "Issue ##{e["number"]} : #{e["title"]}"
    puts "- [#{mdescape(summary)}](#{e["html_url"]})"
  end
end


def cmd_scriptjson(_args)
  puts JSON.dump(vimscript_all())
end


def cmd_scriptranking(args)
  oldstate = JSON.load(open(args["oldfile"]))
  curstate = JSON.load(open(args["curfile"]))
  for e in scriptranking(oldstate, curstate)
    puts e
  end
end


# state["updated"] = "%Y-%m-%d"
# state["vim"]["version"] = "X.Y.ZZZ"
# state["script"]["script_id"] = "script_id"
# state["script"]["state"] = [...]
# state["vim-jp/issues"]["opencount"] = 0
# state["vim-jp/issues"]["closedcount"] = 0
# state["vim-jp/issues"]["number"] = 0
def cmd_generate(args)
  state = JSON.load(open(args["statefile"]))

  updated = DateTime.strptime(state["updated"], "%Y-%m-%d")
  vimpatches = vimpatch_all()
  vimscripts = vimscript_all()
  ranking = scriptranking(state["script"]["state"], vimscripts)
  vimjpissues = githubissue_all("vim-jp", "issues")
  opencount = count_if(vimjpissues) {|x| x["state"] == "open"}
  opendiff = opencount - state["vim-jp/issues"]["opencount"]
  closedcount = count_if(vimjpissues) {|x| x["state"] == "closed"}
  closeddiff = closedcount - state["vim-jp/issues"]["closedcount"]

  puts "## リリース情報"
  puts ""
  for e in vimpatches
    if e["version"] <= state["vim"]["version"]
      next
    end
    summary = "#{e["version"]} : #{e["summary"]}"
    puts "- [#{mdescape(summary)}](#{e["url"]})"
  end
  puts ""

  puts "## 新着スクリプト"
  puts ""
  for e in vimscripts
    if e["script_id"].to_i <= state["script"]["script_id"].to_i
      next
    end
    summary = "#{e["name"]} : #{e["summary"]}"
    puts "- [#{mdescape(summary)}](#{e["url"]})"
  end
  puts ""

  puts "## 月間ダウンロードランキング"
  puts ""
  for e in ranking[0...10]
    puts e
  end
  puts ""

  puts "## vim-jp issues"
  puts ""
  puts sprintf("Open : %d (%+d) | Closed : %d (%+d)", opencount, opendiff, closedcount, closeddiff)
  puts ""
  for e in vimjpissues
    if e["number"] <= state["vim-jp/issues"]["number"]
      next
    end
    summary = "Issue ##{e["number"]} : #{e["title"]}"
    puts "- [#{mdescape(summary)}](#{e["html_url"]})"
  end
  puts ""

  newstate = {
    "updated" => DateTime.now().strftime("%Y-%m-%d"),
    "vim" => {
      "version" => vimpatches[-1]["version"],
    },
    "script" => {
      "script_id" => vimscripts[-1]["script_id"],
      "state" => vimscripts,
    },
    "vim-jp/issues" => {
      "opencount" => opencount,
      "closedcount" => closedcount,
      "number" => vimjpissues.dup().map{|x| x["number"]}.max(),
    },
  }

  if args["update"]
    open(args["statefile"], "w") {|f|
      JSON.dump(newstate, f)
    }
  end
end


def main()
  cmd = ARGV.shift()
  args = {}
  if cmd == nil
    puts "vimmagazinetools.rb patchlist|scriptlist|githubissuelist|scriptjson|scriptranking|generate"
  elsif cmd == "patchlist"
    cmd_patchlist(args)
  elsif cmd == "scriptlist"
    cmd_scriptlist(args)
  elsif cmd == "githubissuelist"
    args["user"] = ARGV.shift()
    args["repo"] = ARGV.shift()
    cmd_githubissuelist(args)
  elsif cmd == "scriptjson"
    cmd_scriptjson(nil)
  elsif cmd == "scriptranking"
    args["oldfile"] = ARGV.shift()
    args["curfile"] = ARGV.shift()
    cmd_scriptranking(args)
  elsif cmd == "generate"
    opt = OptionParser.new
    args['update'] = false
    opt.on('--update') {|v| args['update'] = true}
    opt.parse!(ARGV)
    args["statefile"] = ARGV.shift()
    cmd_generate(args)
  else
    puts "Error: no such command"
  end
end


main()

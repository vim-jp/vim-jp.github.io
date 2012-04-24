#!/usr/bin/env python3


import io
import re
import urllib.request
import urllib.error
import json
import html.entities
import argparse
import datetime


VIMPATCH_URL = "ftp://ftp.vim.org/pub/vim/patches/7.3/{version}"
VIMPATCH_README_URL = "ftp://ftp.vim.org/pub/vim/patches/7.3/README"
VIMSCRIPT_URL = "http://www.vim.org/scripts/script.php?script_id={script_id}"
VIMSCRIPT_LIST_URL = "http://www.vim.org/scripts/script_search_results.php?&show_me=99999"
VIMHG_COMMIT_URL = "http://code.google.com/p/vim/source/detail?r={r}"
VIMHG_TAGS_URL = "http://vim.googlecode.com/hg/.hgtags"
GITHUBAPI_ISSUE_LIST_URL = "https://api.github.com/repos/{user}/{repo}/issues"
GITHUBAPI_ISSUE_URL = "https://api.github.com/repos/{user}/{repo}/issues/{number}"


def count_if(seq, cond):
    return len([x for x in seq if cond(x)])


def httpget(url):
    return urllib.request.urlopen(url)


# http://effbot.org/zone/re-sub.htm#unescape-html
##
# Removes HTML or XML character references and entities from a text string.
#
# @param text The HTML (or XML) source text.
# @return The plain text, as a Unicode string, if necessary.
def htmlentityunescape(text):
    def fixup(m):
        text = m.group(0)
        if text[:2] == "&#":
            # character reference
            try:
                if text[:3] == "&#x":
                    return chr(int(text[3:-1], 16))
                else:
                    return chr(int(text[2:-1]))
            except ValueError:
                pass
        else:
            # named entity
            try:
                text = chr(html.entities.name2codepoint[text[1:-1]])
            except KeyError:
                pass
        return text  # leave as is
    return re.sub("&#?\w+;", fixup, text)


def mdescape(s):
    s = s.replace("\\", "\\\\")
    s = s.replace("<", "\\<")
    s = s.replace("[", "\\[")
    s = s.replace("]", "\\]")
    return s


def vimpatch_all():
    items = []
    tag2rev = vimhg_tags()
    url = VIMPATCH_README_URL
    with httpget(url) as r:
        readme = r.read().decode("utf-8")
    for line in io.StringIO(readme):
        line = line.rstrip("\r\n")
        m = re.match(r"^\s*(?P<size>\d+)  (?P<version>\d\.\d\.\d{3})  (?P<summary>.*)$", line)
        if not m:
            continue
        e = {}
        e["size"] = int(m.group("size"))
        e["version"] = m.group("version")
        e["summary"] = m.group("summary")
        # e["url"] = VIMPATCH_URL.format(version=e["version"])
        major, minor, patchlevel = e["version"].split(".")
        tag = "v{}-{}-{}".format(major, minor, patchlevel)
        e["tag"] = tag
        if tag in tag2rev:
            e["revision"] = tag2rev[tag]
            e["url"] = VIMHG_COMMIT_URL.format(r=e["revision"])
        else:
            # tag is missing
            e["revision"] = ""
            e["url"] = ""
        items.append(e)
    items.sort(key=lambda e: e['version'])
    return items


def vimhg_tags():
    url = VIMHG_TAGS_URL
    with httpget(url) as r:
        hgtags = r.read().decode("utf-8")
    tags = {}
    for line in io.StringIO(hgtags):
        line = line.rstrip("\r\n")
        rev, tag = line.split()
        tags[tag] = rev
    return tags


def vimscript_all():
    with httpget(VIMSCRIPT_LIST_URL) as r:
        data = r.read().decode("ISO-8859-1")
    items = []
    i = 0
    e = {}
    for line in io.StringIO(data):
        line = line.strip()
        m = re.search(r"rowodd|roweven", line)
        if not m:
            continue
        if i == 0:
            e["script_id"] = re.search(r"script_id=(\d+)", line).group(1)
            e["url"] = VIMSCRIPT_URL.format(script_id=e["script_id"])
            e["name"] = htmlentityunescape(re.sub(r"<[^>]*>", "", line))
        elif i == 1:
            e["type"] = htmlentityunescape(re.sub(r"<[^>]*>", "", line))
        elif i == 2:
            e["rating"] = int(re.sub(r"<[^>]*>", "", line))
        elif i == 3:
            e["downloads"] = int(re.sub(r"<[^>]*>", "", line))
        elif i == 4:
            e["summary"] = htmlentityunescape(re.sub(r"<[^>]*>", "", line))
        i = i + 1
        if i == 5:
            items.append(e)
            i = 0
            e = {}
    items.sort(key=lambda e: int(e["script_id"]))
    return items


def parse_linkheader(s):
    res = []
    for link in s.split(','):
        e = {}
        m = re.search(r"<([^>]*)>", link)
        e["url"] = m.group(1)
        for m in re.finditer(r'(\w+)="([^"]*)"', link):
            e[m.group(1).lower()] = m.group(2)
        res.append(e)
    return res


def githubissue_getallpages(nexturl):
    items = []
    while nexturl is not None:
        with httpget(nexturl) as r:
            pageitems = json.loads(r.read().decode("utf-8"))
            items.extend(pageitems)
            nexturl = None
            if "link" in r.info():
                for link in parse_linkheader(r.info()["link"]):
                    if link["rel"] == "next":
                        nexturl = link["url"]
                        break
    return items


def githubissue_all(user, repo):
    url = GITHUBAPI_ISSUE_LIST_URL.format(user=user, repo=repo) \
            + "?per_page=9999"
    items = []
    items.extend(githubissue_getallpages(url + "&state=open"))
    items.extend(githubissue_getallpages(url + "&state=closed"))
    items.sort(key=lambda e: e["number"])
    return items


def scriptranking(oldstate, curstate):
    ranking = curstate[:]
    olddownloads = {}
    for e in oldstate:
        olddownloads[e["script_id"]] = e["downloads"]
    for e in ranking:
        e["downloads_diff"] = \
                e["downloads"] - olddownloads.get(e["script_id"], 0)
    ranking.sort(key=lambda e: int(e['script_id']), reverse=True)
    ranking.sort(key=lambda e: e['downloads_diff'], reverse=True)
    for i, e in enumerate(ranking, start=1):
        summary = "{} : {}".format(e["name"], e["summary"])
        yield "{}. [{}]({}) ({})".format(
                i, mdescape(summary), e["url"], e['downloads_diff'])


def cmd_patchlist(_args):
    for e in vimpatch_all():
        summary = "{} : {}".format(e["version"], e["summary"])
        print("- [{}]({})".format(mdescape(summary), e["url"]))


def cmd_scriptlist(_args):
    for e in vimscript_all():
        summary = "{} : {}".format(e["name"], e["summary"])
        print("- [{}]({})".format(mdescape(summary), e["url"]))


def cmd_githubissuelist(args):
    for e in githubissue_all(args.user, args.repo):
        summary = "Issue #{} : {}".format(e["number"], e["title"])
        print("- [{}]({})".format(mdescape(summary), e["html_url"]))


def cmd_scriptjson(_args):
    print(json.dumps(vimscript_all()))


def cmd_scriptranking(args):
    with open(args.oldfile) as f:
        oldstate = json.load(f)
    with open(args.curfile) as f:
        curstate = json.load(f)
    for e in scriptranking(oldstate, curstate):
        print(e)


# state["updated"] = "%Y-%m-%d"
# state["vim"]["version"] = "X.Y.ZZZ"
# state["script"]["script_id"] = "script_id"
# state["script"]["state"] = [...]
# state["vim-jp/issues"]["opencount"] = 0
# state["vim-jp/issues"]["closedcount"] = 0
# state["vim-jp/issues"]["number"] = 0
def cmd_generate(args):
    with open(args.statefile) as f:
        state = json.load(f)

    updated = datetime.datetime.strptime(state["updated"], "%Y-%m-%d")
    vimpatches = vimpatch_all()
    vimscripts = vimscript_all()
    ranking = list(scriptranking(state["script"]["state"], vimscripts))
    vimjpissues = githubissue_all("vim-jp", "issues")
    opencount = count_if(vimjpissues, lambda x: x["state"] == "open")
    opendiff = opencount - state["vim-jp/issues"]["opencount"]
    closedcount = count_if(vimjpissues, lambda x: x["state"] == "closed")
    closeddiff = closedcount - state["vim-jp/issues"]["closedcount"]

    print("## リリース情報")
    print("")
    for e in vimpatches:
        if e["version"] <= state["vim"]["version"]:
            continue
        summary = "{} : {}".format(e["version"], e["summary"])
        print("- [{}]({})".format(mdescape(summary), e["url"]))
    print("")

    print("## 新着スクリプト")
    print("")
    for e in vimscripts:
        if int(e["script_id"]) <= int(state["script"]["script_id"]):
            continue
        summary = "{} : {}".format(e["name"], e["summary"])
        print("- [{}]({})".format(mdescape(summary), e["url"]))
    print("")

    print("## 月間ダウンロードランキング")
    print("")
    for e in ranking[:10]:
        print(e)
    print("")

    print("## vim-jp issues")
    print("")
    print("Open : {} ({:+}) | Closed : {} ({:+})".format(
            opencount, opendiff, closedcount, closeddiff))
    print("")
    for e in vimjpissues:
        if e["number"] <= state["vim-jp/issues"]["number"]:
            continue
        summary = "Issue #{} : {}".format(e["number"], e["title"])
        print("- [{}]({})".format(mdescape(summary), e["html_url"]))
    print("")

    newstate = {
        "updated": datetime.datetime.now().strftime("%Y-%m-%d"),
        "vim": {
            "version": vimpatches[-1]["version"],
        },
        "script": {
            "script_id": vimscripts[-1]["script_id"],
            "state": vimscripts,
        },
        "vim-jp/issues": {
            "opencount": opencount,
            "closedcount": closedcount,
            "number": max(x["number"] for x in vimjpissues),
        },
    }

    if args.update:
        with open(args.statefile, "w") as f:
            json.dump(newstate, f)


def main():
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers()

    parser_patchlist = subparsers.add_parser('patchlist')
    parser_patchlist.set_defaults(func=cmd_patchlist)

    parser_scriptlist = subparsers.add_parser('scriptlist')
    parser_scriptlist.set_defaults(func=cmd_scriptlist)

    parser_githubissuelist = subparsers.add_parser('githubissuelist')
    parser_githubissuelist.add_argument('user')
    parser_githubissuelist.add_argument('repo')
    parser_githubissuelist.set_defaults(func=cmd_githubissuelist)

    parser_scriptjson = subparsers.add_parser('scriptjson')
    parser_scriptjson.set_defaults(func=cmd_scriptjson)

    parser_scriptranking = subparsers.add_parser('scriptranking')
    parser_scriptranking.add_argument('oldfile')
    parser_scriptranking.add_argument('curfile')
    parser_scriptranking.set_defaults(func=cmd_scriptranking)

    parser_generate = subparsers.add_parser('generate')
    parser_generate.add_argument('statefile')
    parser_generate.add_argument('--update', action='store_true')
    parser_generate.set_defaults(func=cmd_generate)

    args = parser.parse_args()
    args.func(args)


if __name__ == "__main__":
    main()

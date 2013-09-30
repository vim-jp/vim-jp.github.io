#!/usr/bin/env python3
# I WANT YOU FOR VIM
# Find non-member from issue and add them to vim-jp team.


import base64
import getpass
import json
import re
import urllib.error
import urllib.request


class GithubApi:
    def __init__(self, username=None, password=None):
        self.username = username
        self.password = password

    def http_request(self, url, method, data=None):
        req = urllib.request.Request(url, data)
        req.get_method = lambda: method
        if self.username is not None and self.password is not None:
            seed = self.username + ":" + self.password
            token = base64.b64encode(seed.encode("utf-8")).decode("utf-8")
            req.add_header("Authorization", "Basic " + token)
        return urllib.request.urlopen(req)

    def http_get(self, url):
        return self.http_request(url, "GET")

    def http_put(self, url):
        return self.http_request(url, "PUT", b"")

    def parse_linkheader(self, s):
        res = []
        for link in s.split(','):
            e = {}
            m = re.search(r"<([^>]*)>", link)
            e["url"] = m.group(1)
            for m in re.finditer(r'(\w+)="([^"]*)"', link):
                e[m.group(1).lower()] = m.group(2)
            res.append(e)
        return res

    def fetch_allpages(self, url, limit=None):
        items = []
        nexturl = url
        count = 0
        while nexturl is not None:
            with self.http_get(nexturl) as res:
                data = json.loads(res.read().decode("utf-8"))
                items.extend(data)
                count = count + 1
                if limit is not None and count >= limit:
                    break
                nexturl = None
                if "link" in res.info():
                    for link in self.parse_linkheader(res.info()["link"]):
                        if link["rel"] == "next":
                            nexturl = link["url"]
                            break
        return items

    def fetch_one(self, url):
        with self.http_get(url) as res:
            data = json.loads(res.read().decode("utf-8"))
            return data
        return data

    def issue_list(self, user, repo):
        url = "https://api.github.com/repos/{user}/{repo}/issues" \
                .format(user=user, repo=repo) \
                + "?per_page=9999"
        items = []
        items.extend(self.fetch_allpages(url + "&state=open"))
        items.extend(self.fetch_allpages(url + "&state=closed"))
        items.sort(key=lambda e: e["number"])
        return items

    def issue_get(self, user, repo, number):
        url = "https://api.github.com/repos/{user}/{repo}/issues/{number}" \
                .format(user=user, repo=repo, number=number)
        return self.fetch_one(url)

    def issue_comment_list(self, user, repo, number):
        url = "https://api.github.com/repos/{user}/{repo}/issues/{number}/comments" \
                .format(user=user, repo=repo, number=number)
        return self.fetch_allpages(url)

    def issue_comment_get(self, user, repo, id_):
        url = "https://api.github.com/repos/{user}/{repo}/issues/comments/{id}" \
                .format(user=user, repo=repo, id=id_)
        return self.fetch_one(url)

    def org_member_list(self, org):
        url = "https://api.github.com/orgs/{org}/members" \
                .format(org=org)
        return self.fetch_allpages(url)

    def org_member_get(self, org, user):
        url = "https://api.github.com/orgs/{org}/members/{user}" \
                .format(org=org, user=user)
        return self.fetch_one(url)

    def org_public_member_list(self, org):
        url = "https://api.github.com/orgs/{org}/public_members" \
                .format(org=org)
        return self.fetch_allpages(url)

    def org_public_member_get(self, org, user):
        url = "https://api.github.com/orgs/{org}/public_members/{user}" \
                .format(org=org, user=user)
        return self.fetch_one(url)

    def org_team_list(self, org):
        url = "https://api.github.com/orgs/{org}/teams" \
                .format(org=org)
        return self.fetch_allpages(url)

    def team_get(self, id_):
        url = "https://api.github.com/teams/{id}" \
                .format(id=id_)
        return self.fetch_one(url)

    def team_member_list(self, id_):
        url = "https://api.github.com/teams/{id}/members" \
                .format(id=id_)
        return self.fetch_allpages(url)

    def team_member_get(self, id_, user):
        url = "https://api.github.com/teams/{id}/members/{user}" \
                .format(id=id_, user=user)
        return self.fetch_one(url)

    def team_member_add(self, id_, user):
        url = "https://api.github.com/teams/{id}/members/{user}" \
                .format(id=id_, user=user)
        return self.http_put(url)

    def org_event_list(self, org):
        url = "https://api.github.com/orgs/{org}/events" \
                .format(org=org)
        return self.fetch_allpages(url)


class NotFound(BaseException):
    pass


def find(seq, cond):
    for e in seq:
        if cond(e):
            return e
    raise NotFound()


def main():
    blacklist = []

    username = input("username: ")
    password = getpass.getpass("password: ")

    github = GithubApi(username, password)

    teams = github.org_team_list("vim-jp")

    team_writer = find(teams, lambda x: x["name"] == "writer")
    team_writer_members = github.team_member_list(team_writer["id"])
    team_writer_members_id = [x["id"] for x in team_writer_members]

    team_vimdoc_ja = find(teams, lambda x: x["name"] == "vimdoc-ja")
    team_vimdoc_ja_members = github.team_member_list(team_vimdoc_ja["id"])
    team_vimdoc_ja_members_id = [x["id"] for x in team_vimdoc_ja_members]

    todo = []

    for event in github.org_event_list("vim-jp"):
        if event["type"] == "IssuesEvent":
            if event["payload"]["action"] == "opened":
                todo.append([event["repo"], event["actor"]])
        elif event["type"] == "IssueCommentEvent":
            todo.append([event["repo"], event["actor"]])

    for repo, user in todo:
        if user["id"] in blacklist:
            continue
        if user["id"] not in team_writer_members_id:
            print("add writer: {}".format(user["login"]))
            github.team_member_add(team_writer["id"], user["login"])
            team_writer_members_id.append(user["id"])
        if user["id"] not in team_vimdoc_ja_members_id and repo["name"] == "vimdoc-ja":
            print("add vimdoc_ja: {}".format(user["login"]))
            github.team_member_add(team_vimdoc_ja["id"], user["login"])
            team_vimdoc_ja_members_id.append(user["id"])


if __name__ == "__main__":
    main()

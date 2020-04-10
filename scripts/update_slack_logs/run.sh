#!/bin/bash

cd "$(dirname "$0")/../.." || exit "$?"
go run scripts/update_slack_logs/update_slack_logs.go scripts/update_slack_logs/ slacklog_data/ slacklog/

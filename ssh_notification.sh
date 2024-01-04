#!/bin/bash

LOGGED_USER="$(whoami)"
LOGGED_HOST="$(hostname -f)"

LOGGED_TTY="$(ps -p "$$" -o tname h)"
LOGGED_IP="$(w -i | grep -w "$LOGGED_TTY" | awk '{print $3}')"

sudo /usr/local/bin/sshtgbot/ssh_tg_bot \
    -config_path="/usr/local/bin/sshtgbot/config.json" \
    -user="$LOGGED_USER" \
    -host="$LOGGED_HOST" \
    -ip="$LOGGED_IP"

exit 0
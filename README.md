# Ssh Telegram Bot
## Simple Telegram bot for ssh login notifications

## Message Example
```text
Host: example.com
User Name: user
Date Time UTC: 2024-01-04T17:38:09Z
Ip: 127.0.0.1
```

## File config.json
### Path `/usr/local/bin/sshtgbot/config.json`
| Key | Type | Description |
| - | - | - |
| userList | Array[int64] | array contains id's of users to whom the bot will send messages ||
| tgBotApiToken | string | token of the Telegram bot ||
| tgBotDebug | Boolean | will output debugging information ||

## How to use
```bash
wget https://github.com/A1ex3/ssh_tg_bot/releases/download/main/ssh_tg_bot_build.tar.gz
```
```bash
make setup
```

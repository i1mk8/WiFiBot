#!/bin/bash

WORK_DIR=/etc/WiFiBot
mkdir $WORK_DIR

read -p "Введите токен бота: " BOT_TOKEN
read -p "Введите user id (если вы хотите добавить несколько аккаунтов, то введите их user id через запятую): " BOT_USERS
read -p "Введите временную зону (список - https://github.com/Lewington-pitsos/golang-time-locations): " TIMEZONE

echo "{\"bot_token\": \"${BOT_TOKEN}\", \"bot_users\": ${BOT_USERS}, \"timezone\": \"${TIMEZONE}\", \"schedule_enabled\": false, \"schedule_down_hour\": 0, \"schedule_down_minute\": 0, \"schedule_up_hour\": 0, \"schedule_up_minute\": 0}" > "${WORK_DIR}/wifi_bot.json"

wget -O "${WORK_DIR}/WiFiBot" https://github.com/i1mk8/WiFiBot/releases/latest/download/WiFiBot

mkdir /etc/ssl
wget -O /etc/ssl/cert.pem https://curl.se/ca/cacert-2023-12-12.pem

echo "@reboot ${WORK_DIR}/WiFiBot" > "${WORK_DIR}/cron"
crontab "${WORK_DIR}/cron"

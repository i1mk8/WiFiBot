#!/bin/bash

WORK_DIR=/etc/WiFiBot
mkdir $WORK_DIR

read -p "Введите токен бота: " BOT_TOKEN
read -p "Введите user id (если вы хотите добавить несколько аккаунтов, то введите их user id через запятую): " BOT_USERS
read -p "Введите временную зону (список - https://github.com/Lewington-pitsos/golang-time-locations): " TIMEZONE

echo "{\"bot_token\": \"${BOT_TOKEN}\", \"bot_users\": [${BOT_USERS}], \"timezone\": \"${TIMEZONE}\", \"schedule_enabled\": false, \"schedule_down_hour\": 0, \"schedule_down_minute\": 0, \"schedule_up_hour\": 0, \"schedule_up_minute\": 0}" > "${WORK_DIR}/wifi_bot.json"

EXECUTABLE="${WORK_DIR}/on-startup.sh"
wget -O $EXECUTABLE https://raw.githubusercontent.com/i1mk8/WiFiBot/master/on-startup.sh
chmod +x $EXECUTABLE

CRON="${WORK_DIR}/cron"
echo "@reboot ${EXECUTABLE}" > $CRON
crontab $CRON

fs save
reboot

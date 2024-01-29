#!/bin/bash

WORK_DIR=/etc/WiFiBot
mkdir $WORK_DIR

read -p "Введите токен бота: " BOT_TOKEN
read -p "Введите user id (если вы хотите добавить несколько аккаунтов, то введите их user id через запятую): " BOT_USERS
read -p "Введите временную зону (список - https://github.com/Lewington-pitsos/golang-time-locations): " TIMEZONE

echo "{\"bot_token\": \"${BOT_TOKEN}\", \"bot_users\": [${BOT_USERS}], \"timezone\": \"${TIMEZONE}\", \"schedule_enabled\": false, \"schedule_down_hour\": 0, \"schedule_down_minute\": 0, \"schedule_up_hour\": 0, \"schedule_up_minute\": 0}" > "${WORK_DIR}/wifi_bot.json"

EXECUTABLE="/etc/scripts/wan_up.sh"
wget -O - https://raw.githubusercontent.com/i1mk8/WiFiBot/master/usersrvc.sh >> $EXECUTABLE
chown daemon:daemon $EXECUTABLE
chmod +x $EXECUTABLE

fs save
reboot

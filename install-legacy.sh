#!/bin/bash

# Установщик для ПО 7.2.x и старее

read -p "Введите токен бота: " BOT_TOKEN
read -p "Введите user id (если вы хотите добавить несколько аккаунтов, то введите их user id через запятую): " BOT_USERS

echo "{\"bot_token\": \"${BOT_TOKEN}\", \"bot_users\": [${BOT_USERS}], \"schedule_enabled\": false, \"schedule_down_hour\": 0, \"schedule_down_minute\": 0, \"schedule_up_hour\": 0, \"schedule_up_minute\": 0}" > "/etc/wifi_bot.json"

EXECUTABLE="/etc/scripts/wan_up.sh"
wget -O - https://raw.githubusercontent.com/i1mk8/WiFiBot/master/usersrvc.sh >> $EXECUTABLE
chown daemon:daemon $EXECUTABLE
chmod +x $EXECUTABLE

fs save
reboot

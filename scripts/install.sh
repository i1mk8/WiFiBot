#!/bin/bash

WORK_DIR="/etc/WiFiBot"
mkdir $WORK_DIR

read -p "Введите токен бота: " BOT_TOKEN
read -p "Введите user id (если вы хотите добавить несколько аккаунтов, то введите их user id через запятую): " BOT_USERS

echo "{\"bot_token\": \"${BOT_TOKEN}\", \"bot_users\": [${BOT_USERS}], \"schedule_enabled\": false, \"schedule_down_hour\": 0, \"schedule_down_minute\": 0, \"schedule_up_hour\": 0, \"schedule_up_minute\": 0}" > "${WORK_DIR}/config.json"

EXECUTABLE="${WORK_DIR}/startup.sh"
wget -O - https://raw.githubusercontent.com/i1mk8/WiFiBot/master/scripts/startup.sh > $EXECUTABLE
chmod +x $EXECUTABLE

CRON="/tmp/cron"
crontab -l > $CRON
echo "@reboot sleep 60; ${EXECUTABLE}" >> $CRON # Задержка необходима для полной инициализации системы и корректного запуска
crontab $CRON

fs save
reboot

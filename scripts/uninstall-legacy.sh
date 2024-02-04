#!/bin/bash

# Скрипт удаления для ПО 7.2.x и старее
rm /etc/wifi_bot.json
rm /etc/scripts/wan_up.sh

fs save
reboot

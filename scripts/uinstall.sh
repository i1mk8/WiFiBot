#!/bin/bash

# Скрипт удаления для ПО 7.5.30 и новее
rm /etc/wifi_bot.json
rm /etc/scripts/usersrvc.sh

fs save
reboot

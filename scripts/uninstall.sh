#!/bin/bash

CRON="/tmp/cron"
crontab -l > $CRON
sed -e "s~@reboot /etc/WiFiBot/startup.sh~~g" -i $CRON
crontab $CRON

rm -rf "/etc/WiFiBot"

fs save
reboot

#!/bin/bash

# К сожалению, исполняемый файл бота слишком большой и не влезает в ПЗУ, поэтому приходится его скачивать в ОЗУ при каждой перезагрузке системы

WORK_DIR=/etc/WiFiBot
cd $WORK_DIR

EXECUTABLE="${WORK_DIR}/WiFiBot"
wget -O $EXECUTABLE https://github.com/i1mk8/WiFiBot/releases/download/v1.0/WiFiBot
chmod +x $EXECUTABLE
$EXECUTABLE

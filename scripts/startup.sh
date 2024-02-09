#!/bin/bash

# К сожалению, исполняемый файл бота слишком большой и не влезает в ПЗУ, поэтому приходится его скачивать в ОЗУ при каждой перезагрузке системы

EXECUTABLE="/tmp/WiFiBot"
wget -O $EXECUTABLE https://github.com/i1mk8/WiFiBot/releases/download/v1.2/WiFiBot
chmod +x $EXECUTABLE
$EXECUTABLE

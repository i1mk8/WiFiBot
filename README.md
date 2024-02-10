# WiFi Bot

Бот в телеграм для управления беспроводным режимом роутера Wi-CAT Альфин. Бот умеет включать/выключать Wi-FI как по команде, так и по расписанию.

## Функционал
- Возможность включения/выключения беспроводного режима
- Возможность включения/выключения расписания
- Изменение расписания (изменение времени)
- Просмотр расписания

## Скриншоты
![Главное меню](./screenshots/bot_main_menu.png)
![Меню управления расписанием](./screenshots/bot_schedule_menu.png)

## Установка
Включаем crontab: Админ панель > Сервисы > Разное > Планировщик
```
$ ssh Admin@192.168.0.1
# cd /tmp
# wget -O install.sh https://raw.githubusercontent.com/i1mk8/WiFiBot/master/scripts/install.sh
# chmod +x install.sh
# ./install.sh
```

## Удаление
```
$ ssh Admin@192.168.0.1
# cd /tmp
# wget -O uninstall.sh https://raw.githubusercontent.com/i1mk8/WiFiBot/master/scripts/uninstall.sh
# chmod +x uninstall.sh
# ./uninstall.sh
```
Выключаем crontab: Админ панель > Сервисы > Разное > Планировщик

Примечание: Выключение crontab не обязательно, скрипт автоматически удалит себя из crontab. Если crontab больше нигде не испольузется, то можно его отключить.

## Сборка
Перед сборкой необходимо изменить [os_linux.go](https://github.com/golang/go/blob/8b23b7b04234424791e26b8d2d26f61ef1311a9f/src/runtime/os_linux.go#L532), добавив `&& sig != 128` в условие на 532 строке.
```
$ go env -w GOOS=linux
$ go env -w GOARCH=mipsle
$ go build -ldflags "-s -w" ./main.go
```
В теории бот должен работать на любых роутерах, главное указать соответствующие архитектуру и интерфейсы в [WiFiManager.go](./src/WiFiManager/WiFiManager.go)

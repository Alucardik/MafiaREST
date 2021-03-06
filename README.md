# Чекин Игорь, Сервис-ориентированные архитектуры, Домашнее задание 6

## Общая структура проекта

В папке `MafiaREST` содержатся все основные файлы приложения. Обозначим назначения пакетов внутри модуля (название пакетов совпадают с названиями папок).
* `config` - здесь задаются константы, используемые остальными пакетами модуля
* `db` - имплементирует интерфейс для взаимодействия с базой данных (в моем случае с `MongoDB`)
* `endpoints` - содержит обработчики различных ресурсов сервера
* `msgbroker` - имплементирует интерфейс взаимодействия с очередью сообщений
* `pdfgen` - создает PDF-отчет
* `schemes` - содержит схемы данных для базы данных
* `server` - непосредственно  RestFul http-сервер, обрабатывающий запросы пользователей
* `utils` - общие вспомогательные функции
* `worker` - воркер, генерирующий отчеты в соответствии с паттерном "очередь задач"


В качестве базы данных используется NoSQL-решение `MongoDB`, в связи с простотой интерфейса и нативностью сериализации данных из `json` в Mongo и обратно. Очередь задач реализуется через брокер сообщений `RabbitMQ`.

Для локального запуска сервера необходимо наличие `go` версии `1.17 - 1.18`. Требуется перейти в папку `MafiaREST` и ввести команду `go run .`
Локальный запуск воркера: `go run . --mode=worker`. Используется один Докер-Образ для обоих вариантов работы и регулируется значением переменной окружения `MafiaREST_MODE`. По умолчанию она выставляется в значение `server`, можно заменить его на `worker`, тогда будет запущен воркер. [Ссылка на образ](https://hub.docker.com/layers/soa-images/alucardik/soa-images/MafiaREST/images/sha256-3eedf991a93496601bce41eb5d14c83a9c94ee3b3609ab027532f9c2e2ab9b4a?context=explore) (`alucardik/soa-images:MafiaREST`)

> ВНИМАНИЕ: для работы и сервера, и воркера нужны поднятые сервисы брокера сообщений и базы данных, это может усложнять локальный запуск 

Тем не менее, настоятельно рекомендуется запускать все сервисы через `docker-compose`, так как между ними уже настроены все связи. Достаточно выполнить следующие команды из корневой папки:

```bash
docker-compose build
docker up --scale server=1 --scale worker=n
```

Где `n` - число воркеров, которые будут подключены к очереди. Далее можно отправлять запросы на `localhost:8080/`.

## Описание форматов данных

`User` - представляет профиль игрока. Представляется следующим `JSON`:
```
{
    "name": {имя пользователя}, // строка
    "avatar": {ссылка на аватар}, // строка
    "sex": {пол} // число: 0 - мужской, 1 - женский
    "email": {e-mail} // строка
}
```

`UserStats` - представляет статистику игр некоторого игрока. Представляется следующим `JSON`:
```
{
	"uid": {ID игрока} // тип ID
	"session_count": {общее количество сессий} // число
	"wins": {количество побед} // число          
    "losses": {количество поражений} // число
	"total_time": {общее время в игре} // число, секунды
}
```

`SessionReport` - представляет отчет об игровой сессии. Представляется следующим `JSON`:
```
{
    "outcome": {исход сессии}, // число: 0 - поражение, 1 - победа
    "duration": {длительность сессии}, // число: количество секунд
}
```

> ВНИМАНИЕ: на сервере производится валидация схем, поэтому для успешной отправки запроса нужно включать все указанные поля и указывать в них корректные значения (например, пустое имя или неправильная схема e-mail приведут к провалу запроса)

## Описание API сервиса

Сервис поддерживает следующие эндпоинты:

1. `GET /users` - получить список профилей всех пользователей с их ID (остальные методы работают уже с конкретным ID)
> ID выводятся для возможности взаимодействия с сервисом, не имея прямого доступа к БД, при отсутсвии регистрации и авторизации 
2. `GET /users/{uid}` - получить профиль игрока с ID, соответствующим uid
3. `POST /users` - добавить нового игрока. В тело запроса необходимо приложить `JSON` по схеме `User` (также автоматически создается `UserStats` для данного игрока)
4. `PATCH /users/{uid}` - изменить данные профиля игрока с ID, соответствующим uid. В тело запроса необходимо приложить `JSON` по схеме `User`
5. `DELETE /users/{uid}` - удалить игрока с ID, соответствующим uid
6. `PUT /stats/{uid}` - обновить статистику игрока с ID, соответствующим uid. В тело запроса необходимо приложить `JSON` по схеме `SessionReport`
7. `GET /stats/{uid}` - запросить генерацию PDF-отчета для игрока с ID, соответствующим uid. В ответе на запрос будет содержаться ссылка для скачивания ответа (ей удобнее воспользоваться через браузер, чтобы скачать файл отчета без лишних сложностей). Если на момент перехода по ссылке отчет будет не готов - пользователя уведомят об этом в ответном сообщении


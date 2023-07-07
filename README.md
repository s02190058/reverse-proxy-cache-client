# Reverse Proxy Cache Client

Простой клиент для [прокси-сервиса](https://github.com/s02190058/reverse-proxy-cache).
Принимает ссылки на YouTube видео вида `https://www.youtube.com/watch?v=<video_id>`
(если точнее, строки удовлетворяющие регулярному выражению `^(https?://)?www\.youtube\.com/watch\?v=[a-zA-Z0-9_-]{11}`). 
## Быстрый старт

Устанавливает утилиту `rpc-cli` в директорию `$GOPATH/bin`:

```shell
go install github.com/s02190058/reverse-proxy-cache-client/cmd/rpc-cli@latest
```

## Опции

```
--host  - хост, на котором запущен сервис (localhsot по умолчанию)
--port  - порт, на котором запущен сервис (9090 по умолчанию)
--dir   - директорию, в которую будут загружаться изображения (рабочая директория по умолчанию)
--async - булев флаг, позволяющий обрабатывать каждый запрос в отдельной goroutine (false по умолчанию)
```

Запуск в асинхронном режиме:

```shell
rpc-cli --async=true 2>logs.txt
```
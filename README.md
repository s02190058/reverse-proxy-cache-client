# Reverse Proxy Cache Client

простой клиент для [прокси-сервиса](https://github.com/s02190058/reverse-proxy-cache).

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
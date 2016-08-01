# go-telegram-weather

Simple bot for telegram, which give weather by your location.

You need Telegram bot token an openweather token.

# Telegram bot token

To get token just follow [instruction](https://core.telegram.org/bots/#3-how-do-i-create-a-bot).

# Weather API token

To get token just follow [instruction](http://openweathermap.org/appid)

# Build & start

1. Write token to `env.sh`

```
    echo "export BOT_TOKEN=<token>" >> env.sh
    echo "export WEATHER_TOKEN=<token>" >> env.sh
```

2. Build add
```
    make deps
    make
```

3. Export envs

```
    source env.sh
```

4. Run app

```
    ./bin/weather-bot
```
FROM golang:1.12.5-alpine3.9

RUN apk --update add docker --no-cache

ENV APP_DIR=/app
WORKDIR $APP_DIR
COPY . $APP_DIR

RUN go build -o bin/ojibot src/main.go

CMD echo "$CRON_ENTRY" | crontab - && crond -f

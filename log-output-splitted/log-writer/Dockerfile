FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go build .


FROM scratch

WORKDIR /app

COPY --from=build /app/log-writer .

CMD ["./log-writer"]
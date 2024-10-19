FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go build .


FROM alpine

WORKDIR /app

COPY --from=build /app/broadcaster .

CMD ["./broadcaster"]
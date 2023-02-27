FROM golang:1.20.1 as build-env

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 go build -o issues-commenter .

FROM alpine:3.15

COPY --from=build-env /app /

CMD ["/issues-commenter"]

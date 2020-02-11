FROM golang:1.13.7-alpine3.11 as build
WORKDIR /builddir

EXPOSE 8080

ADD main.go .
RUN go build -o server main.go

FROM alpine:3.11

RUN adduser -D server
USER server

COPY --from=build /builddir/server .

CMD ./server

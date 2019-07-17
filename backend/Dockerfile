FROM golang:1.12.0 as builder
WORKDIR /go/src/main
ENV GO111MODULE=on
COPY . .
WORKDIR /go/src/main/cmd/server
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /go/bin/api

FROM alpine:latest AS final
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk add --no-cache tzdata
COPY --from=builder /go/bin/api /go/bin/api

ENTRYPOINT [ "/go/bin/api"]

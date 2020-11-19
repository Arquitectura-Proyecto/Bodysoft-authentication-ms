FROM golang

ENV GO111MODULE=on

WORKDIR /go/src/github.com/jpbmdev/BODYSOFT-AUTHENTICATION-MS

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 4002
ENTRYPOINT ["/go/src/github.com/jpbmdev/BODYSOFT-AUTHENTICATION-MS/Bodysoft-authentication-ms"]
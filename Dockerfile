FROM golang:alpine

ADD bin/syborg-linux-x64 /syborg-linux-x64
RUN chmod 755 /syborg-linux-x64

RUN apk add ffmpeg youtube-dl

ENTRYPOINT /syborg-linux-x64
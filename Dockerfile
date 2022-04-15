FROM golang:bullseye

ADD bin/syborg /syborg

RUN apt-get install -y python ffmpeg youtube-dl

CMD ['/syborg']
FROM ubuntu:16.04

ADD gopath/assets /opt/app/assets
ADD linuxamd64/bin/server /opt/app/bin/server

WORKDIR /opt/app
ENTRYPOINT /opt/app/bin/server
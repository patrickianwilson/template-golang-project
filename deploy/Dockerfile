FROM ubuntu:16.04

ADD assets /opt/app/assets
ADD templates /opt/app/templates
ADD linuxamd64/bin/server /opt/app/bin/server
ENV TEMPLATE_PATH /opt/app/templates

WORKDIR /opt/app
ENTRYPOINT /opt/app/bin/server
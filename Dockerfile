FROM golang

COPY ./src /tmp/build/src

#For go projects, it is best to build the app directly in the container.
ENV GOPATH /tmp/build/
ENV GOBIN /usr/local/bin/
RUN go install github.com/patrickianwilson/hello

ENTRYPOINT /usr/local/bin/hello

EXPOSE 8080


#run via
#  docker run -t -p 8080:8080 <image id>
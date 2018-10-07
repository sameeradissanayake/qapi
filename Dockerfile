FROM golang:1.9

MAINTAINER Sameera Dissanayake (samees3dissanayake.com)

# install dependencies
RUN	go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2

# env
ENV MONGO_HOST 172.17.0.1

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/qapi src/*.go

RUN chmod 777 docker-entrypoint.sh
RUN chmod 777 build/qapi

# running on 8080
EXPOSE 12345

ENTRYPOINT ["/app/docker-entrypoint.sh"]

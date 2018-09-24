FROM golang:1.9

MAINTAINER Sameera Dissanayake (samees3dissanayake.com)

# install dependencies
RUN	go get github.com/gorilla/mux
RUN go get github.com/mongodb/mongo-go-driver/mongo
RUN go get github.com/mongodb/mongo-go-driver/bson

# env
ENV MONGO_HOST 192.169.0.1

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/qapi src/*.go

# running on 8080
EXPOSE 12345

ENTRYPOINT ["/app/docker-entrypoint.sh"]

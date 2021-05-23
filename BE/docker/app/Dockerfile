FROM golang:1.13.1

RUN apt-get update && apt-get install -y unzip && go get -u github.com/tockins/realize
WORKDIR /app
COPY / /app
CMD [ "realize", "start", "--run" ]
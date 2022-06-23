FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget
RUN apt-get install -y tzdata
RUN apt-get install -y software-properties-common
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb
RUN rm google-chrome-stable_current_amd64.deb 

VOLUME /assets

ADD loteca-backend_unix .
ADD firebase.json .

CMD ["./loteca-backend_unix"]
EXPOSE 8080
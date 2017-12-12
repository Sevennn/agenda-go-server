FROM golang:1.8
RUN mkdir -p /gocode/hw8/agenda-go-server
WORKDIR /gocode/hw8/agenda-go-server
COPY . .

WORKDIR /gocode/hw8/agenda-go-server/cli
RUN go-wrapper download
RUN go build -o agenda
RUN mv ./agenda /go/bin/

WORKDIR /gocode/hw8/agenda-go-server/service
RUN go-wrapper download
RUN go build -o agenda
RUN mv ./agendad/go/bin/

RUN mv ../agenda.sh /

WORKDIR /
ENTRYPOINT [ "agenda.sh" ]
CMD [ "agendad" ]
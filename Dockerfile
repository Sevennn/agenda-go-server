FROM golang:1.8
RUN mkdir -p /go/src/github.com/freakkid/service-agenda/
WORKDIR /gocode/src/github.com/freakkid/service-agenda/
COPY . .

WORKDIR /gocode/src/github.com/freakkid/service-agenda/cli
RUN go-wrapper download
RUN go build -o agenda
RUN mv ./agenda /go/bin/

WORKDIR /gocode/src/github.com/freakkid/service-agenda/service
RUN go-wrapper download
RUN go build -o agendad
RUN mv ./agendad /go/bin/

RUN mv ../agenda.sh /

WORKDIR /
ENTRYPOINT [ "agenda.sh" ]
CMD [ "agendad" ]
FROM golang:1.8 
COPY . "$gocode/src/github.com/Sevennn/agenda-go-server" 
RUN cd "$gocode/src/github.com/Sevennn/agenda-go-server/cli" && go get -v && go install -v 
RUN cd "$gocode/src/github.com/Sevennn/agenda-go-server/service" && go get -v && go install -v 
WORKDIR / 
EXPOSE 8080 
VOLUME ["/data"] 

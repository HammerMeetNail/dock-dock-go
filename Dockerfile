FROM golang:1.12.1

WORKDIR /go/src/github.com/HammerMeetNail/dock-dock-go
COPY . .

RUN go get -d -v 
RUN go install -v

ENTRYPOINT ["dock-dock-go"]

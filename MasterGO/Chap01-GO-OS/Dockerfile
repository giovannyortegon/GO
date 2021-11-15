FROM golang:alpine

cOPY hw.go ./files
WORKDIR ./files

RUN go build -o ./files/hw hw.go
ENTRYPOINT ["./files/hw"]

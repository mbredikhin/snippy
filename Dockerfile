FROM golang:1.15-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update && apt-get -y install postgresql-client

# install go-air
RUN go get github.com/cosmtrek/air

# exec wait-for-postgres.sh
RUN chmod +x ./scripts/wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o snippets ./cmd/main.go

CMD [ "./snippets" ]

FROM golang:1.21

RUN go version
ENV GOPATH=/

COPY ./ ./

# install postgres, packages for live reloading and migrations 
RUN apt-get update && apt-get -y install postgresql-client
RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# exec wait-for-postgres.sh
RUN chmod +x ./scripts/wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o snippets ./cmd/main.go

CMD [ "./snippets" ]

FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

# install postgres, packages for live reloading and migrations 
RUN apt-get update && apt-get -y install postgresql-client
RUN go install github.com/air-verse/air@v1.52.2
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

# exec wait-for-postgres.sh
RUN chmod +x ./scripts/wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o snippets ./cmd/main.go

CMD [ "./snippets" ]

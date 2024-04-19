FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

# install postgres, redis, packages for live reloading and migrations 
RUN apt-get update && apt-get -y install postgresql-client

RUN apt-get install -y lsb-release curl gpg 
RUN curl -fsSL https://packages.redis.io/gpg | gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/redis.list
RUN apt-get update
RUN apt-get install -y redis

RUN go install github.com/air-verse/air@v1.52.2
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

# exec wait-for-postgres.sh
RUN chmod +x ./scripts/wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o snippets ./cmd/main.go

CMD [ "./snippets" ]

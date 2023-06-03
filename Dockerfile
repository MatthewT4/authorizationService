FROM golang:latest

WORKDIR /app

COPY . .

#RUN go build -o bin .

#CMD ["make", "migration_up"]
#CMD ["make", "migration_down"]
#CMD ["make", "run"]
CMD ["go", "run", "cmd/main.go"]

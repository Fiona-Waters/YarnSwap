FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . .
EXPOSE 8080
CMD ["go", "run", "."]


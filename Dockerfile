FROM golang:1.18
RUN useradd fiona
WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . .
EXPOSE 8080
USER fiona
CMD ["go", "run", "."]


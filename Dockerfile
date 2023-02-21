

FROM golang:1.20 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY /pkg /workspace/pkg
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app main.go

FROM ubuntu
WORKDIR /
COPY --from=builder /workspace/app .

ENTRYPOINT ["/app"]
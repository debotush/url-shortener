FROM golang:1.19-alpine AS builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOFLAGS=-mod=mod GOOS=linux  go build -a -o /app ./cmd/url-shortener


FROM alpine AS final

EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["/app", "server"]
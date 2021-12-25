# build container
FROM golang:1.16.0-buster as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg

# environment variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# build
RUN go build -mod=readonly -ldflags="-w -s" -v -o bin/torasemi-todo cmd/gqlserver/main.go

# execute container
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/bin/torasemi-todo /app/torasemi-todo

# run
CMD ["./torasemi-todo"]

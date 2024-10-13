# builder
FROM golang:1.23-alpine AS builder
WORKDIR /usr/local/src

# 
RUN apk --no-cache add bash gcc gettext musl-dev

# dependencies
COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

# build
COPY . ./
RUN go build -o ./bin/blog-api cmd/blog-api/main.go

# run
FROM alpine AS runner
COPY --from=builder /usr/local/src/bin/blog-api ./
COPY config/local.yaml /local.yaml
RUN chmod +x blog-api
CMD ["/blog-api"]
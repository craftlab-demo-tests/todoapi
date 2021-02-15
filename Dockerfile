#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go test ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/todoapi /app
COPY ./templates /app/templates
ENTRYPOINT ./app
LABEL Name=todoapi Version=0.0.1
EXPOSE 8080

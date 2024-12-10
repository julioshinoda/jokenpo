# Build Stage

FROM golang:1.23-bullseye AS BuildStage

WORKDIR /app


COPY . .

RUN go mod download


EXPOSE 8080

RUN  CGO_ENABLED="0" GOOS=linux GOARCH=amd64 go build -o /api cmd/main.go

# Deploy Stage

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /api /api


EXPOSE 8080


ENTRYPOINT ["/api"]
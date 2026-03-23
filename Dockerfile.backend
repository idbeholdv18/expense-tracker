FROM golang:1.26.1-alpine AS base

WORKDIR /usr/src/app
COPY go.sum go.mod ./
RUN go mod download

COPY . .

FROM base AS backend
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/api ./cmd/api/main.go

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /usr/server/bin
COPY --from=backend /usr/local/bin/api .
COPY --from=base /usr/src/app/cert ./cert

EXPOSE 8080

ENTRYPOINT [ "./api" ]

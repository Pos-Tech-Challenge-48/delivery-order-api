FROM golang:1.21-alpine

# Create an unprivileged user. https://stackoverflow.com/a/49955098/2387190
RUN adduser -D -H -h "/nonexistent" -s "/sbin/nologin" -g "" -u "10001" "appuser"

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /delivery-api ./cmd/api
RUN chmod a+x /delivery-api

WORKDIR /app

EXPOSE 8080
EXPOSE 2345

USER appuser:appuser

CMD ["/delivery-api"]

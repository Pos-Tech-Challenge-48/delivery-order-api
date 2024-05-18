FROM golang:1.22.0-bookworm

# Create an unprivileged user. https://stackoverflow.com/a/49955098/2387190
# RUN adduser -D -H -h "/nonexistent" -s "/sbin/nologin" -g "" -u "10001" "appuser"

# Create an unprivileged user. https://stackoverflow.com/a/49955098/2387190
RUN groupadd --gid 1001 appuser && useradd --uid 1001 --gid 1001 --shell /bin/bash --create-home appuser

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

USER appuser

CMD ["/delivery-api"]

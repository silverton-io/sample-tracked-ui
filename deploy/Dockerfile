FROM golang:1.18 AS build

WORKDIR /app

COPY . /app/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.VERSION=$(cat .VERSION)" -o tracked-ui ./cmd/*.go

FROM busybox

LABEL maintainer="Jake Thomas <jake@silverton.io>"
LABEL org.opencontainers.image.description "A lightweight, Snowplow-compatible streaming event collection system."

WORKDIR /app
COPY --from=build /app/.VERSION .
COPY --from=build /app/sample.html .
COPY --from=build /app/tracked-ui .
COPY --from=build /etc/ssl/certs /etc/ssl/certs


EXPOSE 8080

ENTRYPOINT [ "./tracked-ui" ]

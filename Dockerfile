FROM golang:1.26 AS builder

WORKDIR /src

COPY go.mod ./
COPY app/ ./app/

WORKDIR /src/app

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /preview-app .

FROM scratch

COPY --from=builder /preview-app /preview-app

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/preview-app"]

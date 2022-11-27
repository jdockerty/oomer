FROM golang:1.19-alpine as builder

LABEL org.opencontainers.image.source https://github.com/jdockerty/oomer

WORKDIR /app

COPY go.mod main.go ./

RUN go build -o oomer .

FROM scratch

COPY --from=builder /app/oomer /app/oomer

USER 1001

ENTRYPOINT ["/app/oomer"]

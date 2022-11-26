FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod main.go ./

RUN go build -o oomer .

FROM scratch

COPY --from=builder /app/oomer /app/oomer

ENTRYPOINT ["/app/oomer"]

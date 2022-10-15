FROM golang:1.19 as builder

WORKDIR /app
ADD . .
RUN go build -o gurl cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/gurl .
CMD [ "./gurl" ]
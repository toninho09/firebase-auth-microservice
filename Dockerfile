FROM golang:1.9.2 as builder
WORKDIR /app
RUN go get -u firebase.google.com/go
RUN go get -u github.com/gin-gonic/gin
COPY .  .
RUN CGO_ENABLED=0 GOOS=linux go build -o microservice .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/microservice .
COPY --from=builder /app/api.json .
EXPOSE 8080
CMD ["./microservice"]
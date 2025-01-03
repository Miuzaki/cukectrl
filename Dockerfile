FROM  golang:alpine as builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /cmd/dash/main .

FROM alpine

WORKDIR /app    
COPY --from=builder /app/main .

CMD ["./main"]

FROM  golang:alpine as builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o cukectrl cmd/cukectrl/main.go 

FROM alpine

WORKDIR /app    
COPY --from=builder /app/cukectrl .

CMD ["./cukectrl"]

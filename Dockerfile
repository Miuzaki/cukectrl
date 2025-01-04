FROM  golang:1.22-alpine as builder

WORKDIR /build


COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build  -o cukectrl cmd/cukectrl/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/cukectrl /app/

CMD ["./cukectrl"]
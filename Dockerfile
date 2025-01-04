FROM  golang:alpine as builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

ENV GOCACHE=/root/.cache/go-build

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=cache,target="/root/.cache/go-build" \
  --mount=type=bind,target=. \
  go build -o cukectrl cmd/cukectrl/main.go 

FROM alpine

WORKDIR /app    
COPY --from=builder /app/cukectrl .

CMD ["./cukectrl"]

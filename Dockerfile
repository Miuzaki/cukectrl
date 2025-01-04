FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Adicione este comando para verificar o conteúdo do diretório
RUN ls -R /app

# Ajuste o caminho caso necessário
RUN go build -o /cmd/cukectrl/main .

FROM alpine

WORKDIR /app    

COPY --from=builder /cmd/cukectrl/main .

CMD ["./main"]

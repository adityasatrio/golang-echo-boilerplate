FROM golang:1.22-alpine AS builder

ENV CGO_ENABLED=0
ENV OPENAPI_ENTRY_POINT="cmd/main.go"
ENV OPENAPI_OUTPUT_DIR="cmd/docs"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN apk add --no-cache make
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swagger
RUN go build -ldflags="-s -w" -o main ./cmd/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/main /app/main
USER 65532:65532
ENTRYPOINT ["/app/main"]
EXPOSE 8888

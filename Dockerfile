FROM golang:1.22-alpine AS builder

ENV OPENAPI_ENTRY_POINT="cmd/main.go"
ENV OPENAPI_OUTPUT_DIR="cmd/docs"
WORKDIR /app
COPY . /app
RUN apk add --no-cache make
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swagger
RUN make build

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/main /app/main
ENTRYPOINT [ "/app/main" ]
EXPOSE 8888
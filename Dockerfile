FROM 255045239074.dkr.ecr.ap-southeast-1.amazonaws.com/golang:1.19-alpine AS builder

ENV OPENAPI_ENTRY_POINT="cmd/main.go"
ENV OPENAPI_OUTPUT_DIR="cmd/docs"
WORKDIR /app
COPY . /app
RUN apk add --no-cache make
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swagger
RUN make build
# USER 1000
ENTRYPOINT [ "/app/main" ]
EXPOSE 8888

# FROM gcr.io/distroless/static
# WORKDIR /app
# COPY --from=builder /app/main /app/main
# USER 1000
# ENTRYPOINT [ "/app/main" ]
# EXPOSE 8888
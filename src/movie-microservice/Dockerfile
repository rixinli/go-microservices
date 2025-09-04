# ---------- Build stage ----------
FROM golang:1.22-alpine AS builder
WORKDIR /build


# copy go.mod go.sum ./
COPY go.mod go.sum ./
RUN go mod download

# copy src ./src
COPY src ./src

# compile movie microservice
WORKDIR /build/src/movie-microservice
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/movie-microservice

# ---------- Runtime stage ----------
FROM alpine:3.20
RUN adduser -D -H appuser && apk add --no-cache ca-certificates

# copy the microservice app
WORKDIR /app
COPY --from=builder /out/movie-microservice /app/movie-microservice
COPY src/movie-microservice/config /app/config
RUN mkdir /app/logs && chown -R appuser:appuser /app/logs

EXPOSE 8809
USER appuser
ENTRYPOINT [ "./movie-microservice" ]
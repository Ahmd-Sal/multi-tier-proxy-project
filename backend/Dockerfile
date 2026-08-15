# Stage 1: Build the Go binary
FROM golang:1.26.5 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Distroless Runtime
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/myapp /myapp
USER nonroot:nonroot
EXPOSE 80
ENTRYPOINT ["/myapp"]

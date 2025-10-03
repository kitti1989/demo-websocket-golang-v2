# ------------------------
# Stage 1: Build
# ------------------------
FROM golang:1.25 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod และ go.sum ก่อน เพื่อแยก layer cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o server app.go

# ------------------------
# Stage 2: Run
# ------------------------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary จาก builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run binary
CMD ["./server"]

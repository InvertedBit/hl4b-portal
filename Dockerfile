# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git nodejs npm

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy package.json for npm
COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Build Tailwind CSS
RUN npm run build:css

# Build the application
RUN go build -o hl4b-portal main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary and static files from builder
COPY --from=builder /app/hl4b-portal .
COPY --from=builder /app/static ./static

# Create uploads directory
RUN mkdir -p /app/uploads

# Expose port
EXPOSE 3000

# Run the application
CMD ["./hl4b-portal"]

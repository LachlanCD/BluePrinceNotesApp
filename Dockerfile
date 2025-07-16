# Stage 1: Build React frontend
FROM node:24.1 AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.24 AS backend-builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./backend/
WORKDIR /app/backend
RUN go mod download

COPY backend/ /app/backend/
RUN go build -o /app/server

# Stage 3: Final image
FROM debian:bookworm-slim
WORKDIR /app

# Copy built Go binary
COPY --from=backend-builder /app/server .
# Copy built frontend
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
# Copy database (optional; or mount it at runtime)
COPY backend/data/notes.db ./data/notes.db

EXPOSE 3000

CMD ["./server"]


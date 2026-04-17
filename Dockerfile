# Stage 1: Build frontend
FROM node:22-alpine AS frontend
WORKDIR /build
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci
COPY frontend/ ./
# vite.config.ts aliases @schemas to ../schemas (relative to frontend/),
# which resolves to /schemas when WORKDIR is /build.
COPY schemas/ /schemas/
RUN npm run build
# RUN npx vite build  # TODO: temporary workaround, remove once TS errors are fixed

# Stage 2: Build backend
FROM golang:1.24-alpine AS backend
WORKDIR /build
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 go build -o server ./cmd/server/

# Stage 3: Final image
FROM alpine:3
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=backend /build/server ./
COPY --from=frontend /build/dist/ ./frontend/dist/
COPY schemas/ /schemas/
ENV AVHBS_FRONTEND_PATH=/app/frontend/dist
EXPOSE 8081
CMD ["./server"]

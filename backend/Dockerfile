# --- Build stage ---
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# --- Run stage ---
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/server /app/server
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]

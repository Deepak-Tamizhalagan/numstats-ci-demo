# ---------- build stage ----------
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build -o /out/numstats ./cmd/numstats

# ---------- runtime stage ----------
FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/numstats /app/numstats
ENTRYPOINT ["/app/numstats"]

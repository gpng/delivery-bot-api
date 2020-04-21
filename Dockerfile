FROM golang:1.14.2-alpine

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser.
RUN adduser -D -g '' appuser

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/application cmd/api/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

# Import the user and group files from the 0.
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /etc/passwd /etc/passwd

# Copy our static executable.
COPY --from=0 /app/bin/application /app/application
# Copy production .env
COPY --from=0 /app/.prod.env /app/.env

# Use an unprivileged user.
USER appuser

EXPOSE 5000

WORKDIR /app

# Run the hello binary.
ENTRYPOINT ["./application"]
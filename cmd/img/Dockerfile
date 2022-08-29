# Info here: https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

############################
# STEP 1 build executable binary
############################
FROM golang:1.18-alpine as builder

# Install git, ssl ca certs.
# Git is required for fetching the dependencies.
# SSL ca certs required to call HTTPS endpoints.
RUN apk update \
    && apk add --no-cache git ca-certificates \
    && update-ca-certificates

# Create an unprivileged user and group.
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    "appuser"

# Copy app.
WORKDIR /go/src/github.com/adammy/memepen-services
COPY . .

# Fetch dependencies.
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /memepen-image-svc cmd/image/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

# Import the ca-certificate.crt from the builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy the binary and it's dependencies.
COPY --from=builder --chown=appuser:appuser /memepen-image-svc /memepen-image-svc
COPY --from=builder --chown=appuser:appuser /go/src/github.com/adammy/memepen-services/assets ./assets
COPY --from=builder --chown=appuser:appuser /go/src/github.com/adammy/memepen-services/configs ./configs

# Use an unprivileged user.
USER appuser:appuser

# Run the binary.
EXPOSE 8080
ENTRYPOINT ["/memepen-image-svc"]
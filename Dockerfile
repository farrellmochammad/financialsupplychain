FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

RUN go mod tidy
RUN sh script/seed.sh

EXPOSE 2021
# ENTRYPOINT ["sh"]
CMD ["go", "run", "cmd/main.go"]


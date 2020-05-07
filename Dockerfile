FROM golang:alpine
WORKDIR /app
RUN apk add --no-cache git \
	&& go get github.com/cespare/reflex
CMD reflex -r '\.go' -s go run .
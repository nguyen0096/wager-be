ARG BUILD_OUTPUT

FROM golang:1.17-alpine AS build-env
RUN apk --no-cache add build-base git mercurial gcc bash
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

FROM alpine
ENV BUILD_OUTPUT=$BUILD_OUTPUT
WORKDIR /app
COPY --from=build-env /app/bin/wager-be /app/
COPY --from=build-env /app/migration /app/migration
ENTRYPOINT ./wager-be server

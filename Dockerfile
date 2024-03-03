FROM golang:1.22 as build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /upload-server

FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /

COPY --from=build-stage /upload-server /upload-server
EXPOSE 8080
ENV GIN_MODE=release

ENTRYPOINT ["/upload-server"]
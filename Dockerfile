FROM golang:alpine AS build
ENV CGO_ENABLED 0
ENV GO111MODULE=on
ENV GOOS=linux

RUN apk add make build-base git curl openssh
WORKDIR /go/src/bank-exmaple
COPY . .
RUN go mod vendor
RUN go build -o /go/bin/bank-example cmd/bank-example/main.go

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /go/bin/bank-example /app/
ENTRYPOINT ["./bank-example"]
EXPOSE 8080

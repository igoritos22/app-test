FROM golang:alpine AS builder

LABEL maintainer="Igor Luiz de Sousa Santos, iluiz.sousa@gmail.com"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /build

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

#build with a minimal image
FROM alpine

COPY --from=builder /dist/main /

EXPOSE 8080

ENTRYPOINT [ "/main" ]
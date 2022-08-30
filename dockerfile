FROM golang:alpine as build-env
RUN apk update && apk add --no-cache git

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o binary

FROM alpine
WORKDIR /app
COPY --from=build-env /src/binary /app
ENV PORT 8080
ENV DB_HOST localhost
ENV DB_PORT 5432
ENV DB_USER postgres
ENV DB_NAME postgres
ENV DB_PASSWORD postgres
ENTRYPOINT ["./binary"]
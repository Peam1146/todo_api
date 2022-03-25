FROM golang:1.18-alpine3.15 AS base
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


FROM base AS build
RUN go build -o main src/main.go

FROM alpine:3.15 as prod
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 3000
CMD [ "./main" ]
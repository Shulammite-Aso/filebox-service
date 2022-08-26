# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /filebox-service

FROM scratch

WORKDIR /

COPY --from=build /filebox-service /filebox-service

EXPOSE 50052

CMD [ "/filebox-service" ]
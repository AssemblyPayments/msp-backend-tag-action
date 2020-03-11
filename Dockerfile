FROM golang:1.13-alpine as build

# build
RUN mkdir /project
COPY main.go /project/
COPY go.* /project/
COPY vendor /project/vendor
WORKDIR /project
RUN go build -mod=vendor -o /app/main .

# run
FROM alpine:3.7
COPY --from=build /app/main /app/main
WORKDIR /app
CMD /app/main

ENTRYPOINT ["/app/main"]

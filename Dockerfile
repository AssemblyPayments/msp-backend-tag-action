FROM golang:1.13-alpine

# build the action
RUN mkdir /project
COPY main.go /project/
COPY go.* /project/
COPY vendor /project/vendor
WORKDIR /project
RUN go build -mod=vendor -o /app/main .
WORKDIR /app

# # run the action
# FROM alpine:3.7
# COPY --from=build /app/main /app/main
# WORKDIR /app
# CMD /app/main

ENTRYPOINT ["/app/main"]

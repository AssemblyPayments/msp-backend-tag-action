FROM golang:1.13-alpine as build

# build and run action
# we could use a "builder" dockerfile to build the action into a docker image and another dockerfile
# referencing that. But it's a bit confusing not very straightward. So here just use single dockerfile
# to build and run the action. The downside is the base image is much bigger (+100MB).
RUN mkdir /project
COPY main.go /project/
COPY go.* /project/
COPY vendor /project/vendor
WORKDIR /project
RUN go build -mod=vendor -o /app/main .

# run the action
FROM alpine:3.7
COPY --from=build /app/main /app/main
WORKDIR /app
CMD /app/main

ENTRYPOINT ["/app/main"]

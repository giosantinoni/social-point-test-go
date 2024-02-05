FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/test
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/project-api cmd/api/main.go
FROM scratch
COPY --from=build /go/bin/project-api /go/bin/project-api
ENTRYPOINT ["/go/bin/project-api"]

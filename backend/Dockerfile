FROM golang:1.20.4-alpine3.18 as builder

COPY go.mod go.sum /go/src/github.com/4hm3d92/community-app/backend/
WORKDIR /go/src/github.com/4hm3d92/community-app/backend
RUN go mod download
COPY . /go/src/github.com/4hm3d92/community-app/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/community github.com/4hm3d92/community-app/backend


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/4hm3d92/community-app/backend/build/community /usr/bin/community

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/community"]

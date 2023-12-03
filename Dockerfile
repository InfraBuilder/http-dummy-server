# Go library image
FROM golang:1.21.1 AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o http-dummy-server .

FROM busybox
COPY --from=builder /app/http-dummy-server /usr/bin/local/http-dummy-server
ENV HTTP_RESPONSE_BODY="ok" \
    HTTP_RESPONSE_CODE="200"
CMD ["/usr/bin/local/http-dummy-server"]



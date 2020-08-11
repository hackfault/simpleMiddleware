FROM golang:1.15rc2-alpine3.12
WORKDIR /app/src
COPY . .
RUN go build -o capture
CMD ["/app/src/capture"]
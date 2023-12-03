FROM golang:1.21.1
WORKDIR /app
COPY ./ ./
RUN go mod download
WORKDIR /app/cmd
RUN go build -buildvcs=false -o urlshortener
EXPOSE 8083
ENTRYPOINT ["./urlshortener"]
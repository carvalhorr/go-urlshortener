FROM golang:latest
LABEL maintainer="Rodrigo de Carvalho <carvalhorr@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod download
WORKDIR /app/urlshortener
RUN go build -o urlshortener
EXPOSE 8081
CMD ["./urlshortener"]
FROM golang:latest
LABEL maintainer="Rodrigo de Carvalho <carvalhorr@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod download
WORKDIR /app/query
RUN go build -o query
EXPOSE 8081
CMD ["./query"]
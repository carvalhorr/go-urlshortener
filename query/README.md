## Building the application

From `go-urlshortener`

Run `docker build -t urlshortener-query:0.1 -f query/Dockerfile .`

Then run `docker run urlshortener-query:0.1`

Your application will be listening on port `8081`

To start the docker container:
`docker run -p 8081:8081 urlshortener-query:0.1`
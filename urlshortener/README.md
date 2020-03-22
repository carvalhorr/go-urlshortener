## Building the application

From `go-urlshortener`

Run `docker build -t urlshortener:0.1 -f urlshortener/Dockerfile .`

Then run `docker run urlshortener:0.1`

Your application will be listening on port `8081`

To start the docker container:

`docker run -p 8081:8081 urlshortener:0.1`
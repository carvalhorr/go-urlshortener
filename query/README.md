## Building the application

From `go-urlshortener`

Run `docker build -t urlshortener-query:0.1 -f query/Dockerfile .`

Then run `docker run urlshortener-query:0.1`

Your application will be listening on port `8081`

To start the docker container:
`docker run -p 8081:8081 urlshortener-query:0.1`

## Publishing to Docker Hub

* Login to docker hub using `docker login`
* Tag the local image using `docker tag urlshortener-query:0.1 carvalhorr/urlshortener-query:0.1`
* Push tp registry using `docker push carvalhorr/urlshortener-query:0.1`
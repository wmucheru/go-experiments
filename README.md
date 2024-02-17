# Deploy a Go Web Application with Docker

- [Simple GO web app with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker)
- [GO multi-stage build with Docker](https://medium.com/geekculture/optimizing-golang-docker-images-with-multi-stage-builds-ca7b305faa)

## Run with Docker

- `docker build --no-cache -t <TAG> .`
- `docker run -it --rm -p 8010:8010 <TAG>`

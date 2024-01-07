## setup

https://github.com/OpenAPITools/openapi-generator

```sh
$ docker pull swaggerapi/swagger-editor
$ docker run -p 80:8080 -v $(pwd):/tmp/ -e SWAGGER_FILE=/tmp/docs/swagger.yaml swaggerapi/swagger-editor
```

## run

```sh
$ brew install openapi-generator
$ openapi-generator generate -i ./docs/swagger.yaml -g go-server -o app/ --git-repo-id waxflower --git-user-id k-jun
```

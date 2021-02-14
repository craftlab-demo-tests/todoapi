# generate clients

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/craftlab-demo-tests/user-service/master/server/api/openapi.yaml \
    -g go \
    -o /local/user \
    --additionalproperties=packageName=userclient


docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/craftlab-demo-tests/todo-service/master/openapi.yaml \
    -g go \
    -o /local/todo \
    --additional-properties=packageName=todoclient

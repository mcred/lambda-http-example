# Lambda HTTP Router Example

This is a simple example to test the performance of a simple HTTP router using AWS Lambda.

## Requirements
* [go](https://golang.org/)
* [docker](https://www.docker.com/)
* [docker-compose](https://docs.docker.com/compose/install/)
* [awslocal-cli](https://github.com/localstack/awscli-local)
* [tflocal-cli](https://github.com/localstack/terraform-local)

## Running the example
1. Set LOCALSTACK_API_KEY to the env
```bash
echo LOCALSTACK_API_KEY=xxxxxxx >> .env
``` 
2. Start the stack
```bash
make start
```
3. Running the example for the first time
```bash
make init
make apply
```
4. Refreshing the code after the first time
```bash
make refresh
```
5. Destroying the example
```bash
make destroy
```
6. Stopping the stack
```bash
make stop
``` 

## Testing the router
The output from the `make apply` command will show the URL of the API Gateway. Export this value to the `API_URL` environment variable. Then use `curl` to test the router.

```bash
export API_URL=http://{API_ID}.execute-api.localhost.localstack.cloud:4566
curl -X GET $API_URL
curl -X GET $API_URL/users
curl -X GET $API_URL/users/1
curl -X DELETE $API_URL/users/1
```  
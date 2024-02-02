.PHONY: start stop refresh init apply destroy

start:
	docker-compose up -d

stop:
	docker-compose down

build:
	GOOS=linux go build -o ./bin/main && cd bin && zip main.zip main

refresh: build
	localstack awslocal lambda update-function-code --function-name app --zip-file fileb://bin/main.zip

init:
	cd terraform && tflocal init

apply: build
	cd terraform && tflocal apply -auto-approve

destroy:
	cd terraform && tflocal destroy -auto-approve
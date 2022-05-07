.PHONY: init local_run dev_build dev_deploy

init:
	cd .. && yarn

dev_build:
	mkdir -p ./build ./build/develop
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./build/develop/main main.go
	zip ./build/develop/deployment.zip ./build/develop/main

local_run: init dev_build
	cd .. && npx sls offline

dev_deploy: init dev_build
	cd .. && npx sls deploy function --function bot --verbose

# productionへはGitHub Actionsで

.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/stock_checker ./main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

local: clean build
	sls invoke local --function stock_checker --verbose

setup:
	@if [ ! -e slack.yml ]; then touch slack.yml && echo "webhook_url: https://hooks.slack.com/services/*********/*********/************************" >> slack.yml;else echo "slack.yml is already exists."; fi

destroy:
	sls remove

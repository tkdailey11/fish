appname  := az-fun-go-sp
funcRoot := ./functions

build:
	go build -o ./bin/server ./

build-prod: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -tags "prod" -o $(funcRoot)/bin/server ./

pack: build-prod
	@mkdir -p infra/package
	cd $(funcRoot) && func pack -o ../infra/package/functions

install:
	go get -u ./... && go mod tidy

clean:
	rm -rf bin/ tmp/ $(funcRoot)/bin/ $(funcRoot)/tmp/ infra/package/
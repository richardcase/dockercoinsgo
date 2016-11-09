all: clean copycontent build 

copycontent:
	mkdir -p ./build
	cp -avr ../../webui ./build/webcontent

build: buildapps buildimages

buildapps:
	CGO_ENABLED=0 GOOS=linux go build -a -o build/rng github.com/richardcase/dockercoinsgo/cmd/rng
	CGO_ENABLED=0 GOOS=linux go build -a -o build/hasher github.com/richardcase/dockercoinsgo/cmd/hasher
	CGO_ENABLED=0 GOOS=linux go build -a -o build/worker github.com/richardcase/dockercoinsgo/cmd/worker
	CGO_ENABLED=0 GOOS=linux go build -a -o build/webui github.com/richardcase/dockercoinsgo/cmd/webui

buildimages: buildapps
	docker build -t richardcase/dockercoinsgo-rng -f Dockerfile-rng .
	docker build -t richardcase/dockercoinsgo-hasher -f Dockerfile-hasher .
	docker build -t richardcase/dockercoinsgo-worker -f Dockerfile-worker .
	docker build -t richardcase/dockercoinsgo-webui -f Dockerfile-webui .

clean:
	rm -rf build/
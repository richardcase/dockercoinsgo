VAR_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VAR_REF=$(shell git rev-parse --short HEAD)
VAR_VER=$(shell cat ../../VERSION)


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
	@echo Build Date: $(VAR_DATE) 
	@echo Git Ref: $(VAR_REF)
	@echo Version: $(VAR_VER)
	docker build -t richardcase/dockercoinsgo-rng -f Dockerfile-rng --build-arg BUILD_DATE=$(VAR_DATE) --build-arg VCS_REF=$(VAR_REF) --build-arg VERSION=$(VAR_VER) .
	docker build -t richardcase/dockercoinsgo-hasher -f Dockerfile-hasher --build-arg BUILD_DATE=$(VAR_DATE) --build-arg VCS_REF=$(VAR_REF) --build-arg VERSION=$(VAR_VER) .
	docker build -t richardcase/dockercoinsgo-worker -f Dockerfile-worker --build-arg BUILD_DATE=$(VAR_DATE) --build-arg VCS_REF=$(VAR_REF) --build-arg VERSION=$(VAR_VER) .
	docker build -t richardcase/dockercoinsgo-webui -f Dockerfile-webui --build-arg BUILD_DATE=$(VAR_DATE) --build-arg VCS_REF=$(VAR_REF) --build-arg VERSION=$(VAR_VER) .

clean:
	rm -rf build/

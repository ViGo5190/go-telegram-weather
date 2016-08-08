default: build

clean:
	export GOPATH=$(shell pwd);  \
	rm -f bin/weather-bot

build:
	export GOPATH=$(shell pwd);  \
	cd src/weather-bot; \
	go build -o ../../bin/weather-bot

deps:
	export GOPATH=$(shell pwd);  \
	cd src; \
	gvt restore

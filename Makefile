.PHONY: pack build

build: build/hh_webhook/main

pack: build
	@echo "All done"

build/hh_webhook/main: src/hh_webhook/*
	cd src/hh_webhook && GOOS=linux GOARCH=amd64 go build -o ../../build/hh_webhook/main
	cd build/hh_webhook && zip function.zip main

deploy: pack
	cd terraform && terraform apply

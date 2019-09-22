.PHONY: pack build

build: build/hh_webhook/main build/hh_trigger/main build/hh_cronjob/main

pack: build
	@echo "All done"

build/hh_webhook/main: src/hh_webhook/*
	cd src/hh_webhook && GOOS=linux GOARCH=amd64 go build -o ../../build/hh_webhook/main
	cd build/hh_webhook && zip function.zip main

build/hh_trigger/main: src/hh_trigger/*
	cd src/hh_trigger && GOOS=linux GOARCH=amd64 go build -o ../../build/hh_trigger/main
	cd build/hh_trigger && zip function.zip main

build/hh_cronjob/main: src/hh_cronjob/*
	cd src/hh_cronjob && GOOS=linux GOARCH=amd64 go build -o ../../build/hh_cronjob/main
	cd build/hh_cronjob && zip function.zip main

deploy: pack
	cd terraform && terraform apply -var-file="testing.tfvars"

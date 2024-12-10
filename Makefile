IP_ADDRESS := $(shell ipconfig getifaddr en0)

.PHONY: run
run: 
	go run cmd/main.go

run-air:
	air

tests:
	go test --race ./... 

run-k6-loadtest:
	cd loadtest && docker run -e IP_ADDRESS=$(IP_ADDRESS) --rm -i grafana/k6 run - <script.js

lint:
	golangci-lint run ./...	

vulncheck:
	govulncheck ./... 

build-image:
	docker build -t jokenpo:latest .	
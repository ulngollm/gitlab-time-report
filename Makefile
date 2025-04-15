include .env

build:
	CGO_ENABLED=0 go build -o gl-report

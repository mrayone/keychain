install:
	go mod vendor && go mod tidy


test:
	go test ./... -count=1

start/api:
	$(_reflex_) -- go run main.go api

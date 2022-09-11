_reflex_ = reflex -d none -s -R vendor. -r '.*\.go'

install:
	go mod vendor && go mod tidy

test:
	go test ./... -count=1

dev/api:
	$(_reflex_) -- go run main.go api

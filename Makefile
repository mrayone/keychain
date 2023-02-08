_reflex_ = reflex -d none -s -R vendor. -r '.*\.go'

install: install-tools
	go mod vendor && go mod tidy

install-tools:
	cat tools/tools.go | grep "_" | awk -F '"' '{print $$2"@latest"}' | xargs -L1 go install

test/unit:
	go test -race ./... -count=1 -short

test/e2e:
	echo "running e2e"

lint:
	gocritic check ./...

dev/api:
	$(_reflex_) -- go run main.go api

swagger:
	swag init -generalInfo cmd/webserver/http.go -output ./docs/swagger

install-hooks:
	./scripts/init-git.sh

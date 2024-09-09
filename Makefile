v ?= 1.0.0

pkg-some:
	wails build -platform windows/amd64,windows/arm64,linux/amd64
	upx --best\
		build/bin/safelock-linux-amd64\
		build/bin/safelock-amd64.exe\
		build/bin/safelock-arm64.exe

pkg:
	wails build -platform linux/amd64 -o safelock-$(v)
	upx build/bin/safelock-$(v)

test:
	go test ./... -count=2
	npm --prefix frontend test

test-fe:
	npm --prefix frontend run test:watch

lint:
	golangci-lint run
	npm --prefix frontend run lint

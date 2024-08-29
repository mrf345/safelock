v ?= 1.0.0

pkg-all:
	wails build -platform windows/amd64,windows/arm64,linux/amd64
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o build/bin/safelock-linux-arm64
	upx --best\
		build/bin/safelock-linux-amd64\
		build/bin/safelock-linux-arm64\
		build/bin/safelock-windows-amd64\
		build/bin/safelock-window-arm46

# NOTE: name of the binary needs to be {name}-{version} on linux for desktop entry
pkg:
	wails build -platform linux/amd64 -o safelock-$(v)
	upx build/bin/safelock-$(v)

test:
	go test ./... -count=2
	npm --prefix frontend test

lint:
	golangci-lint run
	npm --prefix frontend run lint

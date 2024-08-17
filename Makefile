app_id ?= mrf345.github.safelock
icon ?= assets/icon.png
sdk ?= ~/Projects/safelock/fyne-cross/MacOSX11.3.sdk

test:
	go test -count=2 ./...
lint:
	golangci-lint run
release:
	# fyne-cross windows -arch=amd64 -icon $(icon) -app-id $(app_id)
	# fyne-cross darwin -arch=amd64 -icon $(icon) -app-id $(app_id) -macosx-sdk-path $(sdk)
	# fyne-cross linux -arch=amd64 -icon $(icon) -app-id $(app_id)
	fyne package -os darwin -appID $(app_id) --icon $(icon) 


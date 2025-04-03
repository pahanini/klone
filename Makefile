build:
	GOOS=linux GOARCH=amd64 go build -o klone main.go && \
    tar -czvf klone.amd64.tar.gz klone

	GOOS=darwin GOARCH=arm64 go build -o klone main.go && \
    tar -czvf klone.arm64.tar.gz klone



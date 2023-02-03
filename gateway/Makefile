bin/gateway:
	go build -o bin/gateway

run: bin/gateway
	bin/gateway --driver sqlite --db "file::memory:?cache=shared" \
		--secret 0123456789abcdef --storage /tmp/mdouyin

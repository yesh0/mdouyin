docker-files: Dockerfile docker-compose.yml

Dockerfile: scripts/gen-dockerfile.sh go.work
	sh scripts/gen-dockerfile.sh > Dockerfile

docker-compose.yml: scripts/gen-docker-compose.sh go.work
	sh scripts/gen-docker-compose.sh > docker-compose.yml

build: docker-files
	sh scripts/for-all.sh make -B

clean:
	rm gateway/bin/gateway */output/bin/*Service

.PHONY: build clean docker-files
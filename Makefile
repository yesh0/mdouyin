docker-files: Dockerfile docker-compose.yml

Dockerfile: scripts/gen-dockerfile.sh go.work
	sh scripts/gen-dockerfile.sh > Dockerfile

docker-compose.yml: scripts/gen-docker-compose.sh go.work
	sh scripts/gen-docker-compose.sh > docker-compose.yml

build: docker-files
	sh scripts/for-all.sh make -B

remote-clean:
	docker --context mdouyin container rm \
			$(shell docker --context mdouyin container list -a --format '{{.ID}} {{.Image}}'\
			| grep mdouyin | awk '{ print $$1 }')
	docker --context mdouyin image prune -a

remote-send: docker-files
	docker save mdouyin-counter:latest  | gzip | docker --context mdouyin load
	docker save mdouyin-feeder:latest   | gzip | docker --context mdouyin load
	docker save mdouyin-gateway:latest  | gzip | docker --context mdouyin load
	docker save mdouyin-message:latest  | gzip | docker --context mdouyin load
	docker save mdouyin-reaction:latest | gzip | docker --context mdouyin load

clean:
	rm gateway/bin/gateway */output/bin/*Service

.PHONY: build clean docker-files remote-clean remote-send
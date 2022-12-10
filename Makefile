docker-build:
	docker build --no-cache --build-arg  VERSION=1.19.3-alpine -t queuefly -f Dockerfile .

docker-up:
	docker-compose -f docker-compose.yml up -d queuefly

docker-run:
	docker run -d -p 3000:3000 --name queuefly queuefly

.PHONY: docker-build docker-run docker-up
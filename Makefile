init:
	bash build/install_server.sh

up:
	docker-compose -f build/docker-compose.yml up
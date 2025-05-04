DC=docker-compose -f build/docker-compose.yml

up:
	$(DC) up --build

down:
	$(DC) down

restart: down up

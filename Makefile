.PHONY: ps up down down-clean

ps:
	docker compose ps

up:
	docker compose up -d

down:
	docker compose down

down-clean:
	docker compose down --rmi all
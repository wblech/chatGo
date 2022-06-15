

start_dev:
	mv .env.local .env
	docker-compose up -d --remove-orphans

stop_dev:
	mv .env .env.local
	docker-compose down

re: stop_dev
	docker volume prune
	docker rmi -f "$(docker images -aq)"


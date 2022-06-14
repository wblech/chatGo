

start_dev:
	docker-compose up -d --remove-orphans
	chmod +x ./.docker/wait-for-it.sh
	#sh ./.docker/wait-for-it.sh --host=localhost --port=8080/auth/realms/chat/.well-known/openid-configuration --timeout=60
#	sleep 30
#	go run ./src/main.go
	#docker build --tag=codingchalleng/go .
#	docker run \
#		--rm \
#		--tty \
#		--interactive \
#		--publish="8081:8081" \
#		--name="chat_app" \
#		--net="chat-network"
#		codingchalleng/go

stop_dev:
	docker-compose down

re: stop_dev
	docker volume prune
	docker rmi -f "$(docker images -aq)"


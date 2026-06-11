app-up : MySQL	
	@sleep 8
	@go mod tidy
	@go run .

MySQL : Dockerfile docker-compose.yml database_init.sql
	@open -a Docker
	@sleep 2
	@docker compose build mysql
	@docker compose up -d

reset-db:
	@docker compose down
	@rm -rf ./db
	@docker rmi saturn_mysql_image:latest
	@echo "Database wiped! Run 'make app-up' to rebuild and initialize."

reset-docker:
	@docker stop saturn_mysql_container
	@docker rm saturn_mysql_container
	@ocker rmi saturn_mysql_image:latest

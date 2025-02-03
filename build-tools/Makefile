# Environment file and variables
ENV_FILE := .env

## up_build: Stops and removes all running containers, builds the project and starts docker-compose
build: stop_all_containers down build_logger_service
	echo "🚀 Building (when required) and starting docker images with environment variables..."
	docker-compose up --build -d
	echo "✅ Docker images built and started!"
	echo "📜 Fetching logs for all services..."
	docker-compose logs --tail=50
	echo "🚀 Runnig Containers:"
	docker ps

## stop_all_containers: Stops and removes all running Docker containers (if any exist)
stop_all_containers:
	echo "🔍 Checking for running containers..."
	@if [ -n "$$(docker ps -q)" ]; then \
		echo "🛑 Stopping all running Docker containers..."; \
		docker stop $$(docker ps -q); \
		echo "🗑️ Removing all stopped containers..."; \
		docker rm $$(docker ps -aq); \
		echo "✅ All containers stopped and removed."; \
	else \
		echo "⚡ No running containers found. Skipping stop and remove."; \
	fi

## down: stop docker-compose
down:
	echo "🛑 Stopping docker-compose..."
	docker-compose down
	echo "✅ Done!"

## build_logger_service: Builds the logger-service binary as a Linux executable
build_logger_service:
	echo "🚀 Building logger-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd ../logger-service && $$GOFULLPATH build -o $$LOGGER_SERVICE_BINARY ./cmd/api
	echo "✅ Done!"

## logs: Shows logs from all services
logs:
	echo "📜 Fetching last 50 logs for all services..."
	docker-compose logs --tail=50 -f

## help: Displays the list of available commands
help:
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //'
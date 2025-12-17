.PHONY: help build run clean test dependencies kafka-up kafka-down kafka-topics kafka-logs kafka-ui kafka-bridge-info

# Use JDK 21 for all Gradle commands
export JAVA_HOME=/opt/homebrew/opt/openjdk@21/libexec/openjdk.jdk/Contents/Home

# Default target
help:
	@echo "KTable Demo - Available Commands"
	@echo "================================="
	@echo "  make build          - Build the project"
	@echo "  make run            - Run the application"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make test           - Run tests"
	@echo "  make dependencies   - Show dependency tree"
	@echo ""
	@echo "Kafka Management:"
	@echo "  make kafka-up       - Start Kafka + Bridge + UI (Docker Compose)"
	@echo "  make kafka-down     - Stop all services"
	@echo "  make kafka-topics   - List Kafka topics"
	@echo "  make kafka-logs     - Show Kafka logs"
	@echo "  make kafka-ui       - Open Kafbat UI in browser"
	@echo "  make kafka-bridge-info - Show Kafka Bridge HTTP endpoints"

# Gradle tasks
build:
	gradle build

run:
	gradle run

clean:
	gradle clean

test:
	gradle test

dependencies:
	gradle dependencies

# Kafka tasks (we'll create docker-compose.yml next)
run_infra:
	docker-compose up

kafka-topics:
	docker exec -it kafka kafka-topics --bootstrap-server localhost:9092 --list

kafka-logs:
	docker-compose logs -f kafka

kafka-ui:
	@echo "Opening Kafbat UI at http://localhost:8081"
	@open http://localhost:8081 2>/dev/null || xdg-open http://localhost:8081 2>/dev/null || echo "Please open http://localhost:8081 in your browser"

kafka-bridge-info:
	@echo "Kafka Bridge HTTP API available at: http://localhost:8080"
	@echo ""
	@echo "Example endpoints:"
	@echo "  GET  /topics              - List topics"
	@echo "  GET  /topics/{name}       - Get topic details"
	@echo "  POST /topics/{name}       - Send messages"
	@echo "  POST /consumers/{group}   - Create consumer"
	@echo ""
	@echo "Full API docs: https://strimzi.io/docs/bridge/latest/"

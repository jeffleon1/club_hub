FRANCHISE_BINARY=franchiseApp
METADATA_BINARY=metadataApp

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

up_build: build_metadata build_franchise
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

build_metadata:
	@echo "Building broker binary..."
	cd ./metadata && env GOOS=linux CGO_ENABLED=0 go build -o ${METADATA_BINARY} ./cmd
	@echo "Done!"

build_franchise:
	@echo "Building broker binary..."
	cd ./franchise && env GOOS=linux CGO_ENABLED=0 go build -o ${FRANCHISE_BINARY} ./cmd
	@echo "Done!"

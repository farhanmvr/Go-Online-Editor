# Frontend

frontend-build:
	docker build -t go-editor-frontend ./frontend

frontend-start:
	make frontend-build
	docker-compose up -d go-editor-frontend

frontend-stop:
	docker-compose stop go-editor-frontend

# Backend

backend-build:
	docker build -t go-editor-backend ./backend

backend-start:
	make backend-build
	docker rm -f go-editor-backend
	docker-compose up -d go-editor-backend
	docker logs -f go-editor-backend

backend-stop:
	docker-compose stop go-editor-backend

# All

# Start frontend,backend,db containers
start:
	make backend-build
	make frontend-build
	docker-compose down
	docker-compose up -d
	@echo "###############################################################"
	@echo "## Click the link to view the website: http://localhost:3333 ##"
	@echo "###############################################################"

stop:
	docker-compose down

# Clean up old images

cleanup:
	docker image prune -f

.PHONY: start, cleanup, stop, backend-start, backend-stop, backend-build, frontend-build, frontend-start, frontend-stop
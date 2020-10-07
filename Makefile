.PHONY: run-backend
run-backend:
	docker-compose up -d

.PHONY: rebuild
rebuild:
	docker-compose up -d --build

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: update
update:
	go get -u ...
	go mod tidy
	go mod verify

.PHONY: install-frontend
install-frontend:
	(cd ./frontend && npm install)

.PHONY: build-frontend
build-frontend: install-frontend
	(cd ./frontend && npm run build)

.PHONY: start
start: build-frontend run-backend

.PHONY: stop
stop:
	docker-compose down
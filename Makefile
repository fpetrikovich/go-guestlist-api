.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune 

.PHONY: bundle
bundle: ## bundles the submission for... submission
	git bundle create guestlist.bundle --all

.PHONY: generate-mocks
generate-mocks:
	mockgen -source pkg/repository/guest_repository_interface.go -destination pkg/repository/mock_guest_repository.go -package repository
	mockgen -source pkg/service/guest_service_interface.go -destination pkg/service/mock_guest_service.go -package service

.PHONY: run-tests
run-tests:
	go test ./pkg/handler ./pkg/service -v
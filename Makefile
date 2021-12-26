# config
GCP_PROJECT := $(shell gcloud config get-value project)
RUN_SERVICE := torasemi-todo-api

# docker
.PHONY: dev
dev:
	docker compose up

.PHONY: migrate
migrate:
	docker compose exec app go run cmd/migration/main.go

# build
build:
	gcloud builds submit --tag gcr.io/$(GCP_PROJECT)/$(RUN_SERVICE)

cloudbuild:
	gcloud builds submit --config cloudbuild.yml

deploy:
	gcloud run deploy $(RUN_SERVICE) --image gcr.io/$(GCP_PROJECT)/$(RUN_SERVICE) --platform managed --region asia-northeast1 --allow-unauthenticated

# config
GCP_PROJECT := $(shell gcloud config get-value project)
RUN_SERVICE := torasemi-todo-api

# docker
.PHONY: dev
dev:
	docker compose up

build:
	gcloud builds submit --tag gcr.io/$(GCP_PROJECT)/$(RUN_SERVICE)

deploy:
	gcloud run deploy $(RUN_SERVICE) --image gcr.io/$(GCP_PROJECT)/$(RUN_SERVICE) --platform managed --region asia-northeast1 --allow-unauthenticated

# gcp project
create-gcp-project:
	@read -p "project name: " PROJECT_NAME; \
	gcloud projects create $${PROJECT_NAME}

project-list:
	gcloud projects list

set-project:
	@read -p "project id: " PROJECT_ID; \
	gcloud config set project $${PROJECT_ID}

show-config:
	gcloud config list

enable:
	gcloud services enable cloudbuild.googleapis.com

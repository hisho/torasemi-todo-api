GCP_PROJECT := $(shell gcloud config get-value project)
INSTANCE_NAME := "torasemi-todo-db"
DB_NAME := "torasemi-todo-db"
TIER := "db-f1-micro"
REGION := "asia-northeast1"
INSTANCE_CONNECTION_NAME := $(GCP_PROJECT):$(REGION):$(INSTANCE_NAME)
RUN_SERVICE := "torasemi-todo-api"

enable:
	gcloud services enable sqladmin.googleapis.com

create-instance:
	gcloud sql instances create $(INSTANCE_NAME) \
	--database-version=MYSQL_8_0 \
	--tier=$(TIER) \
	--region=$(REGION)

create-database:
	gcloud sql databases create $(DB_NAME) \
	--instance=$(INSTANCE_NAME)

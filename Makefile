VER ?= 0.0.1
REPOSITORY := gravitational.io
NAME := django-app

OPS_URL ?= https://opscenter.localhost.localdomain:33009

CONTAINERS := django-controller:$(VER) \
			  django-app:$(VER)

IMPORT_IMAGE_FLAGS := --set-image=django-controller:$(VER) \
					  --set-image=django-app:$(VER)

IMPORT_OPTIONS := --vendor \
				  --ops-url=$(OPS_URL) \
				  --insecure \
				  --repository=$(REPOSITORY) \
				  --name=$(NAME) \
				  --version=$(VER) \
				  --glob=**/*.yaml \
				  --ignore=images \
				  --registry-url=apiserver:5000 \
				  $(IMPORT_IMAGE_FLAGS)

.PHONY: all
all: images

.PHONY: images
images:
	cd images && $(MAKE) -f Makefile VERSION=$(VER)

.PHONY: import
import: images
	-gravity app delete --ops-url=$(OPS_URL) $(REPOSITORY)/$(NAME):$(VER) --force --insecure
	gravity app import $(IMPORT_OPTIONS) .

.PHONY: clean
clean:
	cd images && $(MAKE) clean

.PHONY: dev-push
dev-push: images
	for container in $(CONTAINERS); do \
		docker tag $$container apiserver:5000/$$container ;\
		docker push apiserver:5000/$$container ;\
	done

.PHONY: dev-redeploy
dev-redeploy: dev-clean dev-deploy

.PHONY: dev-deploy
dev-deploy: dev-push
	kubectl create -f dev/bootstrap.yaml

.PHONY: dev-clean
dev-clean:
	-kubectl delete pod/django-app-bootstrap job/stolon-createdb
	-kubectl delete -f resources/django.yaml

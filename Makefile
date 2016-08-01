VER := 0.0.1
REPOSITORY := gravitational.io
NAME := django-app

OPS_URL ?= https://opscenter.localhost.localdomain:33009

CONTAINERS := django-bootstrap:$(VER) \
			  django-uninstall:$(VER)

IMPORT_IMAGE_FLAGS := --set-image=django-bootstrap:$(VER) \
	--set-image=django-uninstall:$(VER)

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

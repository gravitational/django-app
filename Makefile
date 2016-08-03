VER := 0.0.1
REPOSITORY := gravitational.io
NAME := django-app

OPS_URL ?= https://opscenter.localhost.localdomain:33009

CONTAINERS := django-bootstrap:$(VER) \
			  django-uninstall:$(VER) \
			  django-app:$(VER)

IMPORT_IMAGE_FLAGS := --set-image=django-bootstrap:$(VER) \
	--set-image=django-uninstall:$(VER) \
	--set-image=django-app:$(VER) \
	--set-dep=gravitational.io/k8s-onprem:$$(gravity app list --ops-url=$(OPS_URL) --insecure | grep -m 1 k8s-onprem | awk '{print $$3}' | cut -d: -f2 | cut -d, -f1) \
	--set-dep=gravitational.io/stolon-app:$$(gravity app list --ops-url=$(OPS_URL) --insecure | grep -m 1 stolon-app | awk '{print $$3}' | cut -d: -f2 | cut -d, -f1)

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

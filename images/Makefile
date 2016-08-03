VERSION ?= 0.0.1

BOOTSTRAP_TAG := django-bootstrap:$(VERSION)
UNINSTALL_TAG := django-uninstall:$(VERSION)
APP_TAG := django-app:$(VERSION)

.PHONY: all
all: images

.PHONY: images
images: bootstrap \
		uninstall \
		django
	docker tag $(BOOTSTRAP_TAG) django-bootstrap:latest
	docker tag $(UNINSTALL_TAG) django-uninstall:latest
	docker tag $(APP_TAG) django-app:latest

.PHONY: bootstrap
bootstrap:
	docker build -t $(BOOTSTRAP_TAG) $(PWD)/bootstrap

.PHONY: uninstall
uninstall:
	docker build -t $(UNINSTALL_TAG) $(PWD)/uninstall

.PHONY: django
django:
	docker build -t $(APP_TAG) $(PWD)/django

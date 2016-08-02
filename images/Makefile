.PHONY: all
all: images

BOOTSTRAP_TAG := django-bootstrap:$(VERSION)
UNINSTALL_TAG := django-uninstall:$(VERSION)

.PHONY: images
images: bootstrap \
		uninstall
	docker tag $(BOOTSTRAP_TAG) django-bootstrap:latest
	docker tag $(UNINSTALL_TAG) django-uninstall:latest

.PHONY: bootstrap
bootstrap:
	docker build -t $(BOOTSTRAP_TAG) $(PWD)/bootstrap

.PHONY: uninstall
uninstall:
	docker build -t $(UNINSTALL_TAG) $(PWD)/uninstall

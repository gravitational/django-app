VERSION ?= 0.0.1

CONTROLLER_TAG := django-controller:$(VERSION)
APP_TAG := django-app:$(VERSION)

.PHONY: all
all: images

.PHONY: images
images: controller \
		django
	docker tag $(CONTROLLER_TAG) django-controller:latest
	docker tag $(APP_TAG) django-app:latest

CONTROLLER_BINS := controller/bin
$(CONTROLLER_BINS):
	$(MAKE) -C ../tool/djctl
	mkdir $(CONTROLLER_BINS)
	cp ../tool/djctl/bin/djctl $(CONTROLLER_BINS)/

.PHONY: controller
controller: $(CONTROLLER_BINS)
	docker build -t $(CONTROLLER_TAG) $(PWD)/controller

.PHONY: django
django:
	docker build -t $(APP_TAG) $(PWD)/django

.PHONY: clean
clean:
	-rm -rf $(CONTROLLER_BINS)

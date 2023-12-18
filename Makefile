SHELL := /bin/bash
CONTAINER_BUILDER_NAME=sopro-dev-protocompiler-temporal
DIRECTORIES_PKG := pkg
DIRECTORIES= $(DIRECTORIES_PKG)

# Check if "tag" is not empty
ifeq ($(strip $(tag)),)
$(error TAG variable is empty. Please provide a non-empty value for TAG.)
endif

clean-and-create:
	for dir in $(DIRECTORIES); do \
		rm -rf ./$$dir/* ; \
		mkdir -p ./$$dir ; \
	done

compile: clean-and-create
	docker rm ${CONTAINER_BUILDER_NAME} 2>&1 > /dev/null || true
	docker build --build-arg TAG=${tag} -t ${CONTAINER_BUILDER_NAME} .
	docker run --name ${CONTAINER_BUILDER_NAME} ${CONTAINER_BUILDER_NAME}

	# Golang
	docker cp ${CONTAINER_BUILDER_NAME}:/app/${DIRECTORIES_PKG} .

	go mod tidy  2>&1 > /dev/null || true

list:
	git pull --tags
	git for-each-ref --sort=creatordate --format='%(creatordate:short)	%(refname:short)' refs/tags | tail -n 3

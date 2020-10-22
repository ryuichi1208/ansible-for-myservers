WORKDIR := $(shell pwd)
ARG=debug.yml

.PHONY: build
build:
	@export DOCKER_CONTENT_TRUST=0 DOCKER_BUILDKIT=1 && docker image build -t ansible_test .

.PHONY: test
test: build
	@export DOCKER_CONTENT_TRUST=0 DOCKER_BUILDKIT=1 && docker container run -it -v ${WORKDIR}:/root ansible_test ansible-playbook -i inventory/inventory.ini -c local localhost.yml

.PHONY: debug
debug: build
	@export DOCKER_CONTENT_TRUST=0 DOCKER_BUILDKIT=1 && docker container run -it -v ${WORKDIR}:/root ansible_test ansible-playbook -i inventory/inventory.ini -c local debug.yml

.PHONY: ci
ci: build
	@-export DOCKER_CONTENT_TRUST=0 DOCKER_BUILDKIT=1
	@-echo "==== yamllint ====\n"
	@-docker container run -it -v ${WORKDIR}:/root ansible_test yamllint ${ARG}
	@-echo "==== syntax-check ====\n"
	@-docker container run -it -v ${WORKDIR}:/root ansible_test ansible-playbook --syntax-check -vv ${ARG}
	@-echo "==== ansible-lint ====\n"
	@-docker container run -it -v ${WORKDIR}:/root ansible_test ansible-lint -vv ${ARG}

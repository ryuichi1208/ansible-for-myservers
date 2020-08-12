WORKDIR := $(shell pwd)

.PHONY:
test:
	@export DOCKER_CONTENT_TRUST=0 DOCKER_BUILDKIT=1 && docker image build -t ansible_test .
	@export DOCKER_CONTENT_TRUST=0 DOCKER_BUILDKIT=1 && docker container run -it -v ${WORKDIR}:/root ansible_test ansible-playbook -i inventory/inventory.ini -c local localhost.yml

FROM ansible/ansible:centos7

RUN pip install \
	ansible==2.9.12 \
	ansible-lint==4.2.0

WORKDIR /root

CMD ["/bin/bash"]

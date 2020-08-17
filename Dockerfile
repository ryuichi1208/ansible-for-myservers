FROM ansible/ansible:centos7

RUN pip install \
	ansible==2.9.12 \
	ansible-lint==4.2.0 \
	yamllint==1.3.0 \
	&& yum install -y python-netaddr

WORKDIR /root

CMD ["/bin/bash"]

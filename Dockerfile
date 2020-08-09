FROM ansible/ansible:centos7

RUN pip install ansible

WORKDIR /root

CMD ["/bin/bash"]

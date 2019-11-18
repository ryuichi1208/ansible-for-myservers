# ansible-for-myservers
Setup repository for my cute servers.

## Ansible style guide

* 

## Execution method

``` bash
# syntax test
$ ansible-playbook -i inventory/inventory.ini linux_init.yml --syntax-check

# exec test
$ ansible-playbook -i inventory/inventory.ini linux_init.yml --check

# exec
ansible-playbook -i inventory/inventory.ini linux_init.yml
```

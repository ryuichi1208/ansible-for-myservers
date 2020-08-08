![Ansible syntax test](https://github.com/ryuichi1208/ansible-for-myservers/workflows/Ansible%20syntax%20test/badge.svg?branch=master)

## ansible-playbook -i inventory/inventory.ini linux_init.yml ansible-for-myservers
Setup repository for my cute servers.

## Ansible style guide

* For my servers

## Execution method

``` bash
# syntax test
$ ansible-playbook -i inventory/inventory.ini linux_init.yml --syntax-check

# exec test
$ ansible-playbook -i inventory/inventory.ini linux_init.yml --check

# lint
$ ansible-lint

# config check
$ ansible-config dump --changed-only

# exec
$ ansible-playbook -i inventory/inventory.ini linux_init.yml --diff

# step exec
$ ansible-playbook -i inventory/inventory.ini linux_init.yml --step
```

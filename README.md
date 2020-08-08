## Overview

Setup repository for my cute servers.

## How to use

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


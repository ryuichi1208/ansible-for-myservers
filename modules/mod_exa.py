#!/usr/bin/python
# -*- coding: utf-8 -*-

from ansible.module_utils.basic import AnsibleModule

module = AnsibleModule(
    argument_spec=dict()
)

result = dict(
    some_key="some value",
)

module.exit_json(**result)

# static-resource

A Concourse resource for supplying static information.  Written in Go.

## Example

```yaml
resource_types:
- name: static-resource
  type: registry-image
  source:
    repository: logsquaredn/static-resource
    tag: latest

resources:
- name: config
  type: static-resource
  source:
    key: value

jobs:
- name: some-job
  plan:
  ...
  - get: config

  - load_var: key
    file: config/key
  ...
```

## Source Configuration

## Behavior

### `check`

### `in`

### `out`

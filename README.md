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
    interface:
      key: value

jobs:
- name: some-job
  plan:
  ...
  - get: config
    params:
      format: yml

  - load_var: interface
    file: config/interface
  ...
```

## Source Configuration

A map of key-value pairs

## Behavior

### `check`

returns an array with the hash of the source given to it

### `in`

| Parameter | Required | Description                                                                                                                                     |
| ----------| -------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `format`  | no       | the format of the file that the static information should be made available through. One of `json`, `raw`, `trim`, `yaml`, `yml`. Default `raw` |
| `reveal`  | no       | whether or not to reveal the values in the output. Default `false`                                                     |

### `out`

returns the hash of the source given to it

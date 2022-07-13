# Configuration

This folder holds configuration for all applications in this monorepo. The configuration is created on a per-environment basis. Config can be overwritten in the following ways:

## Config File

```yaml
app_name:
  my_variable: "My Value"
```

## Env Variable
```shell
export APP_NAME.MY_VARIABLE="My Value"
```
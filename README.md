# Local AWS (LocalStack)

## Start environment

```shell
docker compose up
```

## Execute on container

### AWS CLI

```shell
docker compose exec awscli aws s3 ls
```

### Node.js

Install dependencies

```shell
docker compose exec node npm ci
```

Run script

```shell
docker compose exec node node s3.mjs
```

### Golang

Install dependencies

```shell
docker compose exec golang go mod vendor
```

Run script

```shell
docker compose exec golang go run ./s3
```

### Terraform

Install dependencies

```shell
docker compose run terraform init
```

Plan

```shell
docker compose run terraform plan
```

Apply

```shell
docker compose run terraform apply
```

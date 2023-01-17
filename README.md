# go-echo-gorm-boilerplate
Go + Echo + Gormを使ったWebアプリケーション


dockerフォルダ直下に以下の内容を記載した.envを作成する。
```
### Paths #################################################
API_CODE_WORKDIR=/src
API_CODE_LOCAL_PATH=../src

### PORT #################################################
API_PORT=9000

### VERSIONS #################################################
GO_VERSION=1.19
```

srcフォルダ直下に以下の内容を記載した.envを作成する。

```
DB_HOST=db
DB_PORT=5432
DB_USER=gouser
DB_NAME=godb
DB_PASS=gopass
```

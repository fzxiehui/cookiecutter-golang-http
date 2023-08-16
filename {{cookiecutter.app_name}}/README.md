# {{cookiecutter.app_name}}

{{cookiecutter.project_short_description}}

## 目录说明

```shell
.
├── AUTHORS.md					# 作者信息
├── cmd									# 启动指令
│   ├── root.go
│   ├── start.go
│   └── version.go
├── config							# 配置文件库
│   └── config.go
├── configs							# 配置文件模板
├── CONTRIBUTING.md
├── Dockerfile
├── go.mod
├── Gopkg.toml
├── log									# 日志库
│   └── log.go
├── main.go
├── Makefile
├── pkg									# 确认对外开放的库
├── README.md
└── version							# 版本信息
    └── version.go
```

## 编译

```console
$ go get
$ make build
$ ./bin/{{cookiecutter.app_name}}
```

### 测试

``make test``

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

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ go get
$ make
$ ./bin/{{cookiecutter.app_name}}
```

### Testing

``make test``

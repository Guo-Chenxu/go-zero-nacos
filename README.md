# go-zero 整合 nacos 配置中心和注册中心

## 实现功能

- [x] 读取nacos配置中心的配置
- [x] 注册到nacos注册中心

## 项目结构

```shell
│  .gitignore
│  go.mod
│  go.sum
│  LICENSE
│  README.md
│
├─configcenter # 配置中心
│      bootstrap_config.go
│      get_config.go
│      nacos.go
│
├─example   # 示例
└─register  # 注册中心
        register.go
        register_config.go
```

## 使用方式

```shell
 go get -u github.com/Guo-Chenxu/go-zero-nacos
```

具体例子请看 [nacos-demo](https://github.com/Guo-Chenxu/go-zero-demo/tree/main/nacos-demo)
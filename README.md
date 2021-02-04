### 项目简介：

日常工作台，暂时无前端。

### 使用方法：

1. 编辑配置文件，在ff/setting/目录下有一份示例配置文件：config-example.yaml，编辑各个配置
2. 检查所需组件运行正常，例如：MySQL、Redis、
3. 在main.go同级目录下执行go build命令
4. 不出意外的话，编译出来的可执行文件名为：ff
5. 执行以下命令进行数据库迁移，生成项目所需表。

```shell
./ff migrate
```

5. 启动服务

```shell
nohup  ./ff webserver start	&
```



### 用到的一些开源框架、库和组件：

- Gin
- GORM
- Swaggo
- Jwt-go
- Casbin
- Cobra
- Viper
- Zap



### 功能模块划分：

#### web server

提供web api，包含的app：

```
bookmark（书签）
memo（备忘录）
天气预报
```



#### CLI工具

利用cobra构建的命令行工具，命令详情可以通过./ff --help查看

命令举例：

```shell
./ff migrate	# 配置文件确认正确后使用，用于生成项目所需表

nohup  ./ff webserver start	& # 启动webserver并后台运行
```

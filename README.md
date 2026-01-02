# Kratos Project Template

## 开发

> 如果已经安装工具或执行了`init`和`all`命令则跳过相关步骤

- 安装 [buf](https://buf.build/docs/cli/installation/) 工具，用来管理 proto
- 安装 [just](https://just.systems/man/zh/%E5%AE%89%E8%A3%85%E5%8C%85.html) 工具，用来编排命令
- 执行命令 `just init`，安装一些开发工具
- 执行命令 `just all`
- 运行项目 `just run`


## 命令

`just help`

```shell
Usage:
  just <recipe>

Recipes:
just --list
Available recipes:
    all      # 执行所有生成任务
    build    # 构建二进制
    check    # 检查代码
    format   # 格式化代码
    generate # 生成代码 & tidy
    help     # 显示帮助
    init     # 安装必要工具
    run      # 运行项目
    update   # 更新依赖
```


## 环境变量

| 名字           | 说明                                            |
|--------------|-----------------------------------------------|
| APP_ENV      | 应用环境: production, staging, development, local |
| PROJECT_NAME | 项目名字                                          |
| PROJECT_REF  | 项目分支                                          |
| PROJECT_SHA  | 项目版本                                          |

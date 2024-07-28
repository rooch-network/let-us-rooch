# 学习日志

## task1

在btc core控制台中使用命令，不能有bitcoin-cli

## task2

一波n折

一开始使用windows编译，遇到各种问题，最后询问官方人员，选择转向wsl

wsl连接vscode

然后一直报错，

```
Unable to resolve packages for package 'hello_rooch'
```

通过翻看仓库的学习笔记找到解决办法

## 第一步

```
rooch server start
```

### 第二步

新开端口

```
先 rooch move build
再 rooch move publish
```


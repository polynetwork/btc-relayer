# BTC Relayer

## 简介

​	BTC Relayer是跨链生态中比特币网络与联盟链之间沟通的桥梁，负责将BTC的跨链交易转发到联盟链上，这是跨链的第一步。它的主要功能包括两部分，首先是监听比特币网络，将要跨链到其他链的BTC交易，转发到联盟链；然后是监听联盟链，将要转回比特币链的交易广播到比特币网络中。

## 使用

​	在项目的cmd文件夹下，进行编译。

```shell
go build -o run_btc_relayer run.go 
```

​	通过下列命令，启动relayer。

```
run_btc_relayer -conf-file=/path/to/conf.json -log-path=/path/to/log/ 
```

​	当然配置conf.json需要自行填写


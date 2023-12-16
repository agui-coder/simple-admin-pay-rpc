# simple-admin-pay-rpc

# 模块介绍

RPC模块提供基本支付服务。

#### 软件架构
grpc

### 框架

[go-zero](https://github.com/tal-tech/go-zero): 微服务框架

[ent](https://github.com/ent/ent): 是一个简单而又功能强大的Go语言实体框架，ent易于构建和维护应用程序与大数据模型。

#### 使用说明

[pay.yaml](./etc/pay.yaml) 配置文件填好配置运行
 调用 [InitDatabase](./internal/logic/base/init_database_logic.go) 方法初始化数据库

```yaml

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
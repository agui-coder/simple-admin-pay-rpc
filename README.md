---

## 简介

Simple Admin Pay 是一个支付服务，基于Simple Admin开发，适合用于微服务学习，开源免费。

## 调试方式

1. 通过swagger进行测试


    进入服务"./api"目录下生成swagger
    swagger generate spec --output=./swagger.yml --scan-models
    运行 swagger
    swagger serve --no-open -F=swagger --port 36666 swagger.yml

2. 通过logic下的http文件调试

[@Kevin](https://gitee.com/agui-coder)

## License

[MIT © Kevin-2023](./LICENSE)

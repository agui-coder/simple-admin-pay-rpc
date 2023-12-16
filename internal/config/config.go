package config

import (
	"errors"
	"github.com/agui-coder/simple-admin-pay-rpc/payment"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/mq/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"io"
	"os"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf     config.DatabaseConf
	RedisConf        redis.RedisConf
	AsynqConf        asynq.AsynqConf
	PayProperties    model.Properties
	AliPayConfigPath AliPayConfigPath `json:",optional"`
	WxPayConfigPath  WxPayConfigPath  `json:",optional"`
}

type AliPayConfigPath struct {
	AppId                       string // appId
	SignType                    string // 签名类型
	PrivateKey                  string // 应用私钥
	Status                      bool   // 是否启用
	AppPublicContentPath        string // 应用公钥证书内容
	AlipayPublicContentRSA2Path string // 支付宝公钥证书内容
	AlipayRootContentPath       string // 支付宝根证书内容
}

type WxPayConfigPath struct {
	AppId                 string //appId
	MchId                 string
	SerialNumber          string //apiclient_cert.pem 证书文件的证书号
	ApiV3Key              string
	Status                bool   // 是否启用
	PrivateKeyContentPath string //apiclient_key.pem 证书文件的对应字符串
}

func (p *AliPayConfigPath) NewAliPayPayConfig(payConfig *payment.PayConfig) error {
	if p.Status {
		payConfig.AliConfig.Status = true
		payConfig.AliConfig.AppId = p.AppId
		payConfig.AliConfig.SignType = p.SignType
		payConfig.AliConfig.PrivateKey = p.PrivateKey
		appPublicContent, err := readFile(p.AppPublicContentPath)
		if err != nil {
			return err
		}
		payConfig.AliConfig.AppPublicContent = appPublicContent
		alipayPublicContentRSA2, err := readFile(p.AlipayPublicContentRSA2Path)
		if err != nil {
			return err
		}
		payConfig.AliConfig.AlipayPublicContentRSA2 = alipayPublicContentRSA2
		alipayRootContent, err := readFile(p.AlipayRootContentPath)
		if err != nil {
			return err
		}
		payConfig.AliConfig.AlipayRootContent = alipayRootContent
	}
	return nil
}

func (w *WxPayConfigPath) NewWxPayPayConfig(payConfig *payment.PayConfig) error {
	if w.Status {
		payConfig.WxConfig.Status = true
		payConfig.WxConfig.AppId = w.AppId
		payConfig.WxConfig.MchId = w.MchId
		payConfig.WxConfig.SerialNumber = w.SerialNumber
		payConfig.WxConfig.ApiV3Key = w.ApiV3Key
		privateKeyContent, err := readFile(w.PrivateKeyContentPath)
		if err != nil {
			return err
		}
		payConfig.WxConfig.PrivateKeyContent = privateKeyContent
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("file does not exist: " + path)
	}

	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file: " + path)
	}
	defer file.Close()

	// 读取文件内容
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read file: " + path)
	}
	return bytes, nil
}

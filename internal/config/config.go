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
	DatabaseConf  config.DatabaseConf
	RedisConf     redis.RedisConf
	AsynqConf     asynq.AsynqConf
	PayProperties model.Properties
	AliPayConfig  AliPayConfig `json:",optional"`
	WxPayConfig   WxPayConfig  `json:",optional"`
}

type AliPayConfig struct {
	AppId                       string `json:",env=ALI_PAY_APPID"`                                                // appId
	SignType                    string `json:",env=ALI_PAY_SIGN_TYPE"`                                            // 签名类型
	PrivateKey                  string `json:",env=ALI_PAY_PRIVATE_KEY"`                                          // 应用私钥
	Status                      bool   `json:",default=false,env=ALI_PAY_STATUS"`                                 // 是否启用
	AppPublicContentPath        string `json:",default=cert/alipay/appPublicCert.crt,env=ALIPAY_APP_PUBLIC_CERT"` // 应用公钥证书内容
	AlipayPublicContentRSA2Path string `json:",default=cert/alipay/alipayPublicCert.crt,env=ALIPAY_PUBLIC_CERT"`  // 支付宝公钥证书内容
	AlipayRootContentPath       string `json:",default=cert/alipay/alipayPublicCert.crt,env=ALIPAY_PUBLIC_CERT"`  // 支付宝根证书内容
}

type WxPayConfig struct {
	AppId                 string `json:",env=WX_PAY_APPID"`                                               //appId
	MchId                 string `json:",env=WX_PAY_MCHID"`                                               //商户号
	SerialNumber          string `json:",env=WX_PAY_SERIAL_NUMBER"`                                       //apiclient_cert.pem 证书文件的证书号
	ApiV3Key              string `json:",env=WX_PAY_APIV3KEY"`                                            //apiclient_key.pem 证书文件的证书号
	Status                bool   `json:",default=false,env=WX_PAY_STATUS"`                                // 是否启用
	PrivateKeyContentPath string `json:",default=cert/wechat/apiclient_key.pem,env=WX_PAY_APICLIENT_KEY"` //apiclient_key.pem 证书文件的对应字符串
}

func (p *AliPayConfig) NewAliPayPayConfig(payConfig *payment.PayConfig) error {
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

func (w *WxPayConfig) NewWxPayPayConfig(payConfig *payment.PayConfig) error {
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 读取文件内容
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read file: " + path)
	}
	return bytes, nil
}

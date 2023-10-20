package register

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

// Register 注册到 nacos
func Register(config any) {
	// 先序列化再反序列化转换类型
	var c Config
	jsonByte, _ := json.Marshal(config)
	_ = json.Unmarshal(jsonByte, &c)

	// 读取配置并注册
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.NacosConfig.Host, c.NacosConfig.Port),
	}

	cc := &constant.ClientConfig{
		Username:            c.NacosConfig.Username,
		Password:            c.NacosConfig.Password,
		NamespaceId:         c.NacosConfig.NamespaceId,
		TimeoutMs:           c.NacosConfig.TimeoutMs,
		NotLoadCacheAtStart: c.NacosConfig.NotLoadCacheAtStart,
		LogLevel:            c.NacosConfig.LogLevel,
	}

	opts := nacos.NewNacosConfig(c.Name, c.ListenOn, sc, cc)
	err := nacos.RegisterService(opts)
	if err != nil {
		panic(err)
	}
}

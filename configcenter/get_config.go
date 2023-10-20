package configcenter

import "github.com/zeromicro/go-zero/core/conf"

// GetConfig 从nacos配置中心中解析配置
func GetConfig(bootstrapConfigFile *string, bootstrapConfig *BootstrapConfig, c any) {
	// 加载 bootstrap 配置
	conf.MustLoad(*bootstrapConfigFile, bootstrapConfig)

	// 解析配置中心的配置
	nn := NewNacos(bootstrapConfig.NacosConfig)
	serviceConfigContent := nn.InitConfig(func(data string) {
		err := conf.LoadFromYamlBytes([]byte(data), c)
		if err != nil {
			panic(err)
		}
	})
	err := conf.LoadFromYamlBytes([]byte(serviceConfigContent), c)
	if err != nil {
		panic(err)
	}
}

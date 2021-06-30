package vo

import "github.com/yule526751/nacos-client/common/constant"

type NacosClientParam struct {
	ClientConfig  *constant.ClientConfig  // optional
	ServerConfigs []constant.ServerConfig // optional
}

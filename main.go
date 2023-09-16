package main

import (
	"hadebbs/app/console"
	"hadebbs/app/http"
	"hadebbs/framework"
	"hadebbs/framework/provider/app"
	"hadebbs/framework/provider/cache"
	"hadebbs/framework/provider/config"
	"hadebbs/framework/provider/distributed"
	"hadebbs/framework/provider/env"
	"hadebbs/framework/provider/id"
	"hadebbs/framework/provider/kernel"
	"hadebbs/framework/provider/log"
	"hadebbs/framework/provider/orm"
	"hadebbs/framework/provider/redis"
	"hadebbs/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&id.HadeIDProvider{})
	container.Bind(&trace.HadeTraceProvider{})
	container.Bind(&log.HadeLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.HadeCacheProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}

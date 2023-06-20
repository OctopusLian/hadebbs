package main

import "hadebbs/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}

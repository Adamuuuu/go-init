package router

import "go-init/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

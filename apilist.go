package main

import (
	"git.shopex.cn/luban/lib/apipkg"
	"github.com/shopex/luban-srv-hello/controller"
)

var Apis []interface{}

func init() {
	Apis = apipkg.ApiList(
		apipkg.ApiInclude(&controller.Item{}),
	)
}

package controller

import (
	"fmt"
	"git.shopex.cn/luban/lib"
	"time"
)

type Item struct {
}

// @Title 更新
// @Doc   填写name
// @Param  name     string  false  Name
func (me *Item) Say(args lib.Args) (string, error) {
	ret := fmt.Sprintf("hi, %s, timestamp is %d\n", args.Str("name"), time.Now().Unix())
	return ret, nil
}

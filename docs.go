package main

import (
	"encoding/json"

	"git.shopex.cn/luban/lib/apipkg"
)

const (
    api_string string = `{
    "apis": [
        {
            "name": "item_say",
            "title": "更新",
            "summary": "<p>填写name</p>\n",
            "parameters": [
                {
                    "name": "name",
                    "description": "Name",
                    "type": "string",
                    "format": "",
                    "allowMultiple": false,
                    "required": false,
                    "minimum": 0,
                    "maximum": 0
                }
            ]
        }
    ]
}
`
)

var ApiDefine apipkg.ApiDeclaration

func init() {
	json.Unmarshal([]byte(api_string), &ApiDefine)
}

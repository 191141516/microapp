// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package auth 登录
package auth

import (
	"bytes"
	"github.com/fastwego/microapp"
)

const (
	apiJsCode2Session = "/api/apps/v2/jscode2session"
)

/*
code2Session

通过login接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session

POST https://developer.toutiao.com/api/apps/v2/jscode2session
*/
func Code2Session(ctx *microapp.MicroApp, params []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiJsCode2Session, bytes.NewReader(params), "application/json;charset=utf-8")
}

// Package code 错误代码，异常码等等
package code

import "fmt"

type ECode struct {
	Code    int    // 错误码
	Message string // 错误信息
	Err     error  // 原生错误
}

// 固定错误码固定信息

var InvalidParam = ECode{Code: 499, Message: "无效的参数"}

// 固定错误码动态信息

var ErrParam = func(k string, v interface{}) ECode {
	return ECode{Code: 499, Message: fmt.Sprintf("参数错误:%s=%v", k, v)}
}

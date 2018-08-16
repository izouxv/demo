package common

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type reply struct {
	Code   int32       `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}


func getJson(code string) reply {
	return CodeMap[code]
}

var CodeMap = map[string]reply{
	"OK":					{10000, "成功", 			 		nil},
	"ERR":					{10001, "系统异常", 					nil},
	"Unknown":				{28001, "未知错误", 					nil},
	"InvalidArgument":		{28002, "非法参数", 					nil},
	"NotFound":				{28003, "未找到", 					nil},
	"AlreadyExists":		{28004, "已经存在",					nil},
	"Unavailable":			{28005, "资源不可用", 				nil},
	"Unimplemented":		{28006, "未能实现", 					nil},
	"PAUSE":				{28007, "接口暂时停用", 				nil},
	"token_is_invalid":		{28008, "token错误", 				nil},
	"User_Kick_out":		{28009, "用户被踢出", 				nil},
	"PARAMS_ERR":			{28010, "参数有误",					nil},
	"Canceled":				{28011, "注销",						nil},
	"DeadlineExceeded":		{28012, "资源过期", 					nil},
	"permission denied":	{28013, "权限不足，拒绝访问",		nil},
	"ResourceExhausted":	{28014, "无资源", 					nil},
	"FailedPrecondition":	{28015, "资源枯竭", 					nil},
	"Aborted":				{28016, "操作取消", 					nil},
	"OutOfRange":			{28017, "超出范围",					nil},
	"Internal":				{28018, "内部的", 					nil},
	"DataLoss":				{28019, "数据丢失", 					nil},
	"UNAUTHENTICATED":		{28020, "未经证实的", 				nil},
	"Email_failed":		    {28021, "邮件发送器不存在", 				nil},
}

func JsonReply(msg string, result interface{}, w http.ResponseWriter) {
	r := getJson(msg)
	r.Result = result
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	value, _ := json.Marshal(r)
	fmt.Fprintf(w, "%s", value)
	return
}


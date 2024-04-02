// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAliyunSms interface {
		Send(ctx context.Context, mobile string, verifyCode string) (err error)
	}
)

var (
	localAliyunSms IAliyunSms
)

func AliyunSms() IAliyunSms {
	if localAliyunSms == nil {
		panic("implement not found for interface IAliyunSms, forgot register?")
	}
	return localAliyunSms
}

func RegisterAliyunSms(i IAliyunSms) {
	localAliyunSms = i
}

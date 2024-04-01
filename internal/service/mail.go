// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IMail interface {
		Send(ctx context.Context, to, subject, html string) (err error)
	}
)

var (
	localMail IMail
)

func Mail() IMail {
	if localMail == nil {
		panic("implement not found for interface IMail, forgot register?")
	}
	return localMail
}

func RegisterMail(i IMail) {
	localMail = i
}

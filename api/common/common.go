// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"gohub/api/common/v1"
)

type ICommonV1 interface {
	GetVerifyCodeByCaptcha(ctx context.Context, req *v1.GetVerifyCodeByCaptchaReq) (res *v1.GetVerifyCodeByCaptchaRes, err error)
	GetVerifyCodeByEmail(ctx context.Context, req *v1.GetVerifyCodeByEmailReq) (res *v1.GetVerifyCodeByEmailRes, err error)
	GetVerifyCodeByPhone(ctx context.Context, req *v1.GetVerifyCodeByPhoneReq) (res *v1.GetVerifyCodeByPhoneRes, err error)
}

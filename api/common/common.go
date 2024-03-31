// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"gohub/api/common/v1"
)

type ICommonV1 interface {
	GetCaptchaOne(ctx context.Context, req *v1.GetCaptchaOneReq) (res *v1.GetCaptchaOneRes, err error)
}

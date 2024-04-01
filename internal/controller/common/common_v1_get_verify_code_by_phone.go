package common

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gohub/api/common/v1"
)

func (c *ControllerV1) GetVerifyCodeByPhone(ctx context.Context, req *v1.GetVerifyCodeByPhoneReq) (res *v1.GetVerifyCodeByPhoneRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

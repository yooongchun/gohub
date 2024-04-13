package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gohub/internal/service"

	"gohub/api/user/v1"
)

func (c *ControllerV1) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	err = service.GhToken().RemoveToken(ctx, service.GhToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}

package captcha

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
	"gohub/internal/app/common/service"
)

func init() {
	service.RegisterCaptcha(NewCaptcha())
}

type sCaptcha struct {
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
}

func NewCaptcha() *sCaptcha {
	return &sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      180,
			ShowLineOptions: 2 | 4,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"wqy-microhei.ttc"},
		},
		store: base64Captcha.DefaultMemStore,
	}
}

// GetVerifyImgString 生成验证码图片
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKeyC string, base64StringC string, err error) {
	driver := s.driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.store)
	idKeyC, base64StringC, _, err = c.Generate()
	return
}

// VerifyCaptcha 验证验证码
func (s *sCaptcha) VerifyCaptcha(idKey string, verifyValue string) bool {
	c := base64Captcha.NewCaptcha(s.driver, s.store)
	return c.Verify(idKey, gstr.ToLower(verifyValue), true)
}

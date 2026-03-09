package env

import (
	"os"

	"github.com/joho/godotenv"
)

var AdminPassword string
var AllowOrigin string

func init() {
	// .env 파일 없으면 무시
	_ = godotenv.Load()

	// 환경변수 기본값
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	if AdminPassword == "" {
		AdminPassword = "dummy-password"
	}

	AllowOrigin = os.Getenv("ALLOW_ORIGIN")
	if AllowOrigin == "" {
		AllowOrigin = "*"  // 모든 프론트 허용
	}
}

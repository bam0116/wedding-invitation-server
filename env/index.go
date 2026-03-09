package env

import (
	//"fmt"
	"os"

	"github.com/joho/godotenv"
)

var AdminPassword string
var AllowOrigin string

func init() {
	// .env 파일 읽기 시도, 실패해도 그냥 넘어감
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using defaults")
	}

	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	if AdminPassword == "" {
		AdminPassword = "6692" // 기본값
	}

	AllowOrigin = os.Getenv("ALLOW_ORIGIN")
	if AllowOrigin == "" {
		AllowOrigin = "https://bam0116.github.io/wedding-invitation" // 기본 CORS 허용 origin
	}
}
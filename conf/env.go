package conf

import (
	"github.com/joho/godotenv"
	"github.com/ponycool/nebula-lib/file"
	"os"
)

// Load 载入ENV配置
func Load() {
	env := os.Getenv("ENV")
	if "" == env {
		env = "dev"
	}

	pwd, _ := os.Getwd()
	if file.Exists(pwd + string(os.PathSeparator) + ".env." + env) {
		_ = godotenv.Load(".env." + env)
	}

	_ = godotenv.Load()
}

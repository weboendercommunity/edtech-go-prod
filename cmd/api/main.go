package main

import (
	"github.com/gin-gonic/gin"

	mysql "edtech.id/pkg/db/mysql"

	oauth "edtech.id/internal/oauth/injector"
	profile "edtech.id/internal/profile/injector"
	register "edtech.id/internal/register/injector"
)

func main() {
	db := mysql.DB()

	r := gin.Default()

	// wire
	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	r.Run("127.0.0.1:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

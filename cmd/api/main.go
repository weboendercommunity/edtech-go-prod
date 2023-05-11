package main

import (
	"github.com/gin-gonic/gin"

	mysql "edtech.id/pkg/db/mysql"

	admin "edtech.id/internal/admin/injector"
	oauth "edtech.id/internal/oauth/injector"
	product "edtech.id/internal/product/injector"
	productCategory "edtech.id/internal/product_category/injector"
	profile "edtech.id/internal/profile/injector"
	register "edtech.id/internal/register/injector"
)

func main() {
	db := mysql.DB()

	r := gin.Default()

	// wire
	admin.InitializeAdminHandler(db).Route(&r.RouterGroup)
	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)
	product.InitializedService(db).Route(&r.RouterGroup)

	r.Run("127.0.0.1:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

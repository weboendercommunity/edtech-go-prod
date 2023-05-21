package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	mysql "edtech.id/pkg/db/mysql"

	admin "edtech.id/internal/admin/injector"
	cart "edtech.id/internal/cart/injector"
	classRoom "edtech.id/internal/class_room/injector"
	discount "edtech.id/internal/discount/injector"
	oauth "edtech.id/internal/oauth/injector"
	order "edtech.id/internal/order/injector"
	product "edtech.id/internal/product/injector"
	productCategory "edtech.id/internal/product_category/injector"
	profile "edtech.id/internal/profile/injector"
	register "edtech.id/internal/register/injector"
	webhook "edtech.id/internal/webhook/injector"
)

func init() {
	pathdir, _ := os.Getwd()
	environment := godotenv.Load(filepath.Join(pathdir, ".env"))

	if environment != nil {
		panic(environment)
	}
}

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
	cart.InitializeService(db).Route(&r.RouterGroup)
	discount.InitializeService(db).Route(&r.RouterGroup)
	order.InitializeService(db).Route(&r.RouterGroup)
	classRoom.InitializeService(db).Route(&r.RouterGroup)
	webhook.InitializedService(db).Route(&r.RouterGroup)

	r.Run("127.0.0.1:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

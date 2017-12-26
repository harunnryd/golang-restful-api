package routing

import (
	"golang-restful-api/models"
	"golang-restful-api/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type WebService struct {}

func (w *WebService) Run() {
	/*
		----------------------------------------
		NOTE : sesuaikan string conn ini dengan
		nama database yang digunakan ..
		----------------------------------------
	*/
	conn := "root:@tcp(127.0.0.1:3306)/apitest?parseTime=true"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}

	/*
		-----------------------------------
		NOTE : jika ingin menambahkan model
		yang ingin di migrate, cukup tambah
		kan parameter berupa struct (model)
		-----------------------------------
	*/
	db.AutoMigrate(&models.Pet {})

	w.routing(db)
}

/*
	-------------------------------
	NOTE : tambah route di sini ..
	-----------------------------
*/
func (w *WebService) routing(db *gorm.DB) {
	/*
		-----------------------------------
		NOTE : anggap saja controller  :'(
		nameController struct {}
		-----------------------------------
	*/
	petController := controller.PetController {DB: db}

	r := gin.Default()
	v1 := r.Group("api/v1/pets")
	{
		v1.GET("/", petController.Index)
		v1.POST("/", petController.Create)
		v1.GET("/:id", petController.Show)
		v1.PUT("/:id", petController.Update)
		v1.PUT("/:id/upload", petController.UploadImage)
	}
	r.Run()
}


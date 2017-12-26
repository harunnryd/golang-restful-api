package controller

import (
	"strconv"
	"fmt"
	"time"
	"golang-restful-api/models"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

type PetController struct {
	DB *gorm.DB
}

func (p *PetController) Index(c *gin.Context) {
	var pets []models.PetResult
	p.DB.Table("pets").Find(&pets)

	if len(pets) <= 0 {
		c.JSON(404, gin.H {"status": 404, "message": "not found."})
		return
	}

	c.JSON(200, gin.H {"status": 200, "data": pets})
}

func (p *PetController) Create(c *gin.Context) {
	var pet models.Pet
	err := c.Bind(&pet)

	if err != nil {
		c.JSON(404, gin.H {"error": err.Error()})
		return
	}

	p.DB.Save(&pet)
	c.JSON(200, gin.H {"status": 200, "message": "pet item created."})
}

func (p *PetController) Show(c *gin.Context) {
	var pet models.PetResult
	petID := c.Param("id")
	p.DB.Table("pets").First(&pet, petID)

	if pet.Id == 0 {
		c.JSON(404, gin.H {"status": 404, "message": "pet not found."})
		return
	}

	c.JSON(200, gin.H {"status": 200, "data": pet})
}

func (p *PetController) Update(c *gin.Context) {
	var pet models.Pet
	petID := c.Param("id")

	var existingPet models.Pet

	if p.DB.First(&existingPet, petID).RecordNotFound() {
		c.JSON(404, gin.H {"status": 404, "message": "record not found."})
		return
	}

	p.DB.First(&pet, petID)
	pet.Name = c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	pet.Age = age
	pet.Photo = c.PostForm("photo")
	p.DB.Save(&pet)
	c.JSON(200, gin.H {"status": 200, "message": "pet item updated."})

}

func (p *PetController) UploadImage(c *gin.Context) {
	var fileLocation = "./tmp/"
	var filename string
	var pet models.Pet

	img, err := imageupload.Process(c.Request, "photo")

	if err != nil {
		panic(err)
	}

	if thumb, err := imageupload.ThumbnailPNG(img, 300, 300); err != nil {
	} else {
		filename = fmt.Sprintf("%d.png", time.Now().Unix())
		err = thumb.Save(fileLocation + filename)

		if err != nil {
			panic(err)
		}
	}

	if thumb, err := imageupload.ThumbnailJPEG(img, 300, 300, 80); err != nil {

	} else {
		filename = fmt.Sprintf("%d.jpeg", time.Now().Unix())
		err = thumb.Save(fileLocation + filename)

		if err != nil {
			panic(err)
		}
	}

	petID := c.Param("id")

	var existingPet models.Pet

	if p.DB.First(&existingPet, petID).RecordNotFound() {
		c.JSON(404, gin.H {"status": 404, "message": "record not found."})
	}

	p.DB.First(&pet, petID)
	pet.Photo = filename
	p.DB.Save(&pet)

	c.JSON(200, gin.H {"status": 200, "message": "photo item updated."})
}

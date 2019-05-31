package controllers

import (
	"catmash/models"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

/*NewCat is controller for add cat with image uploading  */
func NewCat(c echo.Context) error {

	//Get Image
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("upload/img/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//set Model
	var m models.Cat

	//bind data to model
	if err := c.Bind(&m); err != nil {
		return c.String(422, err.Error())
	}

	//generate id/date
	m.ID = bson.NewObjectId()
	m.Created = time.Now()
	m.Img = "/img/" + file.Filename

	//insert to model to database
	if err := Dao.InsertCat(m); err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(http.StatusOK, m)
}

/*VoteUp increase vote score on selected cat  */
func VoteUp(c echo.Context) error {
	//get cat
	cat, err := Dao.FindCatByID(c.Param("id"))
	if err != nil {
		return c.String(500, err.Error())
	}

	//upgrade scoring
	cat.Vote = cat.Vote + 1

	//update model into database
	if err := Dao.UpdateCat(cat); err != nil {
		return c.String(500, err.Error())
	}
	//return success
	return c.String(200, "Voted")
}

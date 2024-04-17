package api

import (
	"gin-api/api/response"
	"gin-api/pkg/app"
	"gin-api/pkg/e"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	appG := response.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": "pong",
		"total": "1",
	})
}

type PongForm struct {
	TagID     int        `form:"tag_id" json:"tag_id" uri:"tag_id" binding:"required"`
	User      string     `form:"user" json:"user" binding:"required,min=3"`
	Password  string     `form:"password" json:"password" binding:"required,max=5"`
	TimeStart app.MyTime `form:"time_start" json:"time_start" binding:"required"`
}

func AddPing(c *gin.Context) {
	var json PongForm
	//appG := response.Gin{C: c}
	//log.Println("jam server : ", time.Now().Format("2006-01-02 15:04:05"))
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println("waktu", json.TimeStart)
		//appG.Response(http.StatusBadRequest, e.UNKNOWN, nil)
		c.JSON(http.StatusOK, gin.H{"status": err.Error()})
		return
	}

	log.Println("waktu 1", time.Time(json.TimeStart).Format("2006-01-02 15:04:05"))
	log.Println("waktu 2", json.TimeStart)
	t := time.Now()
	log.Println("waktu 3", t.String())
	log.Println("waktu 4", t.Format("2006-01-02 15:04:05"))
	c.JSON(http.StatusOK, gin.H{"status": "status succes"})

}

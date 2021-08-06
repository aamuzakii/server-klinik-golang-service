package main

import (
	"net/http"
	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
 	"github.com/astaxie/beego/orm"
	 "fmt"
)

var ORM orm.Ormer

func init() {
    models.ConnectToDb()
    ORM = models.GetOrmObject()
}


func main() {
	router := gin.Default()
	router.POST("/createUser", createUser)
	router.GET("/readUsers", readUsers)
	// router.PUT("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)
	router.Run(":3000")
}




func readUsers(c *gin.Context) {
	var user []models.Users
	fmt.Println(ORM)
	_, err := ORM.QueryTable("users").All(&user)
	if(err == nil) {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": &user})
	} else {
			c.JSON(http.StatusInternalServerError, 
					gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the users"})
	}
}

func createUser(c *gin.Context) {
	var newUser models.Users
	c.BindJSON(&newUser)
	_, err := ORM.Insert(&newUser)
	if err == nil {
			c.JSON(http.StatusOK, gin.H{
					"status": http.StatusOK, 
					"email": newUser.Email,
					"user_name": newUser.UserName, 
					"user_id": newUser.UserId})
	} else {
			c.JSON(http.StatusInternalServerError, 
					gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
	} 
}

func deleteUser(c *gin.Context) {
	var delUser models.Users
	c.BindJSON(&delUser)
	_, err := ORM.QueryTable("users").Filter("email", delUser.Email).Delete()
	if(err == nil) {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
			c.JSON(http.StatusInternalServerError, 
					gin.H{"status": http.StatusInternalServerError, "error": "Failed to delete the users"})
	}
}
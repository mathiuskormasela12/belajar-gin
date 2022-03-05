package main

import "github.com/gin-gonic/gin"

func main() {
	var router = gin.Default()

	// mengelompokan router
	var api = router.Group("/api/v1")
	{
		api.GET("/users", getUsers)
		api.POST("/users", postUser)

		var userApi = api.Group("/users")
		{
			userApi.PUT("/:nama", editUsers)
		}
	}

	router.Run()
}

func getUsers(context *gin.Context) {
	context.JSON(200, gin.H{
		"status":  200,
		"message": "Berhasil mengambil data users",
	})
}

func postUser(context *gin.Context) {
	// untuk mengambil req.body seperti di express ini hanya unutk www-url-encoded (json gk bisa)
	username := context.PostForm("username")
	password := context.PostForm("password")

	context.JSON(200, gin.H{
		"status":   200,
		"message":  "Berhasil post data",
		"username": username,
		"password": password,
	})
}

func editUsers(context *gin.Context) {
	// mengambil req.params seperti di express
	var nama = context.Param("nama")

	context.JSON(200, gin.H{
		"status":  200,
		"message": "Halo nama saya " + nama,
	})
}

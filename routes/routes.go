package routes

import (
	"golang-social-media/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// User routes
	r.POST("/users", controllers.RegisterUser)     			// Registrasi user baru
	r.GET("/users", controllers.GetUsers)          			// Ambil semua user
	r.GET("/users/:id", controllers.GetUserByID)   			// Ambil profile user
	r.PUT("/users/:id", controllers.UpdateUser)    			// update profile user
	r.DELETE("/users/:id", controllers.DeleteUser) 			// Hapus user

	// Post routes
	r.POST("/posts", controllers.CreatePost)                	// Buat post baru
	r.GET("/posts", controllers.GetPosts)                   	// Ambil semua post
	r.GET("/posts/:id", controllers.GetPostByID)            	// Ambil detail post
	r.GET("/users/:id/posts", controllers.GetPostsByUserID) 	// Ambil semua post berdasarkan user
	r.DELETE("/posts/:id", controllers.DeletePost)          	// Hapus post

	// Like routes
	r.POST("/likes", controllers.LikePost)                  	// Suka post
	r.GET("/likes", controllers.GetLikes)                   	// Lihat semua like
	r.GET("/posts/:id/likes", controllers.GetLikesByPostID) 	// Lihat siapa saja yang like post tertentu
	r.GET("/users/:id/likes", controllers.GetLikesByUserID) 	// Lihat semua like dari seorang user
	r.DELETE("/likes/:id", controllers.UnlikePost)          	// Batalkan like pada post

	// Comment routes
	r.POST("/comments", controllers.CreateComment)                	// Tambah komentar pada post
	r.GET("/comments", controllers.GetComments)			// Lihat semua komentar
	r.GET("/posts/:id/comments", controllers.GetCommentsByPostID) 	// Lihat komentar pada post tertentu
	r.DELETE("/comments/:id", controllers.DeleteComment)          	// Hapus komentar
}

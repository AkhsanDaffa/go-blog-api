package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"

	"go-blog-api/config"
	"go-blog-api/handler"
	"go-blog-api/repository"
)

func main() {
	log.Println("Memulai aplikasi blog...")

	connStr := config.GetDBConnectionString()

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke DB: %v", err)
	}
	log.Println("🎉 BERHASIL TERHUBUNG KE DATABASE 'blog_db' 🎉")

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	postRepo := repository.NewPostRepository(db)
	postHandler := handler.NewPostHandler(postRepo)

	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		// POST /users
		userRoutes.POST("", userHandler.CreateUser)

		// GET /users
		userRoutes.GET("", userHandler.GetAllUsers)

		// GET /users/:id
		userRoutes.GET("/:id", userHandler.GetUserByID)

		// GET /users/:id/posts
		userRoutes.GET("/:id/posts", postHandler.GetPostsByUserID)
	}

	postRoutes := r.Group("/posts")
	{
		// POST /posts
		postRoutes.POST("", postHandler.CreatePost)

		// GET /posts
		postRoutes.GET("", postHandler.GetAllPosts)
	}

	log.Println("Server Running")
	if err := r.Run("localhost:8000"); err != nil {
		log.Fatalf("Faild server: %v", err)
	}
}

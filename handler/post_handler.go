package handler

import (
	"go-blog-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Repo *repository.PostRepository
}

func NewPostHandler(repo *repository.PostRepository) *PostHandler {
	return &PostHandler{Repo: repo}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  int    `json:"user_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input JSON tidak valid"})
		return
	}

	post, err := h.Repo.CreatePost(req.Title, req.Content, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.Repo.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPostsByUserID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID user tidak valid"})
		return
	}

	posts, err := h.Repo.GetPostByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

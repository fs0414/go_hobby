package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/fs0414/go_hobby/internal/adapter/repository/interface"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type BoardUseCase struct {
	repo repository_interface.BoardRepositoryIf
}

type BoardCreateRequest struct {
	Content string `json:"content"`
	UserID uint `json:"user_id"`
}

func BoardUseCaseFactory(repo repository_interface.BoardRepositoryIf) *BoardUseCase {
	return &BoardUseCase{repo: repo}
}

func (uc *BoardUseCase) FetchBoards(c *gin.Context) {
	boards, err := uc.repo.GetBoards()
    if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, boards)
}

func (uc *BoardUseCase) CreateBoard(c *gin.Context) {
	var req BoardCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	board := model.Board{
        Content: req.Content,
        UserID: req.UserID,
    }
	newBoard, err := uc.repo.CreateBoard(board)

	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

	c.JSON(201, newBoard)
}
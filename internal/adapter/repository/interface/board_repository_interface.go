package repository_interface

import (
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type BoardRepositoryIf interface {
	GetBoards() ([]model.Board, error)
	CreateBoard(board model.Board) (model.Board, error)
}

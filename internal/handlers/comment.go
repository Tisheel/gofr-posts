package handlers

import (
	"strings"

	"github.com/example/internal/models"
	"gofr.dev/pkg/gofr"
)

type CommentHandler struct{}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

func (h *CommentHandler) CreateComment(ctx *gofr.Context) (any, error) {
	var comment models.Comment
	if err := ctx.Bind(&comment); err != nil {
		return nil, err
	}

	res, err := ctx.SQL.Exec("INSERT INTO comments (postId, name, email, body) VALUES (?, ?, ?, ?)", comment.PostId, comment.Name, comment.Email, comment.Body)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	comment.Id = int(id)
	return comment, nil
}

func (h *CommentHandler) GetCommentsByPostId(ctx *gofr.Context) (any, error) {
	postId := ctx.PathParam("id")

	var comments []models.Comment
	res, err := ctx.SQL.Query("SELECT * FROM comments WHERE postId = ?", postId)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var comment models.Comment
		if err = res.Scan(&comment.Id, &comment.PostId, &comment.Name, &comment.Email, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (h *CommentHandler) GetComments(ctx *gofr.Context) (any, error) {
	postId := ctx.Param("postId")
	name := ctx.Param("name")
	email := ctx.Param("email")
	body := ctx.Param("body")

	var conds []string
	var args []any

	if postId != "" {
		conds = append(conds, "postId = ?")
		args = append(args, postId)
	}

	if name != "" {
		conds = append(conds, "name = ?")
		args = append(args, name)
	}

	if email != "" {
		conds = append(conds, "email = ?")
		args = append(args, email)
	}

	if body != "" {
		conds = append(conds, "body = ?")
		args = append(args, body)
	}

	query := "SELECT * FROM comments"

	if len(conds) > 0 {
		query = query + " WHERE " + strings.Join(conds, " AND ")
	}

	res, err := ctx.SQL.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	for res.Next() {
		var comment models.Comment
		res.Scan(&comment.Id, &comment.PostId, &comment.Name, &comment.Email, &comment.Body)
		comments = append(comments, comment)
	}

	return comments, nil
}

package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/example/internal/models"
	"gofr.dev/pkg/gofr"
)

type PostHandler struct{}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

func (h *PostHandler) GetPosts(ctx *gofr.Context) (any, error) {
	title := ctx.Param("title")
	body := ctx.Param("body")

	var conds []string
	var args []any

	if title != "" {
		conds = append(conds, "title = ?")
		args = append(args, title)
	}

	if body != "" {
		conds = append(conds, "body = ?")
		args = append(args, body)
	}

	query := "SELECT * FROM posts"

	if len(conds) > 0 {
		query = query + " WHERE " + strings.Join(conds, ", ")
	}

	rows, err := ctx.SQL.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Body); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (h *PostHandler) GetPostById(ctx *gofr.Context) (any, error) {
	idStr := ctx.PathParam("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	var post models.Post
	err = ctx.SQL.QueryRow("SELECT id, title, body FROM posts WHERE id = ?", id).Scan(&post.Id, &post.Title, &post.Body)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (h *PostHandler) CreatePost(ctx *gofr.Context) (any, error) {
	var post models.Post
	if err := ctx.Bind(&post); err != nil {
		return nil, err
	}

	res, err := ctx.SQL.Exec("INSERT INTO posts (title, body) VALUES (?, ?)", post.Title, post.Body)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	post.Id = int(id)

	return post, nil
}

func (h *PostHandler) EditPost(ctx *gofr.Context) (any, error) {
	idStr := ctx.PathParam("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	var post models.Post
	if err := ctx.Bind(&post); err != nil {
		return nil, err
	}

	_, err = ctx.SQL.Exec("UPDATE posts SET title = ?, body = ? WHERE id = ?", post.Title, post.Body, id)
	if err != nil {
		return nil, err
	}

	post.Id = id
	return post, nil
}

func (h *PostHandler) EditPostByFields(ctx *gofr.Context) (any, error) {
	idStr := ctx.PathParam("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	var post models.Post
	if err := ctx.Bind(&post); err != nil {
		return nil, err
	}

	var updates []string
	var args []any

	if post.Title != "" {
		updates = append(updates, "title = ?")
		args = append(args, post.Title)
	}

	if post.Body != "" {
		updates = append(updates, "body = ?")
		args = append(args, post.Body)
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query := "UPDATE posts SET " + strings.Join(updates, ", ") + " WHERE id = ?"
	args = append(args, id)

	_, err = ctx.SQL.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (h *PostHandler) DeletePost(ctx *gofr.Context) (any, error) {
	idStr := ctx.PathParam("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	_, err = ctx.SQL.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

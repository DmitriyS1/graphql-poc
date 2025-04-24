package models

import "time"

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"createdAt"`
	Posts     []*Post    `json:"posts,omitempty"`
	Comments  []*Comment `json:"comments,omitempty"`
}

type Post struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	UserID    int        `json:"userId"`
	CreatedAt time.Time  `json:"createdAt"`
	User      *User      `json:"user,omitempty"`
	Comments  []*Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	PostID    int       `json:"postId"`
	UserID    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Post      *Post     `json:"post,omitempty"`
	User      *User     `json:"user,omitempty"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"userId"`
}

type CommentInput struct {
	Content string `json:"content"`
	PostID  int    `json:"postId"`
	UserID  int    `json:"userId"`
}

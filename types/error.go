package types

var PostNotFoundCode, PostNotFoundError = 1, "Post not found!"

type Error struct {
	Error struct{
		ErrorCode int `json:"error_code"`
		ErrorMsg string `json:"error_msg"`
	} `json:"error"`
}
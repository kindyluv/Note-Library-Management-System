package dto

import "github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/data"

type ReaderResponse struct {
	UserName string
	Book     []data.Book
}

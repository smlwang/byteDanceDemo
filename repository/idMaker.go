package repository

import "sync/atomic"

var (
	userId    int64
	videoId   int64
	commentId int64
)

func UserId() int64 {
	return atomic.AddInt64(&userId, 1)
}
func VideoId() int64 {
	return atomic.AddInt64(&videoId, 1)
}
func CommentId() int64 {
	return atomic.AddInt64(&commentId, 1)
}

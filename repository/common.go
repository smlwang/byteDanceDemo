package repository

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Created       int64  `json:"-"`
	AuthorRefer   int64  `json:"-"`
	Author        User   `json:"author" gorm:"foreignKey:AuthorRefer"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id              int64  `json:"id,omitempty"`
	UserKey_Comment int64  `json:"-"`
	User            User   `json:"user" gorm:"foreignKey:UserKey_Comment"`
	Content         string `json:"content,omitempty"`
	CreateDate      string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type IdInfo struct {
	UserId_Max    int64
	VideoId_Max   int64
	CommentId_Max int64
}

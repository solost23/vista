package forms

import "vista/pkg/utils"

type InsertCommentForm struct {
	Content  string `json:"content" binding:"required"`
	ParentID uint   `json:"parentId"`
}

type CommentListForm struct {
	*utils.PageForm
	VideoId *uint `form:"videoId" binding:"required"`
}

// type CommentListRecord struct {
// 	Id            *uint   `json:"id"`
// 	Content       *string `json:"content"`
// 	ParentId      *uint   `json:"parentId"`
// 	Type          *uint   `json:"type"`
// 	CreatedAt     *string `json:"createdAt"`
// 	UpdatedAt     *string `json:"updatedAt"`
// 	CreatorId     *uint   `json:"creatorId"`
// 	CreatorAvatar *string `json:"creatorAvatar"`
// }

// type CommentListResponse struct {
// 	Records  []*CommentListRecord `json:"records"`
// 	PageList *utils.PageList
// }

type GetCommentsUser struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type GetCommentsReply struct {
	Total uint                `json:"total"`
	List  []GetCommentsRecord `json:"list"`
}

type GetCommentsRecord struct {
	ID              uint               `json:"id"`
	ParentID        uint               `json:"parentId"`
	CreatorID       uint               `json:"uid"`
	Content         string             `json:"content"`
	CreatedAt       string             `json:"createdAt"`
	GetCommentsUser GetCommentsUser    `json:"user"`
	Replys          []GetCommentsReply `json:"reply"`
}

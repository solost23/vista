package forms

import "vista/pkg/utils"

type InsertCommentForm struct {
	Content  string `json:"content" binding:"required"`
	ParentID uint   `json:"parentId"`
}

type InsertComment struct {
	ID         uint        `json:"id"`
	ParentID   uint        `json:"parentId"`
	UserID     uint        `json:"userId"`
	Content    string      `json:"content"`
	CreatedAt  string      `json:"createdAt"`
	InsertUser CommentUser `json:"user"`
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

type CommentUser struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// type GetCommentsReply struct {
// 	Total uint                `json:"total"`
// 	List  []GetCommentsRecord `json:"list"`
// }

type CommentList struct {
	Id          uint           `json:"id"`
	ParentId    uint           `json:"parentId"`
	CreatorID   uint           `json:"uid"`
	Content     string         `json:"content"`
	CreatedAt   string         `json:"createdAt"`
	CommentUser CommentUser    `json:"user"`
	Children    []*CommentList `json:"reply"`
}

func (f *CommentList) ID() uint {
	return f.Id
}

func (f *CommentList) ParentID() uint {
	return f.ParentId
}

func (f *CommentList) AppendChildren(children any) {
	f.Children = append(f.Children, children.(*CommentList))
}

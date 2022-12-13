package delivery

import (
	"be13/project/features/comment"
)

type CommentRequest struct {
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Notes      string `json:"notes" form:"notes"`
	Ratings    int    `json:"ratings" form:"ratings"`
}

func UserRequestToUserCore(data CommentRequest) comment.CoreComment {
	return comment.CoreComment{
		HomestayID: data.HomestayID,
		Notes:      data.Notes,
		Ratings:    data.Ratings,
	}
}

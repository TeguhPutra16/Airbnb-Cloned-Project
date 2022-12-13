package delivery

import (
	"be13/project/features/comment"
)

type CommentRequest struct {
	HomestayID uint   `json:"home_stay_id" form:"home_stay_id"`
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

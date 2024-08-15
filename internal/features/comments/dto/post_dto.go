package dto

type AddCommentRequest struct {
	Comment string `json:"comment"`
}

type CommentResponse struct {
}

type CommentListResponse struct {
	TotalCount int                `json:"total_count"`
	TotalPages int                `json:"total_pages"`
	Page       int                `json:"page"`
	Size       int                `json:"size"`
	HasMore    bool               `json:"has_more"`
	Data       []*CommentResponse `json:"data"`
}

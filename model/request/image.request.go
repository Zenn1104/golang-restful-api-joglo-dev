package request

type ImageCreateRequest struct {
	CategoryId uint `validate:"required" form:"category_id" json:"category_id"`
}

package web

type CategoryUpdateRequest struct {
	Id int32 `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}
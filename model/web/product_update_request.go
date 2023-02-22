package web

type ProductUpdateRequest struct {
	Id 					int 	 `validate:"required"`
	Name 				string `validate:"required,max=100,min=1"`
	Description string `validate:"required"`
	Category 		string `validate:"required"`
	Price 			int 	 `validate:"required"`
}
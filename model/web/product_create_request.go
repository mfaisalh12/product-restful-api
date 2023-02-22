package web

type ProductCreateRequest struct {
	Name 				string `validate:"required,max=100,min=1"`
	Description string `validate:"required"`
	Category 		string `validate:"required"`
	Price 			int 	 `validate:"required"`
}
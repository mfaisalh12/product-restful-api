package web

type ProductCreateRequest struct {
	Name 				string `validate:"required,min=1,max=100"`
	Description string `validate:"required"`
	Category 		string `validate:"required"`
	Price 			int 	 `validate:"required"`
}
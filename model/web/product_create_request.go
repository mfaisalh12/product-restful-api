package web

type ProductCreateRequest struct {
	Name 				string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required" json:"description"`
	Category 		string `validate:"required" json:"category"`
	Price 			int 	 `validate:"required" json:"price"`
}
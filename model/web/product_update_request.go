package web

type ProductUpdateRequest struct {
	Id 					int 	 `validate:"required" json:"id"`
	Name 				string `validate:"required,max=100,min=1" json:"name"`
	Description string `validate:"required" json:"description"`
	Category 		string `validate:"required" json:"category"`
	Price 			int 	 `validate:"required" json:"price"`
}
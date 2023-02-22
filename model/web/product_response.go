package web

type ProductResponse struct {
	Id 					int			`json:"id"`
	Name 				string	`json:"name"`
	Description string	`json:"description"`
	Category 		string	`json:"category"`
	Price 			int			`json:"price"`
}
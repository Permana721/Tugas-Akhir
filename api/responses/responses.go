package responses

type BlogResponse struct {
	ID          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Category    *string    `json:"category"`
	Content     *string    `json:"content"`
	HeaderPhoto *string    `json:"header_photo"`
	Photo       *string    `json:"photo"`
}

type UserResponse struct {	
	ID     		*int    	`json:"id"`
	Name 		*string 	`json:"name"`
	Email     	*string  	`json:"email"`
	Password    *string  	`json:"password"`
	Age     	*int    	`json:"age"`
}
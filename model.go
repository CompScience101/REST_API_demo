package main


// Book struct (Model)
type Book struct {
	ID     string  `json:"id" binding: "required"`
	Isbn   string  `json:"isbn binding: "required""`
	Title  string  `json:"title" binding: "required"`
	Author *Author `json:"author" binding: "required"`
	//DateCreated string `json:"date_created"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname" binding: "required"`
	Lastname  string `json:"lastname" binding: "required"`
}


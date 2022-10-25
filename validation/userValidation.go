package validation

var UpdateUserBody struct {
	FirstName      string `binding:"omitempty,alpha"`
	LastName       string `binding:"omitempty,alpha"`
	Height         int    `binding:"omitempty,numeric"`
	Gender         string `binding:"omitempty,alpha"`
	Email          string `binding:"omitempty,email"`
	Password       string `binding:"omitempty,alphanum,min=8"`
	CommonName     string `binding:"omitempty,alpha"`
	DisplayPicture string `binding:"omitempty,url"`
	Birthday       string
}

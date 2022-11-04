package validation

var NewUserBody struct {
	FirstName  string `binding:"required" form:"firstName"`
	LastName   string `binding:"required" form:"lastName"`
	Height     int    `binding:"required" form:"height"`
	Gender     string `binding:"required" form:"gender"`
	Email      string `binding:"required,email" form:"email"`
	Password   string `binding:"required,alphanum,min=8" form:"password"`
	CommonName string `binding:"omitempty,alpha" form:"commonName"`
	Birthday   string `binding:"required" form:"birthday"`
}

var NewOrganizationBody struct {
	Name           string `binding:"required" form:"name"`
	Email          string `binding:"required,email" form:"email"`
	Password       string `binding:"required,alphanum,min=8" form:"password"`
	Est            int    `binding:"required" form:"est"`
	City           string `binding:"required" form:"city"`
	Facebook       string `form:"facebook"`
	Instagram      string `form:"instagram"`
	DisplayPicture string `form:"displayPicture"`
}

var Login struct {
	Email    string `binding:"required" form:"email"`
	Password string `binding:"required" form:"password"`
}

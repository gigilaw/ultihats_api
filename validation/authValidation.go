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

var Login struct {
	Email    string `binding:"required" form:"email"`
	Password string `binding:"required" form:"password"`
}

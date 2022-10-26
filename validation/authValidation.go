package validation

var NewUserBody struct {
	FirstName  string `binding:"required"`
	LastName   string `binding:"required"`
	Height     int    `binding:"required"`
	Gender     string `binding:"required"`
	Email      string `binding:"required,email"`
	Password   string `binding:"required,alphanum,min=8"`
	CommonName string `binding:"omitempty,alpha"`
	Birthday   string `binding:"required"`
}

var Login struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

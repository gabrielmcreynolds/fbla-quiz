package errorCodes

const (
	SameEmailName          = "user with same emails already exists"
	UserDoesNotExist       = "user does not exist"
	InvalidPassword        = "invalid password"
	DatabaseError          = "database error"
	InvalidEmail           = "invalid email"
	InvalidRefreshToken    = "invalid refresh token"
	TokenDoesNotExist      = "token does not exist"
	PasswordHashingFailure = "passing hashing failure"
)

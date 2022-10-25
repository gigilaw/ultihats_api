package config

var (
	ERROR_USER_LOGIN     = map[string]string{"message": "ERROR_USER_LOGIN", "details": "User not logged in"}
	ERROR_INVALID_TOKEN  = map[string]string{"message": "ERROR_INVALID_TOKEN", "details": "Cannot validate token"}
	ERROR_UNAUTHORIZED   = map[string]string{"message": "ERROR_UNAUTHORIZED", "details": "Unauthorized to perform this action"}
	ERROR_VALIDATION     = map[string]string{"message": "ERROR_VALIDATION"}
	ERROR_DATABASE       = map[string]string{"message": "ERROR_DATABASE"}
	ERROR_INVALID_LOGIN  = map[string]string{"message": "ERROR_INVALID_LOGIN", "details": "Invalid email or password"}
	ERROR_HASH_PASSWORD  = map[string]string{"message": "ERROR_HASH_PASSWORD", "details": "Failed to hash password"}
	ERROR_PARSE_BIRTHDAY = map[string]string{"message": "ERROR_PARSE_BIRTHDAY", "details": "Failed to parse birthday"}
)

package static

const (
	SUCCESS_CREATE_USER             = "success create new user"
	REQUEST_REQUIRED_MESSAGE        = "field {field} can't be empty"
	REQUEST_MIN_MESSAGE             = "field {field} must %d characters"
	REQUEST_MAX_MESSAGE             = "field {field} should be less then %d characters"
	REQUEST_EMAIL_MESSAGE           = "wrong format"
	ONLY_IMAGE_ALLOWED              = "only image allowed"
)

var (
	VALIDATOR_MESSAGE = map[string]string{
		"required": REQUEST_REQUIRED_MESSAGE,
		"email":    REQUEST_EMAIL_MESSAGE,
		"minLen":   REQUEST_MIN_MESSAGE,
		"maxLen":   REQUEST_MAX_MESSAGE,
	}
)

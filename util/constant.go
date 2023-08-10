package util

// Application constants
const (
	URL_INVALID = "provide a valid url"
)

// Status codes
const (
	SUCCESS_CODE        int = 200
	INTERNAL_ERROR_CODE int = 500
)

// Api messages
const (
	URLS_READ_SUCCESSFULLY   = "urls read successfully"
	URL_READ_SUCCESSFULLY    = "url read successfully"
	URL_CREATED_SUCCESSFULLY = "url created successfully"
	URL_DELETED_SUCCESSFULLY = "url deleted successfully"
)

// Server error messages
const (
	SERVER_LOADING_CREDENTIALS_ERROR = "error loading .env file"
)

// Server success messages
const (
	DATABASE_CONNECTED        = "database connected"
	SERVER_RUNNING_ON_PORT    = "server running on port"
	SERVER_RECOVER_FROM_ERROR = "recover from an error"
)

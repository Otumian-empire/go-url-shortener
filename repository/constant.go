package repository

// Literal constants
const (
	DEFAULT_INT = 0
)

// Database error messages
const (
	INSERT_ERROR              = "an error occurred while inserting"
	READING_ERROR             = "an error occurred while reading"
	DELETE_ERROR              = "an error occurred while deleting"
	DATABASE_OPENING_ERROR    = "an error occurred opening database"
	DATABASE_CONNECTING_ERROR = "an error occurred connecting database"
)

// Sql queries
const (
	SELECT_URL_QUERY       = "SELECT `id`, `short`, `original`, `created_at` FROM `url_shortener`"
	SELECT_URL_BY_ID_QUERY = "SELECT `id`, `short`, `original`, `created_at` FROM `url_shortener` WHERE`id`=?"
	INSERT_QUERY           = "INSERT INTO `url_shortener`(`short`, `original`) VALUES (?, ?)"
	DELETE_QUERY           = "DELETE FROM `url_shortener` WHERE`id`=?"
)

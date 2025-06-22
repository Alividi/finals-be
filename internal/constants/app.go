package constants

const (

	// AWS
	S3_ROOT_PATH                = "digisatlink"
	S3_ALLOWED_IMAGE_TYPES      = "profile"
	S3_ALLOWED_IMAGE_EXTENSIONS = "jpg,jpeg,png"
	S3_FILE_SIZE_LIMIT          = 10 * 1024 * 1024

	// Ticket Status
	TICKET_STATUS_OPEN        = "open"
	TICKET_STATUS_IN_PROGRESS = "in_progress"
	TICKET_STATUS_CLOSED      = "closed"

	// Roles
	ROLE_ADMIN    = "admin"
	ROLE_CUSTOMER = "customer"
	ROLE_TEKNISI  = "teknisi"

	// Notification Types
	NOTIFICATION_TYPE_SERVICE = "service"
	NOTIFICATION_TYPE_TICKET  = "ticket"
	NOTIFICATION_TYPE_PRODUCT = "product"
	NOTIFICATION_TYPE_AUTH    = "auth"
)

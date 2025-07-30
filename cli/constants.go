package cli

const (
	ErrMissingID            = "missing task ID"
	ErrInvalidNumber        = "Invalid ID, provide a valid number"
	ErrWrongNumberArguments = "Wrong number of Arguments"
	ErrDescription          = "Description cannot be empty."
	ErrInvalidMarkCommand   = "Invalid MarkCommand: %s. valid commands are mark-in-progress, mark-done"
	ErrMissingCommand       = "Missing Command. valid commands are add, delete, update, mark-in-progress, mark-done, list"
)

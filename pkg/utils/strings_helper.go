package utils

// Error Helpers
const (
	ARG_ERR                       = "Empty Argument"
	ARG_ERR_DETAIL                = "No argument (file name) added."
	ARG_ERR_SUGGESTIVE_CORRECTION = "Consider adding one or more file names with the syntax: go run main.go <filename1> <filename2> ..."

	FILE_READ_ERR                       = "File read error"
	FILE_READ_ERR_DETAIL                = "An error occurred while engine attempted reading the file."
	FILE_READ_ERR_SUGGESTIVE_CORRECTION = "Check that the file path exists or that your actual file path is correct."

	EMPTY_FILE_ERR                       = "Empty File"
	EMPTY_FILE_ERR_DETAIL                = "An EOF error occurred."
	EMPTY_FILE_ERR_SUGGESTIVE_CORRECTION = "The file you provided has no content. Kindly check it for corrections."
)

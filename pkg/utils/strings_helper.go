package utils

// Error Helpers Constants
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

	INVALID_INDICES_ERROR                       = "Invalid indices value"
	INVALID_INDICES_ERROR_DETAIL                = "The value you entered for the number of indices is not a valid number."
	INVALID_INDICES_ERROR_SUGGESTIVE_CORRECTION = "Try to enter a valid number for n in (<action>,n). Example (low,2)"

	INVALID_SYNTAX_ERR                       = "Invalid syntax"
	INVALID_SYNTAX_ERR_DETAIL                = "The syntax you entered is incorrect."
	INVALID_SYNTAX_ERR_SUGGESTIVE_CORRECTION = "Check your text to ensure that they follow the pattern (<action>,<indice value>) or (<action>)"
)

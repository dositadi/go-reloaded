package utils

import (
	t "edit-tool/pkg/models"
	"fmt"
)

// Helper function for printing errors out to the user.
func PrintErrorMessage(message t.Error) {
	errorMessage := fmt.Sprintf("\n********** Error **********\n\n%s: %s\n\nENGINE REPORT: %s\n\n******** Thank You ********\n", message.Err, message.Details, message.SuggestiveCorrection)

	fmt.Println(errorMessage)
}

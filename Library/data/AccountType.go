package data

type AccountType struct {
	StudentAccountType   string `json:"student_account_type" binding:"required"`
	LibrarianAccountType string `json:"librarian_account_type" binding:"required"`
}

package models 

// Structure of a job entry 
type Job struct {
	ID			int 	`json:"id"`
	Company 	string	`json:"company"`
	Title 		string	`json:"title"`
	Status 		string 	`json:"status"` // e.g. applied, interviewing, rejected
	DateApplied string 	`json:"date_applied"`
}
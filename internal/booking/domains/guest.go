package domains

import "fmt"

type Guest struct {
	FirstName   string
	MiddleName  string
	LastName    string
	Email       string
	PhoneNumber string
}

func (g *Guest) FullName() string {
	return fmt.Sprintf("%s %s %s", g.FirstName, g.MiddleName, g.LastName)
}

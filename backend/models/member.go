package models

import (
	"fmt"
	"net/http"
	"time"
)

type Member struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name"`
	LastName     string    `json:"last_name"`
	DOB          time.Time `json:"dob"`
	Gender       string    `json:"gender"`
	IDNo         string    `json:"id_no"`
	IDIssueDate  time.Time `json:"id_issue_date"`
	IDIssuePlace string    `json:"id_issue_place"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	CreatedOn    time.Time `json:"created_on,omitempty"`
	CreatedBy    string    `json:"created_by"`
	UpdatedOn    time.Time `json:"updated_on,omitempty"`
	UpdatedBy    string    `json:"updated_by"`
	ExpiresOn    time.Time `json:"expires_on"`
}

type MemberList struct {
	Members []Member `json:"members"`
}

func (i *Member) Bind(r *http.Request) error {
	if i.FirstName == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*MemberList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Member) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

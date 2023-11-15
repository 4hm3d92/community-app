package db

import (
	"context"
	//"database/sql"
	"time"

	"github.com/4hm3d92/community-app/backend/models"
	"github.com/jackc/pgx/v5"
)

func (db Database) GetAllMembers() (*models.MemberList, error) {
	list := &models.MemberList{}

	rows, err := db.Pool.Query(context.Background(), "SELECT id, first_name, middle_name, last_name, dob, gender, id_issue_date, id_issue_place, phone, email, id_no FROM members ORDER BY id DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.DOB, &member.Gender, &member.IDIssueDate, &member.IDIssuePlace, &member.Phone, &member.Email, &member.IDNo)
		if err != nil {
			return list, err
		}
		list.Members = append(list.Members, member)
	}
	return list, nil
}

func (db Database) AddMember(member *models.Member) error {
	var id int
	var created_on time.Time
	query := `INSERT INTO members (first_name, middle_name, last_name, dob, gender, id_issue_date, id_issue_place, phone, email, id_no) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, created_on`
	err := db.Pool.QueryRow(context.Background(), query, member.FirstName, member.MiddleName, member.LastName, member.DOB, member.Gender, member.IDIssueDate, member.IDIssuePlace, member.Phone, member.Email, member.IDNo).Scan(&id, &created_on)
	if err != nil {
		return err
	}

	member.ID = id
	member.CreatedOn = created_on
	return nil
}

/*
func (db Database) GetMemberById(memberId int) (models.Member, error) {
	member := models.Member{}

	query := `SELECT * FROM members WHERE id = $1;`
	row := db.Pool.QueryRow(context.Background(), query, memberId)
	switch err := row.Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName); err {
	case sql.ErrNoRows:
		return member, ErrNoMatch
	default:
		return member, err
	}
}
*/

func (db Database) DeleteMember(memberId int) error {
	query := `DELETE FROM members WHERE id = $1;`
	_, err := db.Pool.Exec(context.Background(), query, memberId)
	switch err {
	case pgx.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateMember(memberId int, memberData models.Member) (models.Member, error) {
	member := models.Member{}
	query := `UPDATE members SET first_name=$1, middle_name=$2 WHERE id=$3 RETURNING id, first_name, updated_on;`
	_, err := db.Pool.Exec(context.Background(), query, memberData.FirstName, memberData.MiddleName, memberId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return member, ErrNoMatch
		}
		return member, err
	}

	return member, nil
}

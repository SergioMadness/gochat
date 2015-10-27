package models

/**
* Profile model
 */

import (
	"chat/helpers"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Profile struct {
	Id                  int
	id                  int
	Username            string
	password            string
	RegDate             int64
	IsBlocked           bool
	AccessToken         string
	AccessTokenDateTime int
	Model
}

/**
* Save user's profile
* Insert to DB
 */
func (p *Profile) Save() bool {
	var err error
	var stmt *sql.Stmt

	if p.Id == 0 {
		stmt, err = p.GetConnection().Prepare("INSERT INTO profile (username, password, reg_date, is_blocked) VALUES (?, ?, ?, ?)")
		_, err = stmt.Exec(p.Username, p.password, p.RegDate, p.IsBlocked)
	} else {
		stmt, err = p.GetConnection().Prepare("UPDATE profile SET username=? WHERE id=?")
		_, err = stmt.Exec(p.Username, p.Id)
	}

	if err != nil {
		return false
	}

	return true
}

/**
* Update user's token
 */
func (p *Profile) UpdateToken() (string, error) {
	p.AccessToken = helpers.GetMD5(p.Username + p.GetPassword())

	stmt, err := p.GetConnection().Prepare("UPDATE profile SET access_token=?, access_token_datetime=NOW() WHERE id=?")
	_, err = stmt.Exec(p.AccessToken, p.Id)

	return p.AccessToken, err
}

/**
* Get user's password
 */
func (p *Profile) GetPassword() string {
	return p.password
}

/**
* Set password
 */
func (p *Profile) SetPassword(password string) {
	p.password = helpers.GetMD5(password)
}

/**
* Create new profile
 */
func NewProfile(conn *sql.DB) *Profile {
	result := new(Profile)

	result.RegDate = time.Now().Unix()
	result.SetConnection(conn)

	return result
}

/**
* Find profile by login and password
 */
func (p *Profile) FindByCredentials(username, password string) *Profile {
	rp := NewProfile(p.GetConnection())
	token := new(sql.NullString)
	tokenDate := new(sql.NullInt64)

	err := p.GetConnection().QueryRow("SELECT * FROM profile WHERE username=? AND password=?", username, password).Scan(&rp.id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked, &token, &tokenDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		rp.Id = rp.id
		if token != nil {
			rp.AccessToken = token.String
			rp.AccessTokenDateTime = int(tokenDate.Int64)
		}
	}

	return rp
}

/**
* Get profile by ID
 */
func (p *Profile) GetById(id int) *Profile {
	result := NewProfile(p.GetConnection())

	token := new(sql.NullString)
	tokenDate := new(sql.NullInt64)

	err := p.GetConnection().QueryRow("SELECT * FROM profile WHERE id=?", id).Scan(&result.Id, &result.Username, &result.password, &result.RegDate, &result.IsBlocked, token, tokenDate)

	if result.Id <= 0 || err != nil {
		return nil
	}

	if token != nil {
		result.AccessToken = token.String
		result.AccessTokenDateTime = int(tokenDate.Int64)
	}

	return result
}

/**
* Get list of users by IDs
 */
func (p *Profile) GetUsersByIds(ids []int) []Profile {
	var result []Profile

	rows, err := p.GetConnection().Query("SELECT id, username, reg_date, is_blocked FROM profile WHERE id IN (" + helpers.JoinI(ids, ",") + ")")

	if err == nil {
		for rows.Next() {
			var profile Profile
			if err := rows.Scan(&profile.id, &profile.Username, &profile.RegDate, &profile.IsBlocked); err != nil {
				log.Fatal(err)
			} else {
				result = append(result, profile)
			}
		}
	}

	return result
}

/**
*Find profile by username
 */
func (p *Profile) FindByUsername(username string) *Profile {
	rp := NewProfile(p.GetConnection())

	err := p.GetConnection().QueryRow("SELECT id, username, password, reg_date, is_blocked FROM profile WHERE username=?", username).Scan(&rp.id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked)

	if rp.Id <= 0 || err != nil {
		return nil
	}

	return rp
}

/**
* Get user by token
 */
func (p *Profile) GetByToken(token string) *Profile {
	rp := NewProfile(p.GetConnection())

	err := p.GetConnection().QueryRow("SELECT id, username, password, reg_date, is_blocked FROM profile WHERE access_token=?", token).Scan(&rp.id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked)

	if rp.Id <= 0 || err != nil {
		return nil
	}

	return rp
}

/**
* Search for profiles
 */
func (p *Profile) Find(searchStr string) []Profile {
	var result []Profile

	rows, err := p.GetConnection().Query(fmt.Sprintf("SELECT id, username, reg_date, is_blocked FROM profile WHERE username LIKE '%%%s%%'", searchStr))

	if err == nil {
		for rows.Next() {
			var profile Profile
			if err := rows.Scan(&profile.id, &profile.Username, &profile.password, &profile.RegDate, &profile.IsBlocked); err != nil {
				log.Fatal(err)
			} else {
				result = append(result, profile)
			}
		}
	} else {
		fmt.Println(err.Error())
	}

	return result
}

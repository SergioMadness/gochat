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
	Id        int
	id        int
	Username  string
	password  string
	RegDate   int64
	IsBlocked bool
	Model
}

/**
* Save user's profile
* Insert to DB
 */
func (p *Profile) Save() bool {
	stmt, _ := p.GetConnection().Prepare("INSERT INTO profile (username, password, reg_date, is_blocked) VALUES (?, ?, ?, ?)")
	insertResult, insertErr := stmt.Exec(p.Username, p.password, p.RegDate, p.IsBlocked)

	if insertErr != nil {
		return false
	} else {
		lastId, _ := insertResult.LastInsertId()
		p.Id = int(lastId)
	}

	return true
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

	err := p.GetConnection().QueryRow("SELECT * FROM profile WHERE username=? AND password=?", username, password).Scan(&rp.id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		rp.Id = rp.id
	}

	return rp
}

/**
* Get profile by ID
 */
func (p *Profile) GetById(id int) *Profile {
	result := NewProfile(p.GetConnection())

	err := p.GetConnection().QueryRow("SELECT * FROM profile WHERE id=?", id).Scan(&result.Id, &result.Username, &result.password, &result.RegDate, &result.IsBlocked)

	if result.Id <= 0 || err != nil {
		return nil
	}

	return result
}

/**
* Get list of users by IDs
 */
func (p *Profile) GetUsersByIds(ids []int) []Profile {
	var result []Profile

	rows, err := p.GetConnection().Query("SELECT * FROM profile WHERE id IN (" + helpers.JoinI(ids, ",") + ")")

	if err == nil {
		for rows.Next() {
			var profile Profile
			if err := rows.Scan(&profile.Id, &profile.Username, &profile.password, &profile.RegDate, &profile.IsBlocked); err != nil {
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

	err := p.GetConnection().QueryRow("SELECT * FROM profile WHERE username=?", username).Scan(&rp.Id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked)

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

	rows, err := p.GetConnection().Query(fmt.Sprintf("SELECT * FROM profile WHERE username LIKE '%%%s%%'", searchStr))

	if err == nil {
		for rows.Next() {
			var profile Profile
			if err := rows.Scan(&profile.Id, &profile.Username, &profile.password, &profile.RegDate, &profile.IsBlocked); err != nil {
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

func (p *Profile) GetOpenKey() *OpenKey {
	openKey := NewOpenKey(p.GetConnection())

	return openKey.GetByIdProfile(p.Id)
}

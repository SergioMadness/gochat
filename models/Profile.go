package models

/**
* Profile model
 */

import (
	"database/sql"
	"io/ioutil"
)

type Profile struct {
	Id                  int
	AccessToken         string
	AccessTokenDateTime int
	Model
}

/**
* Update user's token
 */
func (p *Profile) UpdateToken() error {
	stmt, err := p.GetConnection().Prepare("UPDATE user SET access_token_datetime=NOW() WHERE id=?")
	_, err = stmt.Exec(p.Id)

	return err
}

/**
* Create new profile
 */
func NewProfile(conn *sql.DB) *Profile {
	result := new(Profile)

	result.SetConnection(conn)

	return result
}

/**
* Get user by token
 */
func (p *Profile) GetByToken(token string) *Profile {
	rp := NewProfile(p.GetConnection())

	err := p.GetConnection().QueryRow("SELECT id FROM \"user\" WHERE access_token=$1", token).Scan(&rp.Id)
	if rp.Id <= 0 || err != nil {
		ioutil.WriteFile("errr.log", []byte(err.Error()), 0777)
		return nil
	}

	return rp
}

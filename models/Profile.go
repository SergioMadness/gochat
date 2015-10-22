package models

import (
	"chat/helpers"
	"database/sql"
	"fmt"
	"time"
)

type Profile struct {
	Id        int
	id        int
	Username  string
	password  string
	RegDate   int64
	IsBlocked bool
	Conn      *sql.DB
}

func (p *Profile) Save() bool {
	stmt, err := p.Conn.Prepare("INSERT INTO profile (username, password, reg_date, is_blocked) VALUES (?, ?, ?, ?)")
	fmt.Println("Password3: " + p.GetPassword())
	_, err = stmt.Exec(p.Username, p.password, p.RegDate, p.IsBlocked)

	if err != nil {
		return false
	}

	return true
}

func (p *Profile) GetId() int {
	return p.id
}

func (p *Profile) GetPassword() string {
	return p.password
}

func (p *Profile) SetPassword(password string) {
	fmt.Println("Password: " + password)
	p.password = helpers.GetMD5(password)
	fmt.Println("Password2: " + p.password)
}

func NewProfile(conn *sql.DB) *Profile {
	result := new(Profile)

	result.RegDate = time.Now().Unix()
	result.Conn = conn

	return result
}

func (p *Profile) FindByCredentials(username, password string) *Profile {
	rp := NewProfile(p.Conn)

	err := p.Conn.QueryRow("SELECT * FROM profile WHERE username=? AND password=?", username, password).Scan(&rp.id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		rp.Id = rp.id
	}

	return rp
}

func (p *Profile) FindByUsername(username string) *Profile {
	rp := NewProfile(p.Conn)

	err := p.Conn.QueryRow("SELECT * FROM profile WHERE username=?", username).Scan(&rp.id, &rp.Username, &rp.password, &rp.RegDate, &rp.IsBlocked)

	if rp.id <= 0 || err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return rp
}

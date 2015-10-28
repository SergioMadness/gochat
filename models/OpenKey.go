package models

import (
	"crypto/dsa"
	"database/sql"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"
)

type OpenKey struct {
	id        int
	IdProfile int
	Key       string
	Model
}

func NewOpenKey(conn *sql.DB) *OpenKey {
	result := new(OpenKey)

	result.SetConnection(conn)

	return result
}

/**
* Save open key
* Insert to DB
 */
func (k *OpenKey) Save() bool {
	stmt, err := k.GetConnection().Prepare("INSERT INTO profile_key (id_profile, open_key) VALUES (?, ?)")
	_, err = stmt.Exec(k.IdProfile, k.Key)

	if err != nil {
		return false
	}

	return true
}

/**
* Get key by id
 */
func (k *OpenKey) GetById(id int) *OpenKey {
	result := NewOpenKey(k.GetConnection())

	err := k.GetConnection().QueryRow("SELECT * FROM profile_key WHERE id=?", id).Scan(&result.id, &result.IdProfile, &result.Key)

	if result.id <= 0 || err != nil {
		return nil
	}

	return result
}

/**
* Get key by id_profile
 */
func (k *OpenKey) GetByIdProfile(idProfile int) *OpenKey {
	result := NewOpenKey(k.GetConnection())

	err := k.GetConnection().QueryRow("SELECT * FROM profile_key WHERE id_profile=?", idProfile).Scan(&result.id, &result.IdProfile, &result.Key)

	if result.id <= 0 || err != nil {
		return nil
	}

	return result
}

func (k *OpenKey) Validate(hashed, sig string) bool {
	return true
}

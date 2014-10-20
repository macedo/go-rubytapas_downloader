package main

import "strings"

type Credential struct {
	Username, Password string
}

func NewCredential(credentials string) *Credential {
	var username, password string

	func(s []string, vars ...*string) {
		for i, str := range s {
			*vars[i] = str
		}
	}(strings.Split(credentials, ":"), &username, &password)
	return &Credential{username, password}
}

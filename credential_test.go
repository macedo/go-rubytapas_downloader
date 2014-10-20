package main

import "testing"

func TestNewCredential(t *testing.T) {
	credentials := "username:password"
	credential := NewCredential(credentials)

	if credential.Username != "username" {
		t.Error("Expected username, got", credential.Username)
	}

	if credential.Password != "password" {
		t.Error("Expected password, got", credential.Password)
	}
}

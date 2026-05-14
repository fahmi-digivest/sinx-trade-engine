package server

import "strings"

// Authenticator validates SoupBinTCP login credentials.
type Authenticator struct {
	credentials map[string]string // username (lower) -> password (lower)
}

// NewAuthenticator creates an Authenticator from a map of username -> password pairs.
func NewAuthenticator(creds map[string]string) *Authenticator {
	normalised := make(map[string]string, len(creds))
	for u, p := range creds {
		normalised[strings.ToLower(u)] = strings.ToLower(p)
	}
	return &Authenticator{credentials: normalised}
}

// Validate returns true when the username/password combination is valid.
// Both fields are treated as case-insensitive per the spec.
func (a *Authenticator) Validate(username, password string) bool {
	expected, ok := a.credentials[strings.ToLower(username)]
	if !ok {
		return false
	}
	return expected == strings.ToLower(password)
}

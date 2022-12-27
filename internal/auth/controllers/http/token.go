package controllers

type AuthInfo struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type TokenInfo struct {
	Token string `json:"token,omitempty"`
}

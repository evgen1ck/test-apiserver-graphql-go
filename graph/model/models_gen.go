// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthUserPayload struct {
	Token    string `json:"token"`
	UUID     string `json:"uuid"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type LoginInput struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	EmailCode *string `json:"emailCode"`
	TwoFaCode *string `json:"twoFaCode"`
}

type SignupInput struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupWithCodeInput struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

package domain

import "errors"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginReq) Validate() error {
	if l.Username == "" {
		return errors.New("username wajib di isi")
	}

	if l.Password == "" {
		return errors.New("password wajib di isi")
	}

	return nil
}

type LoginResp struct {
	ID          string `json:"id"`
	RoleID      string `json:"role_id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

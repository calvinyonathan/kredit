package Login

type loginRequest struct {
	Nik      string `json:"nik"`
	Password string `json:"password"`
}

type updatePasswordRequest struct {
	Nik         string `json:"nik"`
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}

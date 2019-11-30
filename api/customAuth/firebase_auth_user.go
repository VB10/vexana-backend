package customAuth

type FirebaseUserRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

//func (f FirebaseUserRequest) Read(p []byte) (n int, err error) {
//	panic("implement me")
//}

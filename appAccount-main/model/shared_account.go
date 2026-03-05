package model

type SharedAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country,omitempty"`
	Status   string `json:"status,omitempty"`
	Time     string `json:"time,omitempty"`
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"` // shadowrocket / surge / quantumult
}

package http

type TokenPair struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}

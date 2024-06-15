package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewSuccesResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccesResponse(data interface{}) *successRes {
	return NewSuccesResponse(data, nil, nil)
}

type userResponesToken struct {
	AccessToken  interface{} `json:"accesstoken"`
	RefreshToken interface{} `json:"refreshtoken"`
}

func NewReponseUserToken(accessToken, refreshToken string) *userResponesToken {
	return &userResponesToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

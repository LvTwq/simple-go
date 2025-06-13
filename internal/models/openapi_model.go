package models

type ApplyAppIdVo struct {
	AppId       string   `json:"appId"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type AccessTokenDto struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type AccessTokenVo struct {
	Token        string `json:"token"`
	ExpireIn     int64  `json:"expireIn"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenDto struct {
	AppId        string `json:"appId"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshAccessTokenVo struct {
	Token    string `json:"token"`
	ExpireIn int64  `json:"expireIn"`
}

type OpenApiIdModel struct {
	AppId       string `gorm:"column:app_id"`
	AppSecret   string `gorm:"column:secret"`
	Description string `gorm:"column:description"`
}

type OpenapiAppIdPermissionMp struct {
	AppId      string `gorm:"column:app_id"`
	Permission string `gorm:"column:permission"`
}

type PageVo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PageData[T any] struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Data     []T   `json:"data"`
	Total    int64 `json:"total"`
}

type KeyValueVo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AppIdDeleteVo struct {
	AppIds []string `json:"appIds"`
}

type IptablesVo struct {
	Chain  string   `json:"chain"`
	Table  string   `json:"table"`
	Rule   string   `json:"rule"`
	Params []string `json:"params"`
}

type AreaVo struct {
	Country string
	Region  string
	City    string
}

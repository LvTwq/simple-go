package models

type ReqSendSmsVo struct {
	Key  string    `json:"key"`
	Data []SmsData `json:"data"`
}

type SmsData struct {
	Phone   string `json:"phone_number"`
	Content string `json:"send_content"`
}

type RespSendSmsResult struct {
	RetCode string  `json:"retCode"`
	RetMsg  string  `json:"retMsg"`
	RetData RetData `json:"retData"`
}

type RetData struct {
	Result int    `json:"result"`
	Code   string `json:"send_sms_code"`
}

type ReqQueryVo struct {
	Key  string `json:"key"`
	Code string `json:"send_sms_code"`
}

type RespQueryVo struct {
	RetCode string `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	RetData string `json:"retData"`
}

type Detail struct {
	Phone  string `json:"phone_number"`
	Status int    `json:"is_send_status"`
}

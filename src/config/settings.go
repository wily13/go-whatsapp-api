package config

type Browser string

var (
	AppVersion string = "3.4.0"
	AppIP      string = "172.20.103.227"
	AppPort    string = "3000"
	AppDebug   bool   = false

	PathQrCode    string = "statics/qrcode"
	PathSendItems string = "statics/senditems"

	DBName string = "hydrogenWaCli.db"

	WhatsappLogLevel            string = "ERROR"
	WhatsappAutoReplyMessage    string
	WhatsappSettingMaxFileSize  int64 = 30000000 // 10MB
	WhatsappSettingMaxVideoSize int64 = 30000000 // 30MB
)

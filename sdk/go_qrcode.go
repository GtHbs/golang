package sdk

import "github.com/skip2/go-qrcode"

func GenerateQrCode() {
	qrcode.WriteFile("http://www.baidu.com", qrcode.Highest, 1024, "./qr.png")
}

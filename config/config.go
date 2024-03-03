package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	TextRecognitionAsyncUrl = "https://ocr.api.cloud.yandex.net/ocr/v1/recognizeTextAsync"
	GetRecognitionUrl       = "https://ocr.api.cloud.yandex.net/ocr/v1/getRecognition?operationId="
)

var (
	Port                 string
	Host                 string
	YandexCloudOcrApiKey string
)

func init() {
	var err error

	if os.Getenv("GO_ENV") != "production" {
		if err = godotenv.Load(); err != nil {
			panic(err)
		}
	}

	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")
	YandexCloudOcrApiKey = os.Getenv("YANDEX_CLOUD_OCR_API_KEY")

	if len(Port) == 0 || len(Host) == 0 || len(YandexCloudOcrApiKey) == 0 {
		panic("Invalid env variables")
	}
}

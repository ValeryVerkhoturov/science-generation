package yandexCloudApi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ValeryVerkhoturov/chat/config"
	"github.com/ValeryVerkhoturov/chat/utils/yandexCloudApi/models"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

func recognizeText(oCRequest models.OCRRequest) ([]models.OCRResponse, error) {
	body, err := json.Marshal(oCRequest)
	if err != nil {
		return []models.OCRResponse{}, err
	}

	req, err := http.NewRequest("POST", config.TextRecognitionAsyncUrl, bytes.NewBuffer(body))
	if err != nil {
		return []models.OCRResponse{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Api-Key %s", config.YandexCloudOcrApiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []models.OCRResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing response body:", err)
		}
	}(resp.Body)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.OCRResponse{}, err
	}

	var oCRAsyncResponse models.OCRAsyncResponse
	err = json.Unmarshal(responseBody, &oCRAsyncResponse)
	if err != nil {
		return []models.OCRResponse{}, err
	}

	for {
		req, err := http.NewRequest("GET", config.GetRecognitionUrl+oCRAsyncResponse.Id, nil)
		if err != nil {
			return []models.OCRResponse{}, err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Api-Key %s", config.YandexCloudOcrApiKey))
		req.Header.Set("Content-Type", "application/json")

		resp, err = http.Get(config.GetRecognitionUrl + oCRAsyncResponse.Id)
		if err != nil {
			return []models.OCRResponse{}, err
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return []models.OCRResponse{}, err
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Error("Error closing response body:", err)
			}
		}(resp.Body)

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return []models.OCRResponse{}, err
		}

		responseBodyString := string(responseBody)
		if strings.HasPrefix(responseBodyString, "{\"error\"") {
			var oCRResponse models.GetRecognitionResultResponseError
			err = json.Unmarshal(responseBody, &oCRResponse)
			if err != nil {
				return []models.OCRResponse{}, err
			}

			if oCRResponse.Error.HttpCode != 404 {
				return []models.OCRResponse{}, fmt.Errorf(oCRResponse.Error.Message)
			}
			time.Sleep(5 * time.Second)
			continue
		}

		responseBodies := strings.Split(responseBodyString, "\n")
		pages := make([]models.OCRResponse, len(responseBodies))
		for i := range responseBodies {
			if responseBodies[i] == "" {
				continue
			}
			var oCRResponse models.OCRResponse
			err = json.Unmarshal([]byte(responseBodies[i]), &oCRResponse)
			if err != nil {
				return []models.OCRResponse{}, err
			}
			pages = append(pages, oCRResponse)
		}
		return pages, nil
	}
}

func extractText(models []models.OCRResponse) string {
	var buffer bytes.Buffer
	for i, model := range models {
		if i != 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(model.Result.TextAnnotation.FullText)
	}
	return buffer.String()
}

func RecognizeTextFromPdf(pdfContent []byte) (string, error) {
	ocrRequest := models.OCRRequest{
		MimeType:      "application/pdf",
		LanguageCodes: []string{"*"},
		Model:         "page",
		Content:       base64.StdEncoding.EncodeToString(pdfContent),
	}
	responseModels, err := recognizeText(ocrRequest)
	if err != nil {
		return "", err
	}

	return extractText(responseModels), nil
}

package cbr

import (
	"fmt"
	"net/http"
	"time"

	"encoding/xml"

	"golang.org/x/net/html/charset"
)

func FetchData(dateReq string) (*ValCurs, error) {
	url := "https://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=" + dateReq

	client := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка HTTP запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP статус %d", resp.StatusCode)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var data ValCurs
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("ошибка парсинга XML: %w", err)
	}

	return &data, nil
}

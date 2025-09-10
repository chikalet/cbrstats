package cbr

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

func FetchData(dateReq string) ([]byte, error) {
	url := "https://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=" + dateReq

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status %d", resp.StatusCode)
	}

	utf8Reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		utf8Reader, err = charset.NewReaderLabel("windows-1251", resp.Body)
		if err != nil {
			return nil, fmt.Errorf("charset convert error: %w", err)
		}
	}
	data, err := io.ReadAll(utf8Reader)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}

	return data, nil
}

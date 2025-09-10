package cbr

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

func ParseValCurs(data []byte) (*ValCurs, error) {
	var vc ValCurs
	if err := xml.Unmarshal(data, &vc); err != nil {
		return nil, fmt.Errorf("xml ошибка: %w", err)
	}
	return &vc, nil
}

func RateForValute(v Valute) (float64, error) {
	valStr := strings.TrimSpace(strings.ReplaceAll(v.Value, ",", "."))
	if valStr == "" {
		return 0, fmt.Errorf("ошибка значения")
	}
	valFloat, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка: %w", err)
	}
	if v.Nominal <= 0 {
		return 0, fmt.Errorf("неверно: %d", v.Nominal)
	}
	return valFloat / float64(v.Nominal), nil
}

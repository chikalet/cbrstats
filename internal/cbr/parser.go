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
		return nil, fmt.Errorf("xml unmarshal error: %w", err)
	}
	return &vc, nil
}

func RateForValute(v Valute) (float64, error) {
	valStr := strings.TrimSpace(strings.ReplaceAll(v.Value, ",", "."))
	if valStr == "" {
		return 0, fmt.Errorf("empty value")
	}
	valFloat, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	if v.Nominal <= 0 {
		return 0, fmt.Errorf("invalid nominal: %d", v.Nominal)
	}
	return valFloat / float64(v.Nominal), nil
}

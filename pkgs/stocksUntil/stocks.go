package stocksUntil

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// Stock 结构体
type Stock struct {
	Code       string
	Name       string
	Percent    float64
	Updown     float64
	Price      float64
	Open       float64
	Low        float64
	High       float64
	YestClose  float64
	Profit     string
	HoldPrice  float64
	HoldNumber int
}

// SinaStockTransform 新浪数据转换结构体
type SinaStockTransform struct {
	code   string
	params []string
}

// Name 获取名称
func (s *SinaStockTransform) Name() string {
	switch s.code[:2] {
	case "sh":
		if len(s.params) > 0 {
			return s.decodeGBK(s.params[0])
		}
		return "---"
	case "sz":
		if len(s.params) > 0 {
			return s.decodeGBK(s.params[0])
		}
		return "---"
	case "hk":
		if len(s.params) > 1 {
			return s.decodeGBK(s.params[1])
		}
		return "---"
	case "us":
		if len(s.params) > 0 {
			return s.decodeGBK(s.params[0])
		}
		return "---"
	case "bj":
		if len(s.params) > 0 {
			return s.decodeGBK(s.params[0])
		}
		return "---"
	default:
		return "---"
	}
}

// Price 获取现价
func (s *SinaStockTransform) Price() float64 {
	switch s.code[:2] {
	case "sh":
		if len(s.params) > 3 {
			price, _ := strconv.ParseFloat(s.params[3], 64)
			return price
		}
		return 0
	case "sz":
		if len(s.params) > 3 {
			price, _ := strconv.ParseFloat(s.params[3], 64)
			return price
		}
		return 0
	case "hk":
		if len(s.params) > 6 {
			price, _ := strconv.ParseFloat(s.params[6], 64)
			return price
		}
		return 0
	case "us":
		if len(s.params) > 1 {
			price, _ := strconv.ParseFloat(s.params[1], 64)
			return price
		}
		return 0
	case "bj":
		if len(s.params) > 3 {
			price, _ := strconv.ParseFloat(s.params[3], 64)
			return price
		}
		return 0
	default:
		return 0
	}
}

// Low 获取最低价
func (s *SinaStockTransform) Low() float64 {
	switch s.code[:2] {
	case "sh":
		if len(s.params) > 5 {
			low, _ := strconv.ParseFloat(s.params[5], 64)
			return low
		}
		return 0
	case "sz":
		if len(s.params) > 5 {
			low, _ := strconv.ParseFloat(s.params[5], 64)
			return low
		}
		return 0
	case "hk":
		if len(s.params) > 5 {
			low, _ := strconv.ParseFloat(s.params[5], 64)
			return low
		}
		return 0
	case "us":
		if len(s.params) > 7 {
			low, _ := strconv.ParseFloat(s.params[7], 64)
			return low
		}
		return 0
	case "bj":
		if len(s.params) > 5 {
			low, _ := strconv.ParseFloat(s.params[5], 64)
			return low
		}
		return 0
	default:
		return 0
	}
}

// High 获取最高价
func (s *SinaStockTransform) High() float64 {
	switch s.code[:2] {
	case "sh":
		if len(s.params) > 4 {
			high, _ := strconv.ParseFloat(s.params[4], 64)
			return high
		}
		return 0
	case "sz":
		if len(s.params) > 4 {
			high, _ := strconv.ParseFloat(s.params[4], 64)
			return high
		}
		return 0
	case "hk":
		if len(s.params) > 4 {
			high, _ := strconv.ParseFloat(s.params[4], 64)
			return high
		}
		return 0
	case "us":
		if len(s.params) > 6 {
			high, _ := strconv.ParseFloat(s.params[6], 64)
			return high
		}
		return 0
	case "bj":
		if len(s.params) > 4 {
			high, _ := strconv.ParseFloat(s.params[4], 64)
			return high
		}
		return 0
	default:
		return 0
	}
}

// Yestclose 获取昨日收盘价
func (s *SinaStockTransform) Yestclose() float64 {
	switch s.code[:2] {
	case "sh":
		if len(s.params) > 2 {
			yestclose, _ := strconv.ParseFloat(s.params[2], 64)
			return yestclose
		}
		return 0
	case "sz":
		if len(s.params) > 2 {
			yestclose, _ := strconv.ParseFloat(s.params[2], 64)
			return yestclose
		}
		return 0
	case "hk":
		if len(s.params) > 3 {
			yestclose, _ := strconv.ParseFloat(s.params[3], 64)
			return yestclose
		}
		return 0
	case "us":
		if len(s.params) > 26 {
			yestclose, _ := strconv.ParseFloat(s.params[26], 64)
			return yestclose
		}
		return 0
	case "bj":
		if len(s.params) > 2 {
			yestclose, _ := strconv.ParseFloat(s.params[2], 64)
			return yestclose
		}
		return 0
	default:
		return 0
	}
}

// Percent 获取涨跌百分比
func (s *SinaStockTransform) Percent() float64 {
	price := s.Price()
	if price == 0 {
		return 0
	}
	return (price - s.Yestclose()) / s.Yestclose()
}

func TransformPercent(num float64) string {
	multiplier := 100.0
	roundedPercentage := math.Round(num*100*multiplier) / multiplier
	return fmt.Sprintf("%.2f%%", roundedPercentage)
}

// Open 获取今开价
func (s *SinaStockTransform) Open() float64 {
	if len(s.params) > 1 {
		open, _ := strconv.ParseFloat(s.params[1], 64)
		return open
	}
	return 0
}

// Updown 获取涨跌价格
func (s *SinaStockTransform) Updown() float64 {
	return s.Price() - s.Yestclose()
}

// 将 GBK 编码字符串转换为 UTF-8
func (s *SinaStockTransform) decodeGBK(str string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	output, _ := io.ReadAll(reader)
	return string(output)
}

// 转换为结构体
func (s *SinaStockTransform) transform() Stock {
	return Stock{
		Code:      s.code,
		Name:      s.Name(),
		Percent:   s.Percent(),
		Updown:    s.Updown(),
		Price:     s.Price(),
		Open:      s.Open(),
		Low:       s.Low(),
		High:      s.High(),
		YestClose: s.Yestclose(),
	}
}

// SinaStockProvider 新浪数据提供结构体
type SinaStockProvider struct {
}

// Fetch 获取数据
func (s *SinaStockProvider) Fetch(code string) (Stock, error) {
	result := Stock{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://hq.sinajs.cn/", nil)
	if err != nil {
		return Stock{}, err
	}
	q := req.URL.Query()
	q.Add("list", code)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("referer", "https://finance.sina.com.cn/")
	resp, err := client.Do(req)
	if err != nil {
		return Stock{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Stock{}, err
	}
	responseStr := string(body)
	parts := strings.Split(responseStr, "=")
	if len(parts) != 2 {
		return Stock{}, fmt.Errorf("invalid response for code %s", code)
	}
	data := strings.ReplaceAll(parts[1], `"`, "")
	params := strings.Split(data, ",")
	transformer := SinaStockTransform{code: code, params: params}
	result = transformer.transform()
	return result, nil
}

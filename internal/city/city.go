package city

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"sync"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	//go:embed city_code.csv
	csvData []byte
)

var CityClient = NewCity()

type City struct {
	CodeMap map[string]string
	mu      sync.Mutex
}

func NewCity() *City {
	return &City{
		CodeMap: make(map[string]string),
	}
}

func (c *City) LoadCodeMap() error {

	reader := csv.NewReader(bytes.NewReader(csvData))
	reader.FieldsPerRecord = -1 // 允许不一致字段数量

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("读取 CSV 文件失败: %v", err)
	}

	tmpMap := make(map[string]string)
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		cityName := record[0]
		adcode := record[1]
		if len(cityName) == 0 || len(adcode) == 0 {
			continue
		}
		tmpMap[cityName] = adcode
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CodeMap = tmpMap
	return nil
}

func (c *City) GetAdcode(cityName string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 先尝试直接获取
	adcode, exists := c.CodeMap[cityName]
	if exists {
		return adcode, true
	}
	// 如果直接获取失败，尝试模糊匹配
	for name, code := range c.CodeMap {
		matches := fuzzy.Match(cityName, name)
		if matches {
			return code, true
		}
	}
	return adcode, false
}

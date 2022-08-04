package salary

import "fmt"

const (
	requestURL = ""
)

type Data = map[string]interface{}

var (
	data = map[string]Data{
		"lanzhou": {
			"name":    "lanzhou",
			"country": "CN",
			"val":     500,
		},
		"beijing": {
			"name":    "beijing",
			"country": "CN",
			"val":     1600,
		},
	}
)

func GetCitySalary(cityName string) (Data, error) {
	d, ok := data[cityName]
	if !ok {
		return Data{}, fmt.Errorf("city %s salary data not found", cityName)
	}

	return d, nil
}

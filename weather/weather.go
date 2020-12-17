package weather

import (
	"encoding/json"
	"ff/g"
	"fmt"
	"io/ioutil"
)

// 高德API文档地址：https://lbs.amap.com/api/webservice/guide/api/weatherinfo/

// 查询城市编码
func QueryCity(searchStr string) (cityArr interface{}) {
	uri := "/v3/geocode/geo"
	key := g.VP.GetString("weather.secret-key")
	api := g.VP.GetString("weather.api")
	url := fmt.Sprintf("%s%s?key=%s&address=%s", api, uri, key, searchStr)
	resp, err := g.HttpClient.Get(url)
	if err == nil {
		res := make(map[string] interface{})
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			err := json.Unmarshal(data, &res)
			if err == nil {
				return res["geocodes"]
			}
		}
	}
	g.Logger.Error(fmt.Sprintln("query city code failed:", err))
	return [0]int{}
}

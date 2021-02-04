package weather

import (
	"context"
	"encoding/json"
	"ff/g"
	"ff/model"
	"ff/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// 高德API文档地址：https://lbs.amap.com/api/webservice/guide/api/weatherinfo/

type InfoResp struct {
	Status   string    `json:"status"`
	Count    string    `json:"count"`
	Info     string    `json:"info"`
	InfoCode string    `json:"infocode"`
	Lives    []Weather `json:"lives"`
}

/*
"province" : "陕西",
"city" : "长安区",
"adcode" : "610116",
"weather" : "晴",
"temperature" : "8",
"winddirection" : "西南",
"windpower" : "≤3",
"humidity" : "36",
"reporttime" : "2021-01-25 17:55:04"
*/

type Weather struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	WindDirection string `json:"winddirection"`
	WindPower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	ReportTime    string `json:"reporttime"`
}

// 查询城市编码
func QueryCity(searchStr string) (cityArr interface{}) {
	uri := "/v3/geocode/geo"
	key := g.VP.GetString("weather.secret-key")
	api := g.VP.GetString("weather.api")
	url := fmt.Sprintf("%s%s?key=%s&address=%s", api, uri, key, searchStr)
	resp, err := g.HttpClient().Get(url)
	if err == nil {
		res := make(map[string]interface{})
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

// 根据城市编码获取城市天气
func WeatherInfoNow(adcode string) (weather *Weather) {
	uri := "/v3/weather/weatherInfo"
	key := g.VP.GetString("weather.secret-key")
	api := g.VP.GetString("weather.api")
	url := fmt.Sprintf("%s%s?key=%s&city=%s", api, uri, key, adcode)
	var resp *http.Response
	var err error
	var data []byte
	resp, err = g.HttpClient().Get(url)
	res := &InfoResp{}
	if err == nil {
		data, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			err = json.Unmarshal(data, &res)
			if err == nil {
				return &res.Lives[0]
			}
		}
	}
	g.Logger.Error(fmt.Sprintln("get city weather failed:", adcode, err))
	return &res.Lives[0]
}

// 根据城市编码去redis缓存中获取天气情况
func WeatherCached(adcode string) (map[string]string, error) {
	ctx := context.Background()
	return g.RDB().HGetAll(ctx, "weather:"+adcode).Result()
}

// 缓存某地区天气信息
func CacheWeather(weather *Weather) {
	if *weather == (Weather{}) || weather == nil {
		g.Logger.Warn("无效的天气信息")
		return
	}
	ctx := context.Background()
	g.RDB().HSet(ctx, "weather:"+weather.Adcode,
		map[string]interface{}{
			"province":      weather.Province,
			"city":          weather.City,
			"weather":       weather.Weather,
			"temperature":   weather.Temperature,
			"winddirection": weather.WindDirection,
			"windpower":     weather.WindPower,
			"humidity":      weather.Humidity,
			"updatetime":    time.Now().Format("2006-01-02 15:04:05"),
			"reporttime":    weather.ReportTime})
}

// 拉取并更新天气信息
func pullUpdateWeatherInfo(adcode string) {
	res := WeatherInfoNow(adcode)
	CacheWeather(res)
}

func weatherInfoTask() error {
	randomValue, err := g.UIDGen().NextID()
	if err != nil {
		g.Logger.Error("生成分布式锁随机值失败，更新天气信息任务终止")
		return err
	}
	var ctx = context.Background()
	weatherLock := utils.RedLock{
		Key:    "weatherLock",
		Value:  strconv.FormatUint(randomValue, 10),
		Expiry: time.Minute * g.VP.GetDuration("weather.update-frequency"),
		Ctx:    ctx}
	if err := weatherLock.Lock(); err != nil {
		g.Logger.Error("加锁失败，终止更新天气信息")
		return err
	}
	defer func() {
		if err := weatherLock.Unlock(); err != nil {
			g.Logger.Warn("更新天气信息完成，但解锁失败")
		}
	}()
	var adcodes []string
	if err := g.DB().Model(&model.AuthUser{}).Pluck("distinct(city_adcode)", &adcodes).Error; err != nil {
		g.Logger.Error("获取城市编码列表失败")
		return err
	}
	weatherChan := make(chan string, 1)
	limiter := time.Tick(time.Second / g.VP.GetDuration("weather.weather-qps"))
	go func() {
		for _, item := range adcodes {
			<-limiter
			weatherChan <- item
		}
		close(weatherChan)
	}()

	for {
		if adcode, ok := <-weatherChan; ok && adcode != "" {
			go pullUpdateWeatherInfo(adcode)
		} else {
			break
		}
	}
	return nil
}

func RealtimeWeather(interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		g.Logger.Info("==>开始更新天气信息...")
		if err := weatherInfoTask(); err != nil {
			g.Logger.Warn("拉取天气信息任务执行失败")
		}
		g.Logger.Info("==>更新天气信息完成...")
	}
}

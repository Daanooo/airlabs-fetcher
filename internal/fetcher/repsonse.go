package fetcher

import "time"

type Response struct {
	Request struct {
		Lang     string `json:"lang"`
		Currency string `json:"currency"`
		Time     int    `json:"time"`
		ID       string `json:"id"`
		Server   string `json:"server"`
		Host     string `json:"host"`
		Pid      int    `json:"pid"`
		Key      struct {
			ID             int       `json:"id"`
			APIKey         string    `json:"api_key"`
			Type           string    `json:"type"`
			Expired        time.Time `json:"expired"`
			Registered     time.Time `json:"registered"`
			Upgraded       any       `json:"upgraded"`
			LimitsByHour   int       `json:"limits_by_hour"`
			LimitsByMinute int       `json:"limits_by_minute"`
			LimitsByMonth  int       `json:"limits_by_month"`
			LimitsTotal    int       `json:"limits_total"`
		} `json:"key"`
		Params struct {
			Lang string `json:"lang"`
		} `json:"params"`
		Version int    `json:"version"`
		Method  string `json:"method"`
		Client  struct {
			IP  string `json:"ip"`
			Geo struct {
				CountryCode string  `json:"country_code"`
				Country     string  `json:"country"`
				Continent   string  `json:"continent"`
				City        string  `json:"city"`
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
				Timezone    string  `json:"timezone"`
			} `json:"geo"`
			Connection struct {
				Type    string `json:"type"`
				IspCode int    `json:"isp_code"`
				IspName string `json:"isp_name"`
			} `json:"connection"`
			Device struct {
			} `json:"device"`
			Agent struct {
			} `json:"agent"`
			Karma struct {
				IsBlocked bool `json:"is_blocked"`
				IsCrawler bool `json:"is_crawler"`
				IsBot     bool `json:"is_bot"`
				IsFriend  bool `json:"is_friend"`
				IsRegular bool `json:"is_regular"`
			} `json:"karma"`
		} `json:"client"`
	} `json:"request"`
	Response []struct {
		Name        string  `json:"name"`
		IataCode    any     `json:"iata_code"`
		IcaoCode    string  `json:"icao_code"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
		CountryCode string  `json:"country_code"`
	} `json:"response"`
}

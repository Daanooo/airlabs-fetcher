package fetcher

import "strconv"

type Airport struct {
	Name        string  `json:"name"`
	IataCode    any     `json:"iata_code"`
	IcaoCode    string  `json:"icao_code"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	CountryCode string  `json:"country_code"`
}

func (a Airport) toStringSlice() []string {
	iata := ""
	if a.IataCode != nil {
		iata = a.IataCode.(string)
	} else {
		iata = ""
	}

	return []string{
		a.Name,
		iata,
		a.IcaoCode,
		strconv.FormatFloat(a.Lat, 'f', -1, 64),
		strconv.FormatFloat(a.Lng, 'f', -1, 64),
		a.CountryCode,
	}
}

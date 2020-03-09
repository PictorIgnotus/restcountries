package restcountries

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://restcountries.eu/rest/v2/%s"


func GetAllCountries() ([]Country, error) {
	resData, err := doRestcountriesCall("all")
	if err != nil {
		return nil, err
	}
	return parseCountries(resData)
}

func GetCountriesByName(name string) ([]Country, error) {
	resData, err := doRestcountriesCall(fmt.Sprintf("name/%s", name))
	if err != nil {
		return nil, err
	}
	return parseCountries(resData)
}

func GetCountryByCapital(name string) ([]Country, error) {
	resData, err := doRestcountriesCall(fmt.Sprintf("capital/%s", name))
	if err != nil {
		return nil, err
	}
	return parseCountries(resData)
}

func GetCountryByName(name string) ([]Country, error) {
	resData, err := doRestcountriesCall(fmt.Sprintf("name/%s?fullText=true", name))
	if err != nil {
		return nil, err
	}
	return parseCountries(resData)
}

func doRestcountriesCall(apiSuffix string) ([]byte, error) {
	url := fmt.Sprintf(baseURL, apiSuffix)
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		e := errors.New(fmt.Sprintf("Unexpected API status code %s", res.Status))
		return []byte{}, e
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}


func parseCountries(jsonData []uint8) ([]Country, error) {
	var c []Country
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}


package greeting

import "io/ioutil"
import "net/http"

func GetCountry(code string) (output string) {
	resp, err := http.Get("http://services.groupkt.com/country/get/iso2code/" + code)
	if err != nil {
		output = "bad url"
	} else {
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			output = "Bad read"
		} else {
			output = string(bytes)
		}
	}
	return
}

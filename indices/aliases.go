package indices

import (
	"github.com/mattbaird/elastigo/api"
	"encoding/json"
)

func UpdateAlias(data interface{}) (api.BaseResponse, error) {
	var retval api.BaseResponse

	body, err := api.DoCommand("POST", "/_aliases", data)

	if err != nil {
		return retval, err
	}
	if err == nil {
		// marshall into json
		jsonErr := json.Unmarshal(body, &retval)
		if jsonErr != nil {
			return retval, jsonErr
		}
	}

	return retval, err
}

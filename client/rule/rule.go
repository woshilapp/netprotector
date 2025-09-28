package rule

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/woshilapp/netprotector/client/global"
)

func GetRules() (*global.Rules, error) {
	res, err := http.Get("http://" + global.ServerAddr + "/api/rules")
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100%10 != 2 {
		return nil, errors.New("Response Error: " + res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	var rules *global.Rules = &global.Rules{}
	err = json.Unmarshal(data, rules)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	return rules, nil
}

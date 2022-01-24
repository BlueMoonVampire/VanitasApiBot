package utils

import (
	"bot/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Check(user int) (*types.CheckResponse, error) {
	var url string = "https://sylviorus-api.up.railway.app/user"
	res, err := http.Get(fmt.Sprintf("%v/%v", url, user))
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(res.Body)
	var k types.CheckResponse
	json.Unmarshal(data, &k)
	return &k, nil
}

func Ban(user, enforcer int, reason, admin_token string) {
	var url string = "https://sylviorus-api.up.railway.app/ban"
	
	jsom := map[string]string{
		"user" : string(user),
		"enforcer" : string(enforcer),
		"admin_token" : admin_token,
		"reason" : reason,
	}

	values,_ := json.Marshal(jsom)

	res , err := http.Post(url , "application/json" , bytes.NewBuffer(values))
	if err != nil {
		println(err)
	}

	str,_ := ioutil.ReadAll(res.Body)
	println(string(str))

}

func Unban(user int) {

	var url string = "https://sylviorus-api.up.railway.app/unban"
	str := fmt.Sprintf(`{"user" : %v}`, user)
	jsom := []byte(str)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsom))
	req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

}

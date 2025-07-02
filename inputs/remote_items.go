package inputs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ComboBoxItemsRemote struct {
}

type TextDefault struct {
	ID string `json:"id"`
}

func (c ComboBoxItemsRemote) GetItems(sr string, token string) []ComboItem {
	cl := http.Client{}
	var ls []ComboItem
	req, err := http.NewRequest("GET", sr, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ls
	}
	req.Header.Add("Authorization", token)
	res, err := cl.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ls
	}
	fmt.Println(res.StatusCode)
	err = json.NewDecoder(res.Body).Decode(&ls)
	if err != nil {
		fmt.Println(err.Error())
		return ls
	}
	return ls
}

func (td *TextDefault) GetDefault(url string, token string) {
	cl := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		td.ID = ""
		return
	}
	req.Header.Add("Authorization", token)
	res, err := cl.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		td.ID = ""
		return
	}
	err = json.NewDecoder(res.Body).Decode(td)
	if err != nil {
		fmt.Println(err.Error())
		td.ID = ""
		return
	}
}

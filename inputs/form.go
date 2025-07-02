package inputs

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Form struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
	Submit  Submit   `json:"submit"`
	Access  []string `json:"access"`
}

type Column struct {
	Inputs []Input `json:"inputs"`
}

type Submit struct {
	Text           string `json:"text"`
	Source         string `json:"source"`
	PageSize       int    `json:"pageSize"`
	ItemsCount     string `json:"itemsCount"`
	Method         string `json:"method"`
	ConfirmMessage string `json:"confirmMessage"`
	LastAction     string `json:"lastAction"` //reload-info,success-message,close
	SuccessMessage string `json:"successMessage"`
	Type           string `json:"type"` //finder,loader
}

type FormExported struct {
	Columns []Column `json:"columns"`
	Submit  Submit   `json:"submit"`
}

var Hosts map[string]string

func (f Form) GetName() string {
	return f.Name
}

func (f Form) GetAccess() []string {
	return f.Access
}

func (f *Form) Generate(access []string, token string, kid, acKid string) *FormExported {
	if !existsAccess(access, f.Access) {
		return nil
	}
	fe := FormExported{Submit: f.Submit}
	ri := ComboBoxItemsRemote{}
	for _, vls := range f.Columns {
		clm := Column{}
		for _, val := range vls.Inputs {
			if !existsAccess(access, val.Access) {
				continue
			}
			if val.Type == "combo-box" {
				if (val.Items == nil || len(val.Items) == 0) && len(val.ItemsSource) > 2 {
					fmt.Println(val.ItemsSource)
					val.Items = ri.GetItems(hostsReplace(val.ItemsSource), token)
					if len(val.Items) > 0 {
						val.DefaultValue = fmt.Sprintf("%v", val.Items[0].ID)
					}
				}
			}
			if val.Type == "text-view" || val.Type == "number-view" || val.Type == "label" || val.Type == "date-time" {
				if val.DefaultValue == "time-now" {
					val.DefaultValue = time.Now().Format("03-04-05")
				}
				val.DefaultValue = hostsReplace(val.DefaultValue)
				_, err := url.ParseRequestURI(val.DefaultValue)
				if err == nil {
					td := TextDefault{}
					td.GetDefault(val.DefaultValue, token)
					val.DefaultValue = td.ID
				}
			}
			val.Access = nil
			val.ItemsSource = ""
			clm.Inputs = append(clm.Inputs, val)
		}
		fe.Columns = append(fe.Columns, clm)
	}
	fe.urlCorrect(kid, acKid)
	return &fe
}

func (fe *FormExported) urlCorrect(kid, acKid string) {
	if kid != acKid {
		if fe.Submit.Source != "" {
			fe.Submit.Source = "/reverse/" + kid + fe.Submit.Source
		}
		if fe.Submit.ItemsCount != "" {
			fe.Submit.ItemsCount = "/reverse/" + kid + fe.Submit.ItemsCount
		}
		for i := 0; i < len(fe.Columns); i++ {
			for j := 0; j < len(fe.Columns[i].Inputs); j++ {
				if fe.Columns[i].Inputs[j].InfoSource != "" {
					fe.Columns[i].Inputs[j].InfoSource = "/reverse/" + kid + fe.Columns[i].Inputs[j].InfoSource
				}
				if fe.Columns[i].Inputs[j].Search != "" {
					fe.Columns[i].Inputs[j].Search = "/reverse/" + kid + fe.Columns[i].Inputs[j].Search
				}
				if fe.Columns[i].Inputs[j].FileSource != "" {
					fe.Columns[i].Inputs[j].FileSource = "/reverse/" + kid + fe.Columns[i].Inputs[j].FileSource
				}
			}
		}
	}
}

func existsAccess(access []string, list []string) bool {
	for _, val := range list {
		for _, acc := range access {
			if val == acc {
				return true
			}
		}
	}
	return false
}

func hostsReplace(url string) string {
	uNew := url
	for key, val := range Hosts {
		uNew = strings.ReplaceAll(uNew, "["+key+"]", val)
	}
	return uNew
}

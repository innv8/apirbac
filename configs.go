package apirbac

import (
	"encoding/json"
	"io/ioutil"
)

func (r *RBAC) LoadConfigs(fileName string) error {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContents, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *RBAC) SaveConfigs(fileName string) error {
	configBytes, _ := json.Marshal(r.Configs)
	err := ioutil.WriteFile(fileName, configBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

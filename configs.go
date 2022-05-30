package apirbac

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (r *RBAC) LoadConfigs(fileName string) error {
	var configs RBAConfigs
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContents, &configs)
	if err != nil {
		return err
	}

	r.Configs = configs

	return nil
}

func (r *RBAC) SaveConfigs(fileName string) error {
	configBytes, _ := json.Marshal(r.Configs)
	err := ioutil.WriteFile(fileName, configBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}

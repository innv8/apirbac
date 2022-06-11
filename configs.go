package apirbac

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadConfigs loads configs from a json file.
// if the file is invalid, it returns an error.
func (r *RBAC) LoadConfigs(fileName string) error {
	var configs RBACConfigs
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

//LoadConfigBytes is same as LoadConfigs only that it loads a file read as bytes
func (r *RBAC) LoadConfigBytes(bytes []byte) error {
	var err error
	var configs RBACConfigs
	err = json.Unmarshal(bytes, &configs)
	if err != nil {
		return err
	}

	r.Configs = configs

	return nil
}

// SaveConfigs saves the configs to a json file for persistence/ export
func (r *RBAC) SaveConfigs(fileName string) error {
	configBytes, _ := json.Marshal(r.Configs)
	err := ioutil.WriteFile(fileName, configBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}

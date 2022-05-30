package apirbac

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

func (r *RBAC) LoadConfigs(fileName string) error {
	var configs RBAConfigs
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContents, &configs)
	if err != nil {
		return err
	}

	for _, role := range configs.Roles {
		for _, grant := role.Grants {
			grant.Resource.Rgx = regexp.MustCompile(grant.Resource.Regex)
		}
	}

	r.Configs = configs

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

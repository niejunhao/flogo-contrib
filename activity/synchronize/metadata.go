package synchronize

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Key     string `md:"key"`
}

type Settings struct {
	Operation string `md:"Operation,required,allowed(Lock,Unlock)"`
}

// ToMap conversion
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"firstString":  i.Key,
	}
}

// FromMap conversion
func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.Key, err = coerce.ToString(values["key"])
	if err != nil {
		return err
	}
	return nil
}

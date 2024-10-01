package json

import (
	"errors"

	"github.com/ssh-connection-manager/file"
)

func Generate(filePath string, nameFile string) error {
	fullPath := GetPathToConnectFile(filePath, nameFile)

	if !file.IsExistFile(fullPath) {
		file.CreateFile(fullPath)

		err := CreateBaseJsonDataToFile(fullPath)
		if err != nil {
			return errors.New("err generate base json: " + err.Error())
		}
	}

	return nil
}

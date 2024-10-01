package json

import (
	"errors"

	"github.com/ssh-connection-manager/file"
)

func GetPathToConnectFile(filePath string, nameFile string) string {
	fullPath := file.GetFullPath(filePath, nameFile)

	return fullPath
}

func CreateBaseJsonDataToFile(fullPath string) error {
	connections := Connections{
		Connects: []Connect{},
	}

	connect, err := connections.deserializationJson()
	if err != nil {
		return errors.New("error create json: " + err.Error())
	}

	err = file.WriteFile(fullPath, connect)
	if err != nil {
		return errors.New("error write to json: " + err.Error())
	}

	return nil
}

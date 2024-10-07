package json

import (
	"errors"

	"github.com/ssh-connection-manager/file"
)

type Connections struct {
	Connects []Connect `json:"connects"`
}

type Connect struct {
	Alias    string `json:"alias"`
	Login    string `json:"login"`
	Address  string `json:"address"`
	Password string `json:"password"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Connections) getConnectData(filePath string, nameFile string, fileKey string) error {
	fileConnect := GetPathToConnectFile(filePath, nameFile)

	data, err := file.ReadFile(fileConnect)
	if err != nil {
		return err
	}

	err = c.SerializationJson(data)
	if err != nil {
		return err
	}

	err = c.SetDecryptData(filePath, fileKey)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connections) GetDataForListConnect(filePath string, nameFile string, fileKey string) ([][]string, error) {
	var result [][]string

	err := c.getConnectData(filePath, nameFile, fileKey)
	if err != nil {
		return result, err
	}

	for _, v := range c.Connects {
		newElement := []string{v.Alias, v.CreatedAt, v.UpdatedAt}
		result = append(result, newElement)
	}

	if len(result) == 0 {
		return result, errors.New("no connection found")
	}

	return result, nil
}

func (c *Connections) GetConnectionsAlias(filePath string, nameFile string, fileKey string) ([]string, error) {
	var result []string

	err := c.getConnectData(filePath, nameFile, fileKey)
	if err != nil {
		return result, err
	}

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	if len(result) == 0 {
		return result, errors.New("no connection found")
	}

	return result, nil
}

func (c *Connections) ExistConnectJsonByIndex(alias string, filePath string, nameFile string, fileKey string) (int, error) {
	var noFound = -1

	err := c.getConnectData(filePath, nameFile, fileKey)
	if err != nil {
		return noFound, err
	}

	defer func() {
		err = c.SetCryptAllData(filePath, fileKey)
	}()

	for i, v := range c.Connects {
		if v.Alias == alias {
			return i, nil
		}
	}

	return noFound, errors.New("not found")
}

func (c *Connections) WriteConnectToJson(connect Connect, filePath string, nameFile string, fileKey string) error {
	_, err := c.ExistConnectJsonByIndex(connect.Alias, filePath, nameFile, fileKey)
	if err == nil {
		return err
	}

	fullPath := GetPathToConnectFile(filePath, nameFile)

	data, err := file.ReadFile(fullPath)
	if err != nil {
		return err
	}

	err = c.SerializationJson(data)
	if err != nil {
		return err
	}

	encodedConnect, err := SetCryptData(connect, filePath, fileKey)
	if err != nil {
		return err
	}

	c.Connects = append(c.Connects, encodedConnect)

	deserializationJson, err := c.deserializationJson()
	if err != nil {
		return err
	}

	err = file.WriteFile(fullPath, deserializationJson)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connections) updateJsonDataByIndex(index int, connect Connect) error {
	if index >= 0 && index < len(c.Connects) {
		c.Connects[index].Alias = connect.Alias
		c.Connects[index].Address = connect.Address
		c.Connects[index].Login = connect.Login
		c.Connects[index].Password = connect.Password
		c.Connects[index].UpdatedAt = connect.UpdatedAt

		return nil
	}

	return errors.New("connection update error")
}

func (c *Connections) UpdateConnectJson(alias string, connect Connect, filePath string, nameFile string, fileKey string) error {
	index, err := c.ExistConnectJsonByIndex(alias, filePath, nameFile, fileKey)
	if err != nil {
		return err
	}

	cryptData, err := SetCryptData(connect, filePath, fileKey)
	if err != nil {
		return err
	}

	err = c.updateJsonDataByIndex(index, cryptData)
	if err != nil {
		return err
	}

	deserializationJson, err := c.deserializationJson()
	if err != nil {
		return err
	}

	err = file.WriteFile(GetPathToConnectFile(filePath, nameFile), deserializationJson)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connections) deleteJsonDataByIndex(index int) {
	copy(c.Connects[index:], c.Connects[index+1:])

	c.Connects[len(c.Connects)-1] = Connect{}
	c.Connects = c.Connects[:len(c.Connects)-1]
}

func (c *Connections) DeleteConnectToJson(alias string, filePath string, nameFile string, fileKey string) error {
	index, err := c.ExistConnectJsonByIndex(alias, filePath, nameFile, fileKey)
	if err != nil {
		return err
	}

	c.deleteJsonDataByIndex(index)

	deserializationJson, err := c.deserializationJson()
	if err != nil {
		return err
	}

	err = file.WriteFile(GetPathToConnectFile(filePath, nameFile), deserializationJson)
	if err != nil {
		return err
	}

	return nil
}

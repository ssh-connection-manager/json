package json

import (
	"errors"
	"github.com/ssh-connection-manager/crypt"
)

func SetCryptData(c Connect, pathConf string, fileNameKey string) (Connect, error) {
	var err error

	errMess := errors.New("err set crypt data")

	c.Alias, err = crypt.Encrypt(c.Alias, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.Address, err = crypt.Encrypt(c.Address, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.Login, err = crypt.Encrypt(c.Login, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.Password, err = crypt.Encrypt(c.Password, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.CreatedAt, err = crypt.Encrypt(c.CreatedAt, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.UpdatedAt, err = crypt.Encrypt(c.UpdatedAt, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}

	return c, nil
}

func decryptData(c Connect, pathConf string, fileNameKey string) (Connect, error) {
	var err error

	errMess := errors.New("err decrypt data")

	c.Alias, err = crypt.Decrypt(c.Alias, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.Address, err = crypt.Decrypt(c.Address, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.Login, err = crypt.Decrypt(c.Login, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.Password, err = crypt.Decrypt(c.Password, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.CreatedAt, err = crypt.Decrypt(c.CreatedAt, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}
	c.UpdatedAt, err = crypt.Decrypt(c.UpdatedAt, pathConf, fileNameKey)
	if err != nil {
		return c, errMess
	}

	return c, nil
}

func (c *Connections) SetDecryptData(pathConf string, fileNameKey string) error {
	errMess := errors.New("err set decrypt data")

	for key, connect := range c.Connects {
		data, err := decryptData(connect, pathConf, fileNameKey)
		if err != nil {
			return errMess
		}

		c.Connects[key] = data
	}

	return nil
}

func (c *Connections) SetCryptAllData(pathConf string, fileNameKey string) error {
	errMess := errors.New("err set decrypt data all")

	for key, connect := range c.Connects {
		data, err := SetCryptData(connect, pathConf, fileNameKey)
		if err != nil {
			return errMess
		}

		c.Connects[key] = data
	}

	return nil
}

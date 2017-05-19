package cmd

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/dropbox/godropbox/errors"
	"github.com/pritunl/pritunl-link/config"
	"github.com/pritunl/pritunl-link/errortypes"
	"net/url"
)

func UnifiUsername(username string) (err error) {
	config.Config.Unifi.Username = username

	err = config.Save()
	if err != nil {
		return
	}

	logrus.WithFields(logrus.Fields{
		"username": config.Config.Unifi.Username,
	}).Info("cmd.unifi: Set Unifi username")

	return
}

func UnifiPassword(password string) (err error) {
	config.Config.Unifi.Password = password

	err = config.Save()
	if err != nil {
		return
	}

	logrus.WithFields(logrus.Fields{
		"password": config.Config.Unifi.Password,
	}).Info("cmd.unifi: Set Unifi password")

	return
}

func UnifiController(controller string) (err error) {
	u, err := url.Parse(controller)
	if err != nil {
		err = &errortypes.ParseError{
			errors.New("cmd.unifi: Failed to parse controller url"),
		}
		return
	}

	config.Config.Unifi.Controller = fmt.Sprintf("%s://%s", u.Scheme, u.Host)

	err = config.Save()
	if err != nil {
		return
	}

	logrus.WithFields(logrus.Fields{
		"controller": config.Config.Unifi.Controller,
	}).Info("cmd.unifi: Set Unifi controller")

	return
}

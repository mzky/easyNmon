package common

import "github.com/sirupsen/logrus"

func Handle(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

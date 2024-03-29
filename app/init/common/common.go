package common

import (
	"github.com/sirupsen/logrus"
	"zuoxingtao/init/log"
)

var MtVersion string
var SessionID int

func CommonInit() error {
	commonClient := NewCommonClient()
	version, err := commonClient.GetMTVersion()
	if err != nil {
		return err
	}
	MtVersion = version
	sessionID, err := commonClient.GetSessionID()
	if err != nil {
		return err
	}
	SessionID = sessionID
	return nil
}

func CommonUpdate() {
	commonClient := NewCommonClient()
	version, err := commonClient.GetMTVersion()
	if err != nil {
		log.Logger.Logln(logrus.ErrorLevel, err.Error())
		return
	}
	MtVersion = version
	sessionID, err := commonClient.GetSessionID()
	if err != nil {
		log.Logger.Logln(logrus.ErrorLevel, err.Error())
		return
	}
	SessionID = sessionID
}

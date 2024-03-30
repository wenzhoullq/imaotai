package common

import (
	"github.com/sirupsen/logrus"
	"zuoxingtao/init/config"
	"zuoxingtao/init/log"
)

var MtVersion string
var SessionID int
var FilterSet map[string]struct{}

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
	//初始化想要的茅台
	FilterSet = make(map[string]struct{})
	for _, v := range config.Config.FilterConfigure.ItemsCode {
		FilterSet[v] = struct{}{}
	}
	return nil
}

func UpdateCommon() {
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

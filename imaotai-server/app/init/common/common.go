package common

import (
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"imaotai_helper/constant"
	"imaotai_helper/init/config"
	"imaotai_helper/init/log"
)

var MtVersion string
var SessionID int
var FilterSet map[string]struct{}
var Node *snowflake.Node
var StatusMap map[int]string

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
	Node, err = snowflake.NewNode(0)
	if err != nil {
		return err
	}
	StatusMap = map[int]string{
		constant.USER_INIT:      "token待更新",
		constant.USER_NORMAL:    "正常",
		constant.User_SUSPENDED: "暂停",
		constant.USER_ABNORMAL:  "token已过期",
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

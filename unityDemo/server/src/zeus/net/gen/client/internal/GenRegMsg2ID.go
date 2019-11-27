package internal

import (
	"fmt"
	"strings"
	"zeus/net/gen/misc"
)

type CltRegMsg2ID struct {
	info string
}

func NewCltRegMsg() *CltRegMsg2ID {
	return &CltRegMsg2ID{info: kTemplateRegMsg2ID}
}

func (cm *CltRegMsg2ID) ReplaceSvcName(svcName string) {
	cm.info = strings.Replace(cm.info, "${SERVICE_NAME}", svcName, -1)
}

func (cm *CltRegMsg2ID) ReplaceFunctions(cfg *misc.Cfg) error {
	// Add deal function
	if err := cm.addFunctions(cfg.C2sMsgs); err != nil {
		return err
	}
	return nil
}

func (cm *CltRegMsg2ID) GetContent() string {
	return cm.info
}

func (cm *CltRegMsg2ID) addFunctions(C2sMsgs map[string]string) error {
	var err error
	var replaceStr string
	misc.OrderedForEach(C2sMsgs, func(key, value string) bool {
		_, typeName := misc.ParseMsgName(value)
		replaceStr += fmt.Sprintf("\tm.RegMsg2ID(&%s{}, %s)\n", typeName, key)
		return true
	})
	cm.info = strings.Replace(cm.info, "${REG_MSG2ID}", replaceStr, -1)
	return err
}

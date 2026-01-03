package util

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/zlog"
	"encoding/json"
	"strconv"
)

func Response(conn interfaces.Connection, response interface{}, messageType uint32) {
	marshal, err := json.Marshal(response)

	if err != nil {
		zlog.PrintErrorf("json.Marshal(response) %s ", err.Error())
		return
	}

	conn.SendMsg(messageType, marshal)
}

func UnitToString(number uint32) string {
	return strconv.Itoa(int(number))
}

func StringToUint32(str string) uint32 {
	atoi, _ := strconv.Atoi(str)
	return uint32(atoi)
}

package e

var MsgFlags = map[int]string{
	SUCCESS:                             "ok",
	ERROR:                               "fail",
	INVALID_PARAMS:                      "请求参数错误",
	ERROR_GET_MYSQL_CLUSTER_LIST_FAIL:   "获取所有mysql集群失败",
	ERROR_COUNT_MYSQL_CLUSTER_LIST_FAIL: "统计mysql集群失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

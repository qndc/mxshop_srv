package comm

import (
	"net/http"
	"time"
)

// UnixSecondToTime 秒级时间戳转time
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// UnixMilliToTime 毫秒级时间戳转time
func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(milli/1000, (milli%1000)*(1000*1000))
}

// UnixNanoToTime 纳秒级时间戳转time
func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(nano/(1000*1000*1000), nano%(1000*1000*1000))
}

// TransformTimestrToTimestamp 时间转换工具 将ISO 8601时间转为当地时间戳（13位 毫秒）
func TransformTimestrToTimestamp(timestr string) int64 {
	result, err := time.ParseInLocation(http.TimeFormat, timestr, time.Local)
	if err != nil {
		return -1
	}
	//转为13位时间戳,13位毫秒时间戳单位
	return result.Unix() * 1000
}

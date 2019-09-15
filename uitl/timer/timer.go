package timer

import "time"

const DEFAULT_FORMAT = "2006-01-02 15:04:05"

/*
 *	获取现在日期字符串
 *	@param string date
 *	@param string format
 *	@return string
 */
func GetNowDate(format ...string) string {
	tf := DEFAULT_FORMAT
	if len(format) == 1 {
		tf = format[0]
	}
	return time.Now().Format(tf)
}

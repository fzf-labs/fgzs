package timeutil

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// TimeLayout 常用日期格式化模板
var TimeLayout = "2006-01-02 15:04:05"

func Time() carbon.Carbon {
	return carbon.NewCarbon().SetTimezone(carbon.PRC)
}

func NowTime() time.Time {
	return carbon.Now().Carbon2Time()
}

func NowUnix() int64 {
	return carbon.Now().Timestamp()
}

// NowString 转换为当前时间 2021-06-29 23:53:32
func NowString() string {
	return carbon.Now().ToDateTimeString()
}

// NowMillisecondString 转换为当前时间 2021-06-29 23:53:32.010
func NowMillisecondString() string {
	now := carbon.Now()
	return now.ToDateTimeString() + "." + now.Format("u")
}

// NowMicrosecondString 转换为当前时间 2021-06-29 23:53:32.100000
func NowMicrosecondString() string {
	now := carbon.Now()

	return now.ToDateTimeString() + "." + strconv.Itoa(now.Microsecond())
}

// TimeToShortString 时间转日期
func TimeToShortString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format("2006.01.02")
}

func ToDateTimeStringByTime(ts time.Time) string {
	return carbon.Time2Carbon(ts).ToDateTimeString()
}

func ToDateTimeStringByTimePointer(ts *time.Time) string {
	if ts != nil {
		return carbon.Time2Carbon(*ts).ToDateTimeString()
	} else {
		return ""
	}
}

// GetShowTime 格式化人类友好时间
func GetShowTime(ts time.Time) string {
	duration := time.Now().Unix() - ts.Unix()
	timeStr := ""
	if duration < 60 {
		timeStr = "刚刚发布"
	} else if duration < 3600 {
		timeStr = fmt.Sprintf("%d分钟前更新", duration/60)
	} else if duration < 86400 {
		timeStr = fmt.Sprintf("%d小时前更新", duration/3600)
	} else if duration < 86400*2 {
		timeStr = "昨天更新"
	} else {
		timeStr = TimeToShortString(ts) + "前更新"
	}
	return timeStr
}

// Date 等同于PHP的date函数
// Date("Y-m-d H:i:s", time.Now())
func Date(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	tt := time.Now()
	if len(ts) > 0 {
		tt = ts[0]
	}
	if tt.Unix() <= 0 {
		return ""
	}
	return tt.Format(format)
}

// StrToTime 等同于PHP的strtotime函数
// StrToTime("2020-12-19 14:16:22")
func StrToTime(value string) (tt time.Time, err error) {
	if value == "" {
		err = errors.New("value is null")
		return
	}

	l, err := time.LoadLocation("Local")
	if err != nil {
		return
	}

	layouts := []string{
		"20060102",
		"20060102150405",
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, layout := range layouts {
		tt, err = time.ParseInLocation(layout, value, l)
		if err == nil {
			return
		}
	}
	return
}

// SubDays 计算日期相差多少天
// 返回值day>0, t1晚于t2; day<0, t1早于t2
func SubDays(t1, t2 time.Time) (day int) {
	swap := false
	if t1.Unix() < t2.Unix() {
		t_ := t1
		t1 = t2
		t2 = t_
		swap = true
	}

	day = int(t1.Sub(t2).Hours() / 24)

	// 计算在被24整除外的时间是否存在跨自然日
	if int(t1.Sub(t2).Milliseconds())%86400000 > int(86400000-t2.Unix()%86400000) {
		day += 1
	}

	if swap {
		day = -day
	}

	return
}

// StrToLocalTime 字符串转本地时间
func StrToLocalTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, errors.New("value is null")
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return StrToTime(value)
}

// DateFormat 格式化time.Time
// DateFormat("YYYY-MM-DD HH:mm:ss", time.Now())
func DateFormat(format string, tt time.Time) string {
	res := strings.Replace(format, "MM", tt.Format("01"), -1)
	res = strings.Replace(res, "M", tt.Format("1"), -1)
	res = strings.Replace(res, "DD", tt.Format("02"), -1)
	res = strings.Replace(res, "D", tt.Format("2"), -1)
	res = strings.Replace(res, "YYYY", tt.Format("2006"), -1)
	res = strings.Replace(res, "YY", tt.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", tt.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", tt.Hour()), -1)
	res = strings.Replace(res, "hh", tt.Format("03"), -1)
	res = strings.Replace(res, "h", tt.Format("3"), -1)
	res = strings.Replace(res, "mm", tt.Format("04"), -1)
	res = strings.Replace(res, "m", tt.Format("4"), -1)
	res = strings.Replace(res, "ss", tt.Format("05"), -1)
	res = strings.Replace(res, "s", tt.Format("5"), -1)
	return res
}

// GetDaysAgoZeroTime 以当天0点为基准，获取前后某天0点时间
// 昨天：GetDaysAgoZeroTime(-1)
// 今天：GetDaysAgoZeroTime(0)
// 明天：GetDaysAgoZeroTime(1)
func GetDaysAgoZeroTime(day int) time.Time {
	date := time.Now().AddDate(0, 0, day).Format("2006-01-02")
	tt, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return tt
}

// TimeToHuman 根据时间戳获得人类可读时间
func TimeToHuman(ts int) string {
	var res = ""
	if ts == 0 {
		return res
	}

	tt := int(time.Now().Unix()) - ts
	data := [7]map[string]interface{}{
		{"key": 31536000, "value": "年"},
		{"key": 2592000, "value": "个月"},
		{"key": 604800, "value": "星期"},
		{"key": 86400, "value": "天"},
		{"key": 3600, "value": "小时"},
		{"key": 60, "value": "分钟"},
		{"key": 1, "value": "秒"},
	}
	for _, v := range data {
		var c = tt / v["key"].(int)
		if 0 != c {
			suffix := "前"
			if c < 0 {
				suffix = "后"
				c = -c
			}
			res = strconv.Itoa(c) + v["value"].(string) + suffix
			break
		}
	}

	return res
}

// 计算当前时间至当天 23:59:59 的时间差 (秒)
func SameDaySubSecond() int {
	currentTime := time.Now()
	t1 := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix()
	sec := time.Unix(t1, 0).Sub(currentTime)
	return int(sec.Seconds())
}

// 计算某一日期与当前日期天数差
func SubNowDays(tt time.Time) int64 {
	days := (time.Now().Unix() - tt.Unix()) / 86400
	return days
}

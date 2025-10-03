package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 判断给定时间戳是否在时间范围内（支持跨天）
// 输入：startStr, endStr 格式为 "hh:mm"，timestamp 为 Unix 时间戳（秒）
// 返回：是否在范围内，错误信息
func IsTimeInRange(startStr, endStr string, timestamp int64) (bool, error) {
	// 解析开始时间
	start, err := parseTime(startStr)
	if err != nil {
		return false, fmt.Errorf("invalid start time: %w", err)
	}

	// 解析结束时间
	end, err := parseTime(endStr)
	if err != nil {
		return false, fmt.Errorf("invalid end time: %w", err)
	}

	// 转换时间戳为本地时间
	t := time.Unix(timestamp, 0).Local()
	target := time.Date(0, 1, 1, t.Hour(), t.Minute(), 0, 0, time.UTC)

	// 处理跨天情况（结束时间小于开始时间）
	if end.Before(start) {
		// 情况1：当天时间段（target >= start）
		if target.After(start) || target.Equal(start) {
			return true, nil
		}
		// 情况2：次日时间段（target <= end）
		return target.Before(end) || target.Equal(end), nil
	}

	// 处理非跨天情况
	return (target.After(start) || target.Equal(start)) &&
		(target.Before(end) || target.Equal(end)), nil
}

// 解析 "hh:mm" 格式的时间
func parseTime(s string) (time.Time, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return time.Time{}, errors.New("invalid format")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil || hour < 0 || hour > 23 {
		return time.Time{}, errors.New("invalid hour")
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil || minute < 0 || minute > 59 {
		return time.Time{}, errors.New("invalid minute")
	}

	return time.Date(0, 1, 1, hour, minute, 0, 0, time.UTC), nil
}

// 获取星期几（星期一=1，星期日=7）
// 可选参数 timestamp：Unix 时间戳（秒），默认使用当前时间
func GetWeekday(timestamp ...int64) int {
	var t time.Time
	if len(timestamp) > 0 {
		t = time.Unix(timestamp[0], 0)
	} else {
		t = time.Now()
	}

	// Go 中 Sunday=0, Monday=1,..., Saturday=6
	// 转换为 Monday=1, Sunday=7
	weekday := t.Weekday()
	if weekday == time.Sunday {
		return 7
	}
	return int(weekday)
}

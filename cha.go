package cha

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

func onlyASCII(s string) bool {
	for _, c := range s {
		if c >= 255 {
			return false
		}
	}
	return true
}

var cjkRanges = []string{
	"3300-33FF",   // CJK Compatibility
	"4E00-9FFF",   // CJK Unified Ideographs
	"3400-4DBF",   // CJK Unified Ideographs Extension A
	"F900-FAFF",   // CJK Compatibility Ideographs
	"FE30-FE4F",   // CJK Compatibility Forms
	"20000-215FF", // CJK Unified Ideographs Extension B
	"21600-230FF", // ^
	"23100-245FF", // ^
	"24600-260FF", // ^
	"26100-275FF", // ^
	"27600-290FF", // ^
	"29100-2A6DF", // ^
	"2A700-2B73F", // CJK Unified Ideographs Extension C
	"2B740-2B81F", // CJK Unified Ideographs Extension D
	"2B820-2CEAF", // CJK Unified Ideographs Extension E
	"2CEB0-2EBEF", // CJK Unified Ideographs Extension F
	"2F800-2FA1F", // CJK Compatibility Ideographs Supplement
}

// ["A-B", "C-D"] =hex->int=> [10,11,12,13]
func parseCJKIntRanges() []uint64 {
	var CJKIntRanges = make([]uint64, 0)
	for _, CJKRange := range cjkRanges {
		r := strings.Split(CJKRange, "-")
		start, _ := strconv.ParseUint(r[0], 16, 32)
		end, _ := strconv.ParseUint(r[1], 16, 32)
		CJKIntRanges = append(CJKIntRanges, start, end)
	}
	return CJKIntRanges
}

// https://en.wikipedia.org/wiki/CJK_Unified_Ideographs
func onlyCJK(s string) bool {
	ranges := parseCJKIntRanges()
	count := 0
	for i := 0; i < len(ranges); i += 2 {
		start := ranges[i]
		end := ranges[i+1]
		for _, c := range s {
			cn := uint64(c)
			if cn >= start && cn <= end {
				count++
			}
		}
	}
	return utf8.RuneCountInString(s) == count
}

// NaEr 返回传入内容中应该插入空格的位置数组
func NaEr(s io.Reader) ([]int, error) {
	if s == nil {
		return nil, fmt.Errorf("输入内容不能为 nil")
	}
	b, err := ioutil.ReadAll(s)
	if err != nil {
		return nil, fmt.Errorf("读取字符内容失败, %v", err)
	}
	str := string(b)

	// 01.全是英文
	if onlyASCII(str) {
		return nil, nil
	}

	// 02.全是中文
	if onlyCJK(str) {
		return nil, nil
	}

	return nil, fmt.Errorf("啥都没干")
}

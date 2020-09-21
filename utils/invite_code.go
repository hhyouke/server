package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	baseRunes = "HVE8S2DZX9C7P5IK3MJUAR4WYLTN6BGQ"
	scaling   = int64(32)
	padRune   = "F"
	length    = 6
)

// InvCode instance if invite code
type InvCode struct {
	base    string // 进制的包含字符, string类型
	decimal int64  // 进制长度
	pad     string // 补位字符,若生成的code小于最小长度,则补位+随机字符, 补位字符不能在进制字符中
	len     int    // code最小长度
}

// NewInvCode create new invite code(32) instance
func NewInvCode() *InvCode {
	return &InvCode{
		base:    baseRunes,
		decimal: scaling,
		pad:     padRune,
		len:     length,
	}
}

// NewCustomInvCode create new invite code(32) instance with custom options
// pad should not be in base!!
func NewCustomInvCode(base, pad string) *InvCode {
	return &InvCode{
		base:    base,
		decimal: int64(32),
		pad:     pad,
		len:     length,
	}
}

// IDToCode create invite code from int64 id
func (c *InvCode) IDToCode(id int64) string {
	mod := int64(0)
	res := ""

	for id != 0 {
		mod = id % c.decimal
		id = id / c.decimal
		res += string(c.base[mod])
	}

	resLen := len(res)
	if resLen < c.len {
		res += c.pad
		for i := 0; i < c.len-resLen-1; i++ {
			rand.Seed(time.Now().UnixNano())
			res += string(c.base[rand.Intn(int(c.decimal))])
		}
	}

	return res
}

// CodeToID restore int64 id from an invite code
func (c *InvCode) CodeToID(code string) int64 {
	res := int64(0)

	lenCode := len(code)
	baseArr := []byte(c.base)     // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}

	// 查找补位字符的位置
	isPad := strings.Index(code, c.pad)
	if isPad != -1 {
		lenCode = isPad
	}

	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == c.pad {
			continue
		}
		index := baseRev[code[i]]
		//fmt.Printf("index: %v\n", index)
		b := int64(1)
		for j := 0; j < r; j++ {
			b *= c.decimal
			//fmt.Printf("b: %v %v\n", index, b)
		}
		res += int64(index) * b
		//fmt.Printf("res: %v\n", res)
		r++
	}
	return res
}

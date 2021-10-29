// SPDX-License-Identifier: MIT

// Package gb32100 统一信用代码校验
//
// GB32100—2015
package gb32100

import "errors"

// ErrInvalidFormat 格式错误
var ErrInvalidFormat = errors.New("无效的格式")

var (
	ministries = map[byte]string{
		'1': "机构编制",
		'5': "民政",
		'9': "工商",
		'Y': "其他",
	}

	types = map[byte]map[byte]string{
		'1': {
			'1': "机关",
			'2': "事业单位",
			'3': "中央编办直接管理机构编制的群众团体",
			'9': "其他",
		},
		'5': {
			'1': "社会团体",
			'2': "民办非企业单位",
			'3': "基金会",
			'9': "其他",
		},
		'9': {
			'1': "企业",
			'2': "个体工商户",
			'3': "农民专业合作社",
		},
		'Y': {
			'1': "其它",
		},
	}
)

type GB32100 struct {
	Raw      string
	Ministry byte   // 登记管理部门
	Type     byte   // 类别
	Region   string // 区域信息，可参考 https://github.com/issue9/cnregion
	ID       string // 主体代码
}

func Parse(bs string) (*GB32100, error) {
	if !IsValid([]byte(bs)) {
		return nil, ErrInvalidFormat
	}

	return &GB32100{
		Raw:      bs,
		Ministry: bs[0],
		Type:     bs[1],
		Region:   bs[2:8],
		ID:       bs[8:17],
	}, nil
}

func (g *GB32100) MinistryName() string { return ministries[g.Ministry] }

func (g *GB32100) TypeName() string { return types[g.Ministry][g.Type] }

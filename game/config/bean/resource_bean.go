package bean

import "sync"

var (
	instance ResourceBeanHolder
	once     sync.Once
)

type ResourceBeanHolder struct {
	SkillResource SkillResource
	BuffResource  BuffResource
	ConfigBean    map[string]map[string]interface{}
}

type SkillResource struct {
	Id    string // id
	Name  string // 名称
	Level int32  // 描述
	// 额外字段
	ExaString1 string
	ExaString2 string
}

type BuffResource struct {
	Id       string // id
	BuffType int32  // buff类型
	// 额外字段
	ExaString1 string
	ExaString2 string
}

func GetInstance() ResourceBeanHolder {
	once.Do(func() {
		instance = ResourceBeanHolder{
			SkillResource{},
			BuffResource{},
			make(map[string]map[string]interface{}),
		}
	})
	return instance
}

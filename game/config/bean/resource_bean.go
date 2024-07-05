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

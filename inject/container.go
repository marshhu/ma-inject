package inject

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var injectTagName = "inject" //依赖注入tag名

//生命周期
// singleton：单例 单一实例，每次使用都是该实例
// transient:瞬时实例,每次使用都创建新的实例
type Container struct {
	sync.Mutex
	singletons map[string]interface{}
	transients map[string]factory
}

type factory = func() (interface{}, error)

//注册单例对象
func (c *Container) SetSingleton(name string, singleton interface{}) {
	c.Lock()
	c.singletons[name] = singleton
	c.Unlock()
}

func (c *Container) GetSingleton(name string) interface{} {
	return c.singletons[name]
}

//注册瞬时实例创建工厂方法
func (c *Container) SetTransient(name string, factory factory) {
	c.Lock()
	c.transients[name] = factory
	c.Unlock()
}

func (c *Container) GetTransient(name string) interface{} {
	factory := c.transients[name]
	instance, _ := factory()
	return instance
}

//注入实例
func (c *Container) Entry(instance interface{}) error {
	err := c.entryValue(reflect.ValueOf(instance))
	if err != nil {
		return err
	}
	return nil
}

func (c *Container) entryValue(value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return errors.New("必须为指针")
	}
	elemType, elemValue := value.Type().Elem(), value.Elem()
	for i := 0; i < elemType.NumField(); i++ {
		if !elemValue.Field(i).CanSet() { //不可设置 跳过
			continue
		}

		fieldType := elemType.Field(i)
		if fieldType.Anonymous {
			//fmt.Println(fieldType.Name + "是匿名字段")
			item := reflect.New(elemValue.Field(i).Type())
			c.entryValue(item) //递归注入
			elemValue.Field(i).Set(item.Elem())
		} else {
			if elemValue.Field(i).IsZero() { //零值才注入
				//fmt.Println(elemValue.Field(i).Interface())
				//fmt.Println(fieldType.Name)
				tag := fieldType.Tag.Get(injectTagName)
				injectInstance, err := c.getInstance(tag)
				if err != nil {
					return err
				}
				c.entryValue(reflect.ValueOf(injectInstance)) //递归注入

				elemValue.Field(i).Set(reflect.ValueOf(injectInstance))
			} else {
				fmt.Println(fieldType.Name)
			}
		}
	}
	return nil
}

func (c *Container) getInstance(tag string) (interface{}, error) {
	var injectName string
	tags := strings.Split(tag, ",")
	if len(tags) == 0 {
		injectName = ""
	} else {
		injectName = tags[0]
	}

	if c.isTransient(tag) {
		factory, ok := c.transients[injectName]
		if !ok {
			return nil, errors.New("transient factory not found")
		} else {
			return factory()
		}
	} else { //默认单例
		instance, ok := c.singletons[injectName]
		if !ok || instance == nil {
			return nil, errors.New(injectName + " dependency not found")
		} else {
			return instance, nil
		}
	}
}

// transient:瞬时实例,每次使用都创建新的实例
func (c *Container) isTransient(tag string) bool {
	tags := strings.Split(tag, ",")
	for _, name := range tags {
		if name == "transient" {
			return true
		}
	}
	return false
}

func (c *Container) String() string {
	lines := make([]string, 0, len(c.singletons)+len(c.transients)+2)
	lines = append(lines, "singletons:")
	for key, value := range c.singletons {
		line := fmt.Sprintf("	%s: %x %s", key, c.singletons[key], reflect.TypeOf(value).String())
		lines = append(lines, line)
	}

	lines = append(lines, "transients:")
	for key, value := range c.transients {
		line := fmt.Sprintf("	%s: %x %s", key, c.transients[key], reflect.TypeOf(value).String())
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

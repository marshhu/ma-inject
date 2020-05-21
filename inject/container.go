package inject

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Container struct {
	sync.Mutex
	singletons map[string]interface{}
	factories  map[string]factory
}

type factory = func() (interface{}, error)

//注册单例对象
func (c *Container) SetSingleton(name string, singleton interface{}) {
	c.Lock()
	c.singletons[name] = singleton
	c.Unlock()
}

//获取单例对象
func (c *Container) GetSingleton(name string) interface{} {
	return c.singletons[name]
}

//设置实例对象工厂
func (c *Container) SetProtoType(name string, factory factory) {
	c.Lock()
	c.factories[name] = factory
	c.Unlock()
}

//获取实例对象
func (c *Container) GetProtoType(name string) (interface{}, error) {
	factory, ok := c.factories[name]
	if !ok {
		return nil, errors.New("factory not found")
	}
	return factory()
}

//注入依赖
func (c *Container) Entry(instance interface{}) error {
	if reflect.TypeOf(instance).Kind() != reflect.Ptr {
		return errors.New("必须为指针")
	}
	c.entryValue(reflect.ValueOf(instance))
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
			fmt.Println(fieldType.Name + "是匿名字段")
			item := reflect.New(elemValue.Field(i).Type())
			c.entryValue(item) //递归注入
			elemValue.Field(i).Set(item.Elem())
		} else {
			fmt.Println(fieldType.Name)
			tag := fieldType.Tag.Get("inject")
			diName := c.injectName(tag)
			if diName == "" {
				continue
			}

			var (
				diInstance interface{}
				err        error
			)
			if c.isSingleton(tag) {
				diInstance = c.GetSingleton(diName)
			}
			if c.isProtoType(tag) {
				diInstance, err = c.GetProtoType(diName)
			}

			if err != nil {
				return err
			}
			if diInstance == nil {
				return errors.New(diName + "dependency not found")
			}
			c.entryValue(reflect.ValueOf(diInstance)) //递归注入

			elemValue.Field(i).Set(reflect.ValueOf(diInstance))
		}
	}
	return nil
}

func (c *Container) String() string {
	lines := make([]string, 0, len(c.singletons)+len(c.factories)+2)
	lines = append(lines, "singletons:")
	for key, value := range c.singletons {
		line := fmt.Sprintf("	%s: %x %s", key, c.singletons[key], reflect.TypeOf(value).String())
		lines = append(lines, line)
	}

	lines = append(lines, "factories:")
	for key, value := range c.factories {
		line := fmt.Sprintf("	%s: %x %s", key, c.singletons[key], reflect.TypeOf(value).String())
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (c *Container) injectName(tag string) string {
	tags := strings.Split(tag, ",")
	if len(tags) == 0 {
		return ""
	}
	return tags[0]
}

//检测是否是实例依赖
func (c *Container) isProtoType(tag string) bool {
	tags := strings.Split(tag, ",")
	for _, name := range tags {
		if name == "prototype" {
			return true
		}
	}
	return false
}

func (c *Container) isSingleton(tag string) bool {
	return !c.isProtoType(tag)
}

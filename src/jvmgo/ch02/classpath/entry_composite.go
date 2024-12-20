package classpath

import (
	"errors"
	"strings"
)

// 定义CompositeEntry是一个切片
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	//遍历boot引导器下所有的entry入口。看这些jar包中是否包含className
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			//如果没有异常
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}

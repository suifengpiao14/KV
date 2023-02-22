package kv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type KV struct {
	Key   string
	Value string
}

type KVS []KV

func (kvs KVS) Json() (jsonStr string) {
	var err error
	for _, kv := range kvs {
		value := strings.TrimSpace(kv.Value)
		if len(value) > 0 && (value[0] == '{' || value[0] == '[') && gjson.Valid(value) {
			jsonStr, err = sjson.SetRaw(jsonStr, kv.Key, kv.Value)
			if err != nil {
				panic(err)
			}
			continue
		}
		jsonStr, err = sjson.Set(jsonStr, kv.Key, kv.Value)
		if err != nil {
			panic(err)
		}
	}
	return jsonStr
}

// GetIndex 获取相同前置后带数字名称的最大数字,加1后作为新元素下标,返回
func (kvs KVS) GetNextIndex(keyPrefix string, keySeparator string) (maxIndex int) {
	maxIndex = -1
	for _, kv := range kvs {
		if strings.HasPrefix(kv.Key, keyPrefix) && len(kv.Key) > len(keyPrefix) {
			numberStr := kv.Key[len(keyPrefix):]
			numberStr = strings.Trim(numberStr, keySeparator)
			dotIndex := strings.Index(numberStr, keySeparator)
			if dotIndex > -1 {
				numberStr = numberStr[:dotIndex]
			}
			if index, err := strconv.Atoi(numberStr); err == nil && maxIndex < index {
				maxIndex = index
			}

		}
	}
	maxIndex++ // 增加1,作为新元素下标
	return maxIndex
}

func (kvs KVS) Exists(key string) (exists bool) {
	for _, kv := range kvs {
		if key == kv.Key {
			return true
		}
	}
	return false
}

func (kvs KVS) GetFirstByKey(key string) (kv KV, index int) {
	index = -1
	for i, tplKV := range kvs {
		if key == tplKV.Key {
			return tplKV, i
		}
	}
	return kv, index
}

func (kvs KVS) GetByIndex(index int) (kv KV, exists bool) {
	if index > len(kvs)-1 || index < 0 {
		return kv, false
	}
	kv = kvs[index]
	return kv, true
}

//Order 对kv 集合排序
func (kvs KVS) Order(keyOrder []string) (orderedKVS KVS) {
	orderedKVS = make(KVS, 0)
	orderIndex := make([]int, 0)
	// 确定顺序
	for _, key := range keyOrder {
		kv, index := kvs.GetFirstByKey(key)
		if index < 0 {
			continue
		}
		orderIndex = append(orderIndex, index)
		orderedKVS = append(orderedKVS, kv)
	}

	if len(orderIndex) == len(kvs) {
		return orderedKVS
	}
	//复制剩余kv
	for i, kv := range kvs {
		for _, index := range orderIndex {
			if i == index {
				continue
			}
		}
		orderedKVS = append(orderedKVS, kv)
	}

	return orderedKVS
}

// AddIgnore 引用解析到的kv，批量添加
func (kvs *KVS) AddIgnore(addkvs ...KV) {
	for _, addKv := range addkvs {
		for _, existsKv := range *kvs {
			if existsKv.Key == addKv.Key {
				continue
			}
		}
		*kvs = append(*kvs, addKv)
	}
}

// AddReplace 模板解析后获取的kv，批量新增/替换
func (kvs *KVS) AddReplace(replacekvs ...KV) {
	for _, addKv := range replacekvs {
		exists := false
		for i, existsKv := range *kvs {
			if existsKv.Key == addKv.Key {
				(*kvs)[i] = addKv
				exists = true
				break
			}
		}
		if !exists {
			*kvs = append(*kvs, addKv)
		}
	}
}

// ReplacePrefix 引用解析获得的新数据，需要批量替换id前缀
func (kvs *KVS) ReplacePrefix(old, new string) {
	for i, kv := range *kvs {
		if strings.HasPrefix(kv.Key, old) {
			kv.Key = fmt.Sprintf("%s%s", new, kv.Key[len(old):])
			(*kvs)[i] = kv
		}
	}
}

// FillterByPrefix 引用解析获得的新数据，获取指定前缀kv
func (kvs *KVS) FillterByPrefix(prefix string) (newKVs KVS) {
	newKVs = KVS{}
	for _, kv := range *kvs {
		if strings.HasPrefix(kv.Key, prefix) {
			newKVs = append(newKVs, kv)
		}
	}
	return newKVs
}

//JsonToKVS 将json 转换为key->value 对,key 的规则为github.com/tidwall/gjson 的path
func JsonToKVS(jsonStr string) (kvs KVS) {
	kvs = make(KVS, 0)
	paths := make([]string, 0)
	result := gjson.Parse(jsonStr)
	allResult := getAllJsonResult(result)
	for _, result := range allResult {
		subPath := result.Path(jsonStr)
		paths = append(paths, subPath)
	}
	for _, path := range paths {
		kv := KV{
			Key:   path,
			Value: result.Get(path).String(),
		}
		kvs = append(kvs, kv)
	}
	return kvs
}

func getAllJsonResult(result gjson.Result) (allResult []gjson.Result) {
	allResult = make([]gjson.Result, 0)
	result.ForEach(func(key, value gjson.Result) bool {
		if !value.IsArray() && !value.IsObject() {
			allResult = append(allResult, value)
		} else {
			subAllResult := getAllJsonResult(value)
			allResult = append(allResult, subAllResult...)
		}
		return true
	})
	return
}

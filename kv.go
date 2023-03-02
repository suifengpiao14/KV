package kvstruct

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	KV_TYPE_STRING  = "string"
	KV_TYPE_INT     = "int"
	KV_TYPE_BOOLEAN = "boolean"
	KV_TYPE_FLOAT   = "float"
	KV_TYPE_JSON    = "json"
)

type KV struct {
	Type  string
	Key   string
	Value string
}

type KVS []KV

//Deprecated
func (kvs KVS) Json(WithType bool) (jsonStr string, err error) {
	for _, kv := range kvs {
		// 任何情况,都处理特殊处理json和boolean 类型
		if kv.Type == KV_TYPE_JSON {
			jsonStr, err = sjson.SetRaw(jsonStr, kv.Key, kv.Value)
			if err != nil {
				return "", err
			}
			continue
		}

		strValue := kv.Value
		if kv.Type == KV_TYPE_BOOLEAN {
			switch strings.ToLower(strValue) {
			case "是", "对", "1", "yes":
				strValue = "true"
			case "否", "错", "0", "no":
				strValue = "false"
			}
		}
		if !WithType {
			jsonStr, err = sjson.Set(jsonStr, kv.Key, strValue)
			if err != nil {
				return "", err
			}
			return jsonStr, nil
		}

		var value interface{}
		value = strValue
		switch kv.Type {
		case KV_TYPE_BOOLEAN:
			value, err = strconv.ParseBool(strValue)
			if err != nil {
				return "", err
			}
		case KV_TYPE_INT:
			value, err = strconv.Atoi(strValue)
			if err != nil {
				return "", err
			}
		case KV_TYPE_FLOAT:
			value, err = strconv.ParseFloat(strValue, 64)
			if err != nil {
				return "", err
			}
		}
		jsonStr, err = sjson.Set(jsonStr, kv.Key, value)
		if err != nil {
			return "", err
		}
	}
	return jsonStr, nil
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

// Order 对kv 集合排序
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
		notExists := true
		for _, index := range orderIndex {
			if i == index {
				notExists = false
				break
			}
		}
		if notExists {
			orderedKVS = append(orderedKVS, kv)
		}
	}

	return orderedKVS
}

// Add 新增,不排除重复
func (kvs *KVS) Add(addkvs ...KV) {
	*kvs = append(*kvs, addkvs...)
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

// JsonToKVS 将json 转换为key->value 对,key 的规则为github.com/tidwall/gjson 的path
func JsonToKVS(jsonStr string, namespace string) (kvs KVS) {
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
			Key:   fmt.Sprintf("%s.%s", strings.Trim(namespace, "."), strings.Trim(path, ".")),
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

// IsJsonStr 判断是否为json字符串
func IsJsonStr(str string) (yes bool) {
	str = strings.TrimSpace(str)
	yes = len(str) > 0 && (str[0] == '{' || str[0] == '[') && gjson.Valid(str)
	return yes

}

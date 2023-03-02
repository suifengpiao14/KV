package kvstruct

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKVSJson(t *testing.T) {
	kv := KV{
		Type:  KV_TYPE_BOOLEAN,
		Key:   "doc.example.response.200.language",
		Value: `是`,
	}
	kvs := make(KVS, 0)
	kvs = append(kvs, kv)
	jsonStr, err := kvs.Json(true)
	require.NoError(t, err)
	fmt.Println(jsonStr)
}

func TestJsonToKVS(t *testing.T) {
	jsonstr := `{"ad":{"advertise":{"id":{"database":"ad","table":"advertise","name":"id","goType":"int","dbType":"int(11)","comment":"主键","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"false","size":"11"},"advertiser_id":{"database":"ad","table":"advertise","name":"advertiser_id","goType":"string","dbType":"varchar(32)","comment":"广告主","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"title":{"database":"ad","table":"advertise","name":"title","goType":"string","dbType":"varchar(32)","comment":"广告标题","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"begin_at":{"database":"ad","table":"advertise","name":"begin_at","goType":"string","dbType":"datetime","comment":"投放开��时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},"end_at":{"database":"ad","table":"advertise","name":"end_at","goType":"string","dbType":"datetime","comment":"投放结束时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},"summary":{"database":"ad","table":"advertise","name":"summary","goType":"string","dbType":"varchar(128)","comment":"广告素材-文字描述","nullable":"true","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"128"},"image":{"database":"ad","table":"advertise","name":"image","goType":"string","dbType":"varchar(256)","comment":"广告素材-图片地址","nullable":"true","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},"link":{"database":"ad","table":"advertise","name":"link","goType":"string","dbType":"varchar(512)","comment":"连接地址","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"512"},"remark":{"database":"ad","table":"advertise","name":"remark","goType":"string","dbType":"varchar(255)","comment":"备注","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"255"},"type":{"database":"ad","table":"advertise","name":"type","goType":"string","dbType":"enum('text','image','vido')","comment":"广告素材(类型),text-文字,image-图片,vido-视频","nullable":"false","enums":"text,image,vido","autoIncrement":"false","default":"text","onUpdate":"false","unsigned":"false","size":"0"},"value_obj":{"database":"ad","table":"advertise","name":"value_obj","goType":"string","dbType":"varchar(1024)","comment":"json扩展,广告的值属性对象","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"1024"},"created_at":{"database":"ad","table":"advertise","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"ad","table":"advertise","name":"updated_at","goType":"string","dbType":"datetime","comment":"修改时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},"deleted_at":{"database":"ad","table":"advertise","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"}},"window":{"id":{"database":"ad","table":"window","name":"id","goType":"int","dbType":"int(11) unsigned","comment":"主键","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"true","size":"11"},"code":{"database":"ad","table":"window","name":"code","goType":"string","dbType":"varchar(32)","comment":"位置编码","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"title":{"database":"ad","table":"window","name":"title","goType":"string","dbType":"varchar(32)","comment":"位置名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"remark":{"database":"ad","table":"window","name":"remark","goType":"string","dbType":"varchar(255)","comment":"位置描述","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"255"},"content_types":{"database":"ad","table":"window","name":"content_types","goType":"string","dbType":"varchar(50)","comment":"广告素材(类型),text-文字,image-图片,vido-视频,多个逗号分隔","nullable":"true","enums":"","autoIncrement":"false","default":"text","onUpdate":"false","unsigned":"false","size":"50"},"width":{"database":"ad","table":"window","name":"width","goType":"int","dbType":"smallint(6)","comment":"橱窗宽度","nullable":"true","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"6"},"high":{"database":"ad","table":"window","name":"high","goType":"int","dbType":"smallint(6)","comment":"橱窗高度","nullable":"true","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"6"},"created_at":{"database":"ad","table":"window","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"ad","table":"window","name":"updated_at","goType":"string","dbType":"datetime","comment":"修改时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},"deleted_at":{"database":"ad","table":"window","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"}},"window_advertise":{"id":{"database":"ad","table":"window_advertise","name":"id","goType":"int","dbType":"int(11)","comment":"主键","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"false","size":"11"},"code":{"database":"ad","table":"window_advertise","name":"code","goType":"string","dbType":"varchar(32)","comment":"橱窗编码","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"advertise_id":{"database":"ad","table":"window_advertise","name":"advertise_id","goType":"int","dbType":"int(11)","comment":"广告ID","nullable":"false","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"11"},"advertise_priority":{"database":"ad","table":"window_advertise","name":"advertise_priority","goType":"int","dbType":"int(11)","comment":"广告优先级(同一个橱窗有多个广告时,按照优先级展示)","nullable":"true","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"11"},"created_at":{"database":"ad","table":"window_advertise","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"ad","table":"window_advertise","name":"updated_at","goType":"string","dbType":"datetime","comment":"修改时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},"deleted_at":{"database":"ad","table":"window_advertise","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"}}}}`
	kvs := JsonToKVS(jsonstr, "root")
	fmt.Println(kvs)

}

func TestGetNextIndex(t *testing.T) {
	keySeparator := "."
	t.Run("end number", func(t *testing.T) {
		prefix := "doc.request.parameter."
		kvs := KVS{
			KV{Key: fmt.Sprintf("%s2", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 3, nextIndex)
	})

	t.Run("middle number", func(t *testing.T) {
		prefix := "doc.request.parameter."
		kvs := KVS{
			KV{Key: fmt.Sprintf("%s2.hahhah", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 3, nextIndex)
	})
	t.Run("star number", func(t *testing.T) {
		prefix := "d"
		kvs := KVS{
			KV{Key: fmt.Sprintf("%s2.hahhah", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 3, nextIndex)
	})
	t.Run("no number", func(t *testing.T) {
		prefix := "d"
		kvs := KVS{
			KV{Key: fmt.Sprintf("%s2eigj.hahhah", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 0, nextIndex)
	})

}

func TestOrder(t *testing.T) {
	kvs := KVS{
		{Key: "a", Value: "value_a"},
		{Key: "c", Value: "value_c"},
		{Key: "b", Value: "value_b"},
	}
	keyOrder := []string{"b", "a", "d"}
	expectedKeys := []string{"b", "a", "c"}
	orderedKVS := kvs.Order(keyOrder)
	ok := len(orderedKVS) == len(expectedKeys)
	if ok {
		for i, kv := range orderedKVS {
			key := expectedKeys[i]
			if key != kv.Key {
				ok = false
			}
		}
	}
	assert.Equal(t, true, ok)
}

package json

import (
	"encoding/json"
	"errors"
	"github.com/bitly/go-simplejson"
	"strings"
)

/**
 * json 编码
 * @param interface v
 * @return slice, error
 */
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

/**
 * json解码
 * @param interface v
 * @return obj, error
 */
func Unmarshal(body []byte) (*simplejson.Json, error) {
	return simplejson.NewJson(body)
}

/**
 * 获取解码后key对应的obj
 * @param obj j
 * @param string key
 * @return obj
 */
func Get(j *simplejson.Json, key string) *simplejson.Json {
	return j.Get(key)
}

/**
 * 获取解码后key对应的字符串
 * @param obj j
 * @param string key
 * @return string
 */
func GetString(j *simplejson.Json, key string) string {
	tmStr, _ := j.Get(key).String()
	return tmStr
}

/**
 * 获取解码后key对应的uint64
 * @param obj j
 * @param string key
 * @return uint64
 */
func GetUInt64(j *simplejson.Json, key string) uint64 {
	tmUint64, _ := j.Get(key).Uint64()
	return tmUint64
}

/**
 * 获取解码后key对应的int64
 * @param obj j
 * @param string key
 * @return int64
 */
func GetInt64(j *simplejson.Json, key string) int64 {
	tmInt64, _ := j.Get(key).Int64()
	return tmInt64
}

/**
 * 获取解码后key对应的float64
 * @param obj j
 * @param string key
 * @return float64
 */
func GetFloat64(j *simplejson.Json, key string) float64 {
	tmFloat64, _ := j.Get(key).Float64()
	return tmFloat64
}

/**
 * 获取解码后key对应的slice
 * @param obj j
 * @param string key
 * @return slice
 */
func GetBytes(j *simplejson.Json, key string) []byte {
	tmBytes, _ := j.Get(key).Bytes()
	return tmBytes
}

/**
 * 获取解码后key对应的bool
 * @param obj j
 * @param string key
 * @return bool
 */
func GetBool(j *simplejson.Json, key string) bool {
	tmBool, _ := j.Get(key).Bool()
	return tmBool
}

/**
 * 获取解码后key对应的obj
 * @param obj j
 * @param string key
 * @return obj
 */
func GetObj(j *simplejson.Json, key string) (*simplejson.Json, error) {
	if js, is := j.CheckGet(key); is {
		return js, nil
	}

	return nil, errors.New("get " + key + " data failed")
}

/**
 * 获取解码后key对应的string
 * @param obj j
 * @param string key
 * @return string, error
 */
func GetStringVal(j *simplejson.Json, key string) (string, error) {
	tmpStr, err := j.Get(key).String()

	return strings.TrimSpace(tmpStr), err
}

/**
 * 获取解码后key对应的uint64
 * @param obj j
 * @param string key
 * @return uint64, error
 */
func GetUInt64Val(j *simplejson.Json, key string) (uint64, error) {
	return j.Get(key).Uint64()
}

/**
 * 获取解码后key对应的int64
 * @param obj j
 * @param string key
 * @return int64, error
 */
func GetInt64Val(j *simplejson.Json, key string) (int64, error) {
	return j.Get(key).Int64()
}

/**
 * 获取解码后key对应的float64
 * @param obj j
 * @param string key
 * @return float64, error
 */
func GetFloat64Val(j *simplejson.Json, key string) (float64, error) {
	return j.Get(key).Float64()
}

/**
 * 获取解码后key对应的slice
 * @param obj j
 * @param string key
 * @return slice, error
 */
func GetBytesVal(j *simplejson.Json, key string) ([]byte, error) {
	return j.Get(key).Bytes()
}

/**
 * 获取解码后key对应的bool
 * @param obj j
 * @param string key
 * @return bool, error
 */
func GetBoolVal(j *simplejson.Json, key string) (bool, error) {
	return j.Get(key).Bool()
}

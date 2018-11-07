package hash

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

//无链表的散列表
type NoLinkHashTable struct {
	Buckets  []Bucket
	Capacity int32
	Used     int32
}

//桶
type Bucket struct {
	HKey string
	Data int
}

func NolinkHashTable(length int32) {
	//新建算列表
	ht := NoLinkHashTable{}
	ht.Capacity = length
	ht.Buckets = make([]Bucket, length)
	for i := 0; i < 5; i++ {
		ht.hSet("abc", i)
	}
	fmt.Println(ht)
	one := ht.hGet("abc")
	fmt.Println(one)
}

func (ht *NoLinkHashTable) hSet(key string, value int) {
	bucket := Bucket{key, value}
	hashCode := GetHashCode(key, 1)
	index := hashCode % (ht.Capacity - 1)
	ht.addBucket(index, key, bucket)
}

func (ht *NoLinkHashTable) hGet(key string) Bucket {
	hashCode := GetHashCode(key, 1)
	index := hashCode % (ht.Capacity - 1)
	return ht.findBucket(index, key)
}

func (ht *NoLinkHashTable) findBucket(index int32, key string) Bucket {
	var empty = Bucket{}
	if ok := ht.Buckets[index]; ok != empty {
		if ht.Buckets[index].HKey == key {
			return ht.Buckets[index]
		} else {
			index++
		}
		for {
			if ok := ht.Buckets[index]; ok != empty {
				if ht.Buckets[index].HKey == key {
					return ht.Buckets[index]
				} else {
					index++
				}
			} else {
				return empty
			}

		}
	}
	return empty
}

func (ht *NoLinkHashTable) addBucket(index int32, key string, bucket Bucket) {
	if ht.Capacity == ht.Used {
		panic("Table is full")
	}
	var empty = Bucket{}
	if ok := ht.Buckets[index]; ok != empty {
		if ok.HKey == key {
			ht.Buckets[index] = bucket
			return
		}
		//已经存在
		index++
		for {
			//fmt.Println(index)
			if index >= ht.Capacity {
				index = 0
			}
			if ok := ht.Buckets[index]; ok != empty {
				if ok.HKey == key {
					ht.Buckets[index] = bucket
					return
				}
				index++
			} else {
				ht.Buckets[index] = bucket
				ht.Used++
				break
			}
		}
	} else {
		ht.Buckets[index] = bucket
		ht.Used++
	}

}
func GetHashCode(key string, hashType int) int32 {
	var Result []byte
	switch hashType {
	case 1:
		Md5Inst := md5.New()
		Md5Inst.Write([]byte(key))
		Result = Md5Inst.Sum([]byte(""))
	case 2:
		Sha1Inst := sha1.New()
		Sha1Inst.Write([]byte(key))
		Result = Sha1Inst.Sum([]byte(""))
	}
	bin_buf := bytes.NewBuffer(Result)
	var x int32
	binary.Read(bin_buf, binary.BigEndian, &x)
	return x
}

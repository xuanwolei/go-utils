package goutils

import(
	"github.com/garyburd/redigo/redis"
)

type BloomFilter struct{
	hashFunc []Func
	size uint64
	bitmap redis.Conn
	bitmapKey string
}

type Func func(key string) (uint64)

func NewBloomFilter(bitmap redis.Conn, size uint64, bitmapKey string) * BloomFilter{
	return &BloomFilter{
		size : size,
		bitmap : bitmap,
		hashFunc : []Func{RSHash,PJWHash,SDBMHash,},
		bitmapKey : bitmapKey,
	}
}

func (this *BloomFilter) Add(str string) (bool, error){
	bits := this.hash(str)
	for _,offset := range bits{
		this.bitmap.Do("SETBIT", this.bitmapKey, offset, 1)
	}

	return true, nil
}

func (this *BloomFilter) Exists(str string)(int64, error){
	var exists int64 = 1;
	bits := this.hash(str)
	for _,offset := range bits{	
		value, err := this.bitmap.Do("GETBIT", this.bitmapKey, offset)
		if err != nil {
			return 0, err
		}
		exists = exists & value.(int64)
	}

	return exists, nil
}

func (this *BloomFilter) hash(str string) ([]uint64){
	var bits []uint64
	for index := 0; index < len(this.hashFunc); index++ {	
		bits = append(bits,this.hashFunc[index](str) % this.size)
	}

	return bits
}




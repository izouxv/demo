package cache


import (
	"sync"
	log "github.com/cihub/seelog"
)

//供外部函数调用
type Cache struct {
	*cache
}

//缓存
type cache struct {
	items		map[int]map[int]int
	mu			sync.RWMutex
	onEvicted	func(int, map[int]int) //callback
	size		int32
}


//强制添加项
func (c *cache) Set(k, v, f int) {
	c.mu.Lock()
	c.set(k, v, f)
	c.mu.Unlock()
}

//添加不存在的项。否则返回错误
func (c *cache) Add(k, v, f int32) error {
	c.mu.Lock()
	//check exit
	_, found := c.get(int(k))
	if found {
		return log.Errorf("Item:%s has already exit", k)
	}
	//set
	c.set(int(k), int(v), int(f))
	c.mu.Unlock()
	return nil
}

//存入项
func (c *cache) set(k, v, f int) {
	vs := make(map[int]int)
	vs[v] = f
	c.items[k] = vs
	c.size++
}

//返回存在的项与true，否则返回false
func (c *cache) Get(k int) (map[int]int, bool) {
	c.mu.Lock()
	item, found := c.items[k]
	if !found {
		c.mu.Unlock()
		return nil, false
	}
	c.mu.Unlock()
	return item, true
}

//获取项
func (c *cache) get(k int) (map[int]int, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	return item, true
}

//更新项需存在，否则返回错误
func (c *cache) Replace(k, v, f int) error {
	c.mu.Lock()
	//check exit
	_, found := c.get(k)
	if !found {
		c.mu.Unlock()
		return log.Errorf("Item:%d absence exit", k)
	}
	c.set(k, v, f)
	c.mu.Unlock()
	return nil
}

//删除项，若有回调则返回
func (c *cache) delete(k int) (map[int]int, bool) {
	if v, found := c.items[k]; found {
		delete(c.items, k)
		c.size--
		if c.onEvicted != nil {
			return v, true
		}
	}
	return nil, false
}

//删除该项，若不存在则结束
func (c *cache) Delete(k int) {
	c.mu.Lock()
	v, evicted := c.delete(k)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(k, v)
	}

}

//返回缓存的所有项
func (c *cache) Item() map[int]map[int]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.items
}

// 返回缓存中项的数量
func (c *cache) Size() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.size

}

// 清空缓存
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = make(map[int]map[int]int)
	c.size = 0
	c.mu.Unlock()

}

// 创建一个缓存空间
func New() *Cache {
	c := &cache{
		items:	make(map[int]map[int]int),
		onEvicted:callBack,
	}
	//initContext Cache
	C := &Cache{c}
	return C
}

//实例化缓存
var c = New()
func GetCache() *Cache {
	return c
}

//定义回调函数
func callBack(k int, v map[int]int)  {
	log.Info("cache delete:",k, v)
}

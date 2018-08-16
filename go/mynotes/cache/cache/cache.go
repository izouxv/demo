package cache


import (
	"sync"
	"time"
	log "github.com/cihub/seelog"
	"runtime"
)

const (
	// 用于过期时间函数
	NoExpiration time.Duration = -1
	DefaultExpiration time.Duration = 0
)

//缓存对象
type Item struct {
	Object     interface{}
	Expiration int64
}

//对象过期返回true
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

//供外部函数调用
type Cache struct {
	*cache
}

//缓存
type cache struct {
	defaultExpiration time.Duration
	items		map[string]Item
	mu			sync.RWMutex
	onEvicted	func(string, interface{}) //callback
	janitor		*janitor
	size		int32
}


//强制添加项
func (c *cache) Set(k string, x interface{}, d time.Duration) {
	c.mu.Lock()
	c.set(k, x, d)
	c.mu.Unlock()
}

//添加不存在的项。否则返回错误
func (c *cache) Add(k string, x interface{}, d time.Duration) error {
	c.mu.Lock()
	//check exit
	_, found := c.get(k)
	if found {
		return log.Errorf("Item:%s has already exit", k)
	}
	//set
	c.set(k, x, d)
	c.mu.Unlock()
	return nil
}

//存入项
func (c *cache) set(k string, x interface{}, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}

	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item{
		Object:     x,
		Expiration: e,
	}
	c.size++
}

//返回存在的项与true，否则返回false
func (c *cache) Get(k string) (interface{}, bool) {
	c.mu.Lock()
	item, found := c.items[k]
	if !found || item.Expired()  {
		c.mu.Unlock()
		return nil, false
	}
	c.mu.Unlock()
	return item.Object, true
}

//获取项
func (c *cache) get(k string) (interface{}, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	//检查过期
	if item.Expiration > 0 && item.Expiration < time.Now().UnixNano() {
		return nil, false
	}
	return item.Object, true
}

//更新项需存在，否则返回错误
func (c *cache) Replace(k string, x interface{}, d time.Duration) error {
	c.mu.Lock()
	//check exit
	_, found := c.get(k)
	if !found {
		c.mu.Unlock()
		return log.Errorf("Item:%d absence exit", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
	return nil
}

//该项存在并且未过期时，为该项的int类型的值做加n，否则返回错误
func (c *cache) Increment(k string, n int) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return log.Errorf("Item:%v not found or expired",k)
	}
	switch v.Object.(type) {
	case int:
		v.Object = v.Object.(int) + int(n)
	default:
		c.mu.Unlock()
		return log.Errorf("not support value type")
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

//删除项，若有回调则返回
func (c *cache) delete(k string) (interface{}, bool) {
	if v, found := c.items[k]; found {
		delete(c.items, k)
		c.size--
		if c.onEvicted != nil {
			return v.Object, true
		}
	}
	return nil, false
}

//删除该项，若不存在则结束
func (c *cache) Delete(k string) {
	c.mu.Lock()
	v, evicted := c.delete(k)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(k, v)
	}

}

type kv struct {
	key   string
	value interface{}
}

//清除缓存中的过期项
func (c *cache) DeleteExpired() {
	var evictedItems []kv
	timeNow := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && v.Expiration < timeNow {
			v, evicted := c.delete(k)
			if evicted {
				evictedItems = append(evictedItems, kv{k, v})
			}
		}
	}
	c.mu.Unlock()
	for _, v := range evictedItems {
		c.onEvicted(v.key, v.value)
	}
}

//返回缓存的所有项
func (c *cache) Item() map[string]Item {
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

//设置回调函数，项从缓存中移除时触发，可以为nil
func (c *cache) OnEvicted(f func(string, interface{})) {
	c.mu.Lock()
	c.onEvicted = f
	c.mu.Unlock()
}

// 清空缓存
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = map[string]Item{}
	c.size = 0
	c.mu.Unlock()

}

//调用gc
type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	j.stop = make(chan bool)
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			//delete expired
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
	}
	c.janitor = j
	go j.Run(c)

}

func newCache(de time.Duration, m map[string]Item) *cache {
	if de == 0 {
		de = -1
	}
	c := &cache{
		defaultExpiration: de,
		items:             m,
		//size:0,
	}
	return c
}

//创建缓存、回调函数并启动清除函数
func newCacheWithJanitor(defaultExpiration time.Duration, cleanupInterval time.Duration, items map[string]Item) *Cache {
	c := newCache(defaultExpiration, items)
	//init Cache
	C := &Cache{c}
	if cleanupInterval > 0 {
		runJanitor(c, cleanupInterval)
		runtime.SetFinalizer(C, stopJanitor)
	}
	C.OnEvicted(callBack)
	return C
}

// 创建一个缓存空间
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	return newCacheWithJanitor(defaultExpiration, cleanupInterval, items)
}

//实例化缓存
var c = New(NoExpiration, time.Minute)
func GetCache() *Cache {
	return c
}

//定义回调函数
func callBack(k string, v interface{})  {
	log.Infof("cache delete:%s, %s",k, v)
}

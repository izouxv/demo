package test

import (
	"runtime"
	"sync"
	"testing"
	"time"
	"github.com/panjf2000/ants"
)

var n = 100000

func TestAntsPoolWithFunc(t *testing.T) {
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(AntsSize, func(i interface{}) error {
		demoPoolFunc(i)
		wg.Done()
		return nil
	})
	defer p.Release()

	for i := 0; i < n; i++ {
		wg.Add(1)
		p.Serve(Param)
	}
	wg.Wait()
	t.Logf("pool with func, running workers number:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d", mem.TotalAlloc/GiB)
}

func TestAntsPool(t *testing.T) {
	defer ants.Release()
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		ants.Submit(func() error {
			demoFunc()
			wg.Done()
			return nil
		})
	}
	wg.Wait()

	t.Logf("pool, capacity:%d", ants.Cap())
	t.Logf("pool, running workers number:%d", ants.Running())
	t.Logf("pool, free workers number:%d", ants.Free())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d MB", mem.TotalAlloc/MiB)
}

func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}

	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d MB", mem.TotalAlloc/MiB)
}

func TestCodeCov(t *testing.T) {
	_, err := ants.NewTimingPool(-1, -1)
	t.Log(err)
	_, err = ants.NewTimingPool(1, -1)
	t.Log(err)
	_, err = ants.NewTimingPoolWithFunc(-1, -1, demoPoolFunc)
	t.Log(err)
	_, err = ants.NewTimingPoolWithFunc(1, -1, demoPoolFunc)
	t.Log(err)

	p0, _ := ants.NewPool(AntsSize)
	defer p0.Submit(demoFunc)
	defer p0.Release()
	for i := 0; i < n; i++ {
		p0.Submit(demoFunc)
	}
	t.Logf("pool, capacity:%d", p0.Cap())
	t.Logf("pool, running workers number:%d", p0.Running())
	t.Logf("pool, free workers number:%d", p0.Free())
	p0.ReSize(AntsSize)
	p0.ReSize(AntsSize / 2)
	t.Logf("pool, after resize, capacity:%d, running:%d", p0.Cap(), p0.Running())

	p, _ := ants.NewPoolWithFunc(TestSize, demoPoolFunc)
	defer p.Serve(Param)
	defer p.Release()
	for i := 0; i < n; i++ {
		p.Serve(Param)
	}
	time.Sleep(ants.DefaultCleanIntervalTime * time.Second)
	t.Logf("pool with func, capacity:%d", p.Cap())
	t.Logf("pool with func, running workers number:%d", p.Running())
	t.Logf("pool with func, free workers number:%d", p.Free())
	p.ReSize(TestSize)
	p.ReSize(AntsSize)
	t.Logf("pool with func, after resize, capacity:%d, running:%d", p.Cap(), p.Running())
}

// func TestNoPool(t *testing.T) {
// 	var wg sync.WaitGroup
// 	for i := 0; i < n; i++ {
// 		wg.Add(1)
// 		go func() {
// 			demoPoolFunc(n)
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	mem := runtime.MemStats{}
// 	runtime.ReadMemStats(&mem)
// 	t.Logf("memory usage:%d", mem.TotalAlloc/GiB)
// }

//func TestCustomPool(t *testing.T) {
//	p, _ := ants.NewPool(30000)
//	var wg sync.WaitGroup
//	for i := 0; i < n; i++ {
//		wg.Add(1)
//		p.Submit(func() {
//			demoFunc()
//			//demoFunc()
//			wg.Done()
//		})
//	}
//	wg.Wait()
//
//	//t.Logf("pool capacity:%d", p.Cap())
//	//t.Logf("free workers number:%d", p.Free())
//
//	t.Logf("running workers number:%d", p.Running())
//	mem := runtime.MemStats{}
//	runtime.ReadMemStats(&mem)
//	t.Logf("memory usage:%d", mem.TotalAlloc/1024)
//}

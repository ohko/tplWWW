package controller

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/ohko/hst"
)

// statusController ...
type statusController struct {
	controller
}

func (o *statusController) Index(ctx *hst.Context) {
	ctx.Data(200, getSystemStatus())
}

var (
	sysStatusStartTime = time.Now()
	sysStatus          struct {
		Uptime       string
		NumGoroutine int

		// General statistics.
		MemAllocated string // bytes allocated and still in use
		MemTotal     string // bytes allocated (even if freed)
		MemSys       string // bytes obtained from system (sum of XxxSys below)
		Lookups      uint64 // number of pointer lookups
		MemMallocs   uint64 // number of mallocs
		MemFrees     uint64 // number of frees

		// Main allocation heap statistics.
		HeapAlloc    string // bytes allocated and still in use
		HeapSys      string // bytes obtained from system
		HeapIdle     string // bytes in idle spans
		HeapInuse    string // bytes in non-idle span
		HeapReleased string // bytes released to the OS
		HeapObjects  uint64 // total number of allocated objects

		// Low-level fixed-size structure allocator statistics.
		//	Inuse is bytes used now.
		//	Sys is bytes obtained from system.
		StackInuse  string // bootstrap stacks
		StackSys    string
		MSpanInuse  string // mspan structures
		MSpanSys    string
		MCacheInuse string // mcache structures
		MCacheSys   string
		BuckHashSys string // profiling bucket hash table
		GCSys       string // GC metadata
		OtherSys    string // other system allocations

		// Garbage collector statistics.
		NextGC       string // next run in HeapAlloc time (bytes)
		LastGC       string // last run in absolute time (ns)
		PauseTotalNs string
		PauseNs      string // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
		NumGC        uint32
	}
)

func getSystemStatus() string {

	logn := func(n, b float64) float64 {
		return math.Log(n) / math.Log(b)
	}
	humanateBytes := func(s uint64, base float64, sizes []string) string {
		if s < 10 {
			return fmt.Sprintf("%d B", s)
		}
		e := math.Floor(logn(float64(s), base))
		suffix := sizes[int(e)]
		val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
		f := "%.0f %s"
		if val < 10 {
			f = "%.1f %s"
		}

		return fmt.Sprintf(f, val, suffix)
	}
	// https://github.com/dustin/go-humanize/blob/master/bytes.go
	// IBytes(82854982) -> 79 MiB
	IBytes := func(s uint64) string {
		sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
		return humanateBytes(s, 1024, sizes)
	}
	// FileSize calculates the file size and generate user-friendly string.
	FileSize := func(s int64) string {
		return IBytes(uint64(s))
	}

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.Uptime = time.Since(sysStatusStartTime).String() // 服务运行时间
	sysStatus.NumGoroutine = runtime.NumGoroutine()            // 当前 Goroutines 数量

	sysStatus.MemAllocated = FileSize(int64(m.Alloc))  // 当前内存使用量
	sysStatus.MemTotal = FileSize(int64(m.TotalAlloc)) // 所有被分配的内存
	sysStatus.MemSys = FileSize(int64(m.Sys))          // 内存占用量
	sysStatus.Lookups = m.Lookups                      // 指针查找次数
	sysStatus.MemMallocs = m.Mallocs                   // 内存分配次数
	sysStatus.MemFrees = m.Frees                       // 内存释放次数

	sysStatus.HeapAlloc = FileSize(int64(m.HeapAlloc))       // 当前 Heap 内存使用量
	sysStatus.HeapSys = FileSize(int64(m.HeapSys))           // Heap 内存占用量
	sysStatus.HeapIdle = FileSize(int64(m.HeapIdle))         // Heap 内存空闲量
	sysStatus.HeapInuse = FileSize(int64(m.HeapInuse))       // 正在使用的 Heap 内存
	sysStatus.HeapReleased = FileSize(int64(m.HeapReleased)) // 被释放的 Heap 内存
	sysStatus.HeapObjects = m.HeapObjects                    // Heap 对象数量

	sysStatus.StackInuse = FileSize(int64(m.StackInuse))   // 启动 Stack 使用量
	sysStatus.StackSys = FileSize(int64(m.StackSys))       // 被分配的 Stack 内存
	sysStatus.MSpanInuse = FileSize(int64(m.MSpanInuse))   // MSpan 结构内存使用量
	sysStatus.MSpanSys = FileSize(int64(m.MSpanSys))       // 被分配的 MSpan 结构内存
	sysStatus.MCacheInuse = FileSize(int64(m.MCacheInuse)) // MCache 结构内存使用量
	sysStatus.MCacheSys = FileSize(int64(m.MCacheSys))     // 被分配的 MCache 结构内存
	sysStatus.BuckHashSys = FileSize(int64(m.BuckHashSys)) // 被分配的剖析哈希表内存
	sysStatus.GCSys = FileSize(int64(m.GCSys))             // 被分配的 GC 元数据内存
	sysStatus.OtherSys = FileSize(int64(m.OtherSys))       // 其它被分配的系统内存

	sysStatus.NextGC = FileSize(int64(m.NextGC))                                                           // 下次 GC 内存回收量
	sysStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000) // 距离上次 GC 时间
	if m.LastGC <= 0 {
		sysStatus.LastGC = "0.0s"
	}
	sysStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)          // GC 暂停时间总量
	sysStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000) // 上次 GC 暂停时间
	sysStatus.NumGC = m.NumGC                                                                      // GC 执行次数

	htm := `<html lang="zh-CN"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"></head>`
	htm += `<body><table border=1>`
	htm += fmt.Sprintf(`<tr><td>服务运行时间</td><td>%v</td></tr>`, sysStatus.Uptime)
	htm += fmt.Sprintf(`<tr><td>当前 Goroutines 数量</td><td>%v</td></tr>`, sysStatus.NumGoroutine)
	htm += `<tr><th colspan=2>内存</th></tr>`
	htm += fmt.Sprintf(`<tr><td>当前内存使用量</td><td>%v</td></tr>`, sysStatus.MemAllocated)
	htm += fmt.Sprintf(`<tr><td>所有被分配的内存</td><td>%v</td></tr>`, sysStatus.MemTotal)
	htm += fmt.Sprintf(`<tr><td>内存占用量</td><td>%v</td></tr>`, sysStatus.MemSys)
	htm += fmt.Sprintf(`<tr><td>指针查找次数</td><td>%v</td></tr>`, sysStatus.Lookups)
	htm += fmt.Sprintf(`<tr><td>内存分配次数</td><td>%v</td></tr>`, sysStatus.MemMallocs)
	htm += fmt.Sprintf(`<tr><td>内存释放次数</td><td>%v</td></tr>`, sysStatus.MemFrees)
	htm += `<tr><th colspan=2>Heap</th></tr>`
	htm += fmt.Sprintf(`<tr><td>当前 Heap 内存使用量</td><td>%v</td></tr>`, sysStatus.HeapAlloc)
	htm += fmt.Sprintf(`<tr><td>Heap 内存占用量</td><td>%v</td></tr>`, sysStatus.HeapSys)
	htm += fmt.Sprintf(`<tr><td>Heap 内存空闲量</td><td>%v</td></tr>`, sysStatus.HeapIdle)
	htm += fmt.Sprintf(`<tr><td>正在使用的 Heap 内存</td><td>%v</td></tr>`, sysStatus.HeapInuse)
	htm += fmt.Sprintf(`<tr><td>被释放的 Heap 内存</td><td>%v</td></tr>`, sysStatus.HeapReleased)
	htm += fmt.Sprintf(`<tr><td>Heap 对象数量</td><td>%v</td></tr>`, sysStatus.HeapObjects)
	htm += `<tr><th colspan=2>Stack</th></tr>`
	htm += fmt.Sprintf(`<tr><td>启动 Stack 使用量</td><td>%v</td></tr>`, sysStatus.StackInuse)
	htm += fmt.Sprintf(`<tr><td>被分配的 Stack 内存</td><td>%v</td></tr>`, sysStatus.StackSys)
	htm += fmt.Sprintf(`<tr><td>MSpan 结构内存使用量</td><td>%v</td></tr>`, sysStatus.MSpanInuse)
	htm += fmt.Sprintf(`<tr><td>被分配的 MSpan 结构内存</td><td>%v</td></tr>`, sysStatus.MSpanSys)
	htm += fmt.Sprintf(`<tr><td>MCache 结构内存使用量</td><td>%v</td></tr>`, sysStatus.MCacheInuse)
	htm += fmt.Sprintf(`<tr><td>被分配的 MCache 结构内存</td><td>%v</td></tr>`, sysStatus.MCacheSys)
	htm += fmt.Sprintf(`<tr><td>被分配的剖析哈希表内存</td><td>%v</td></tr>`, sysStatus.BuckHashSys)
	htm += fmt.Sprintf(`<tr><td>被分配的 GC 元数据内存</td><td>%v</td></tr>`, sysStatus.GCSys)
	htm += fmt.Sprintf(`<tr><td>其它被分配的系统内存</td><td>%v</td></tr>`, sysStatus.OtherSys)
	htm += `<tr><th colspan=2>GC</th></tr>`
	htm += fmt.Sprintf(`<tr><td>下次 GC 内存回收量</td><td>%v</td></tr>`, sysStatus.NextGC)
	htm += fmt.Sprintf(`<tr><td>距离上次 GC 时间</td><td>%v</td></tr>`, sysStatus.LastGC)
	htm += fmt.Sprintf(`<tr><td>GC 暂停时间总量</td><td>%v</td></tr>`, sysStatus.PauseTotalNs)
	htm += fmt.Sprintf(`<tr><td>上次 GC 暂停时间</td><td>%v</td></tr>`, sysStatus.PauseNs)
	htm += fmt.Sprintf(`<tr><td>GC 执行次数</td><td>%v</td></tr>`, sysStatus.NumGC)
	htm += `</table></body></html>`

	return htm
}

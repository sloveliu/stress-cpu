package handlers

import (
	"log"
	"net/http"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
)

// StressResponse API 回應格式
type StressResponse struct {
	Status    string `json:"status"`
	Stressing bool   `json:"stressing"`
	Workers   int    `json:"workers,omitempty"` // Workers 為 0，就不包含此欄位
}

// StressHandler 封裝 cpu stress 狀態、方法
type StressHandler struct {
	mu          sync.Mutex
	isStressing bool
	stopChan    chan struct{}
}

// NewStressHandler 建立 StressHandler 實例
func NewStressHandler() *StressHandler {
	return &StressHandler{}
}

func (h *StressHandler) cpuWorker(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			return
		default:
			for i := 0; i < 10000; i++ { // 迴圈次數可改變負載強度
				_ = float64(i) * float64(i) / 3.1415926535
			}
			runtime.Gosched() // 讓出 cpu 給其他 Goroutine
		}
	}
}

// Start 開始 cpu stress
func (h *StressHandler) Start(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.isStressing {
		c.JSON(http.StatusOK, StressResponse{Status: "cpu stress already running", Stressing: true})
		return
	}

	h.isStressing = true
	h.stopChan = make(chan struct{}) // 建立新的 stop channel

	numCPU := runtime.NumCPU()
	var wg sync.WaitGroup

	log.Printf("startup cpu stress ,use %d cpu worker...\n", numCPU)
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go h.cpuWorker(h.stopChan, &wg)
	}

	c.JSON(http.StatusOK, StressResponse{Status: "cpu stress startup", Stressing: true, Workers: numCPU})
}

// Stop 停止 cpu stress
func (h *StressHandler) Stop(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.isStressing {
		c.JSON(http.StatusOK, StressResponse{Status: "cpu stress not running", Stressing: false})
		return
	}

	log.Println("stop cpu stress...")
	if h.stopChan != nil {
		close(h.stopChan) // 發送停止訊號給所有 worker
	}
	h.isStressing = false

	c.JSON(http.StatusOK, StressResponse{Status: "cpu stress stop", Stressing: false})
}

// Status 取得 cpu stress 狀態
func (h *StressHandler) Status(c *gin.Context) {
	h.mu.Lock()
	currentStatus := h.isStressing
	h.mu.Unlock()

	statusMsg := "cpu stress not running"
	if currentStatus {
		statusMsg = "cpu stress running"
	}
	c.JSON(http.StatusOK, StressResponse{Status: statusMsg, Stressing: currentStatus})
}

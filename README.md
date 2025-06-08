# 說明

## 使用者故事

網站主備機制下，Oracle VPS 作為備援使用，以往是使用保活腳本讓 CPU 使用率保持 20% 上下，維持備援機不被關閉。但因腳本需手動開關，主備切換時，得進主機開關保活腳本較麻煩，所以實作當主備切換實能同時透過此 API 開關保活。

## 注意事項

- goroutine 數量會抓主機核心數，這適用在 VPS，不適用在 kubernetes。改用 GOMAXPROCS 話，會導致只跑 1 個 goroutine

- 在 docker 運行時，須指定 --cpus 參數以限制 container CPU 用量，否則會把主機炸了。

```bash
# 抓取 cpu 核心數 * 0.2 給 container 啟動用
cpu_limit=$(nproc | awk '{print $1 * 0.2}')

# 容器運行命令
docker run -itd --name $CONTAINER_NAME -p 8080:8080 --cpus=$cpu_limit $REGISTRY_URL/$BINARY_NAME:$VERSION
```

- 可在前面加上 nginx 處理憑證問題，或將憑證實作在程式碼中，避免 API Token 遭中間人竊取。

## 啟動壓測

`bash
curl -X GET http://localhost:8080/stress/start -H "X-API-KEY: your-secret-api-key"
`

## 查詢狀態

`bash
curl -X GET http://localhost:8080/stress/status -H "X-API-KEY: your-secret-api-key"
`

## 停止壓測

`bash
curl -X GET http://localhost:8080/stress/stop -H "X-API-KEY: your-secret-api-key"
`

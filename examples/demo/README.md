# 使用币安数据试算VolumeProfile

使用方法
```bash
API_KEY=xxx SECRET_KEY=yyy go run main.go
```

会自动从币安下载最近的 ETH-USDT 交易对的 15m K线进行计算，打印到控制台。
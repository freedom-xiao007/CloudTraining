# 模块二作业
***

## 作业要求
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200

## 作业提交
所以代码已简单写到mian.go函数中，可以运行测试进行快速请求验证

1. 请求已进行遍历和复制，可以在测试中看到
2. 读取环境变量，如果没有，默认为default，测试中进行了断言
3. 自定义了gin的日志输出，已包含客户端ip、HTTP返回码等（直接测试ip为空，使用浏览器等访问可以看到ip）
4. 对返回进行断言，返回的200

日志输出大致如下：

```text
# 测试输出
2022/01/09 18:46:34 Accept-Encoding [gzip]
2022/01/09 18:46:34 User-Agent [Go-http-client/1.1]
::1 - [Sun, 09 Jan 2022 18:46:34 CST] "GET /v1/hello HTTP/1.1 200 520.2µs "Go-http-client/1.1" "
[GIN] 2022/01/09 - 18:46:34 | 200 |       520.2µs |   

# 浏览器访问输出
2022/01/09 18:50:00 Upgrade-Insecure-Requests [1]
2022/01/09 18:50:00 User-Agent [Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36]
2022/01/09 18:50:00 Sec-Ch-Ua-Mobile [?0]
2022/01/09 18:50:00 Sec-Fetch-Site [none]
2022/01/09 18:50:00 Sec-Fetch-Mode [navigate]
2022/01/09 18:50:00 Cache-Control [max-age=0]
2022/01/09 18:50:00 Sec-Ch-Ua [" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"]
2022/01/09 18:50:00 Sec-Ch-Ua-Platform ["Windows"]
2022/01/09 18:50:00 Sec-Fetch-User [?1]
2022/01/09 18:50:00 Accept-Language [zh-CN,zh;q=0.9]
2022/01/09 18:50:00 Connection [keep-alive]
2022/01/09 18:50:00 Accept [text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]
2022/01/09 18:50:00 Sec-Fetch-Dest [document]
2022/01/09 18:50:00 Accept-Encoding [gzip, deflate, br]
127.0.0.1 - [Sun, 09 Jan 2022 18:50:00 CST] "GET /v1/hello HTTP/1.1 200 3.2357ms "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36" "    
[GIN] 2022/01/09 - 18:50:00 | 200 |      3.7686ms |       127.0.0.1 | GET      "/v1/hello"
```
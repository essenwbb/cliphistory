# Cliphistory
> Golang clipboard recorder. 


## Requirements
- go > 1.9
- GCC
    - For Windows: MinGW-w64 (Use recommended) or other GCC



## Installation:
```
go get github.com/go-vgo/robotgo
```
  It's that easy!

## How to run by source code

- ### windows
    ```
    go run main.go database.go
    ```

- ### linux
    ```
    go run *.go
    ```

## TODO list

- [x] History置顶实现
- [x] 记录数据至sqllite
- [ ] electronjs 编写界面
- [ ] 多平台支持及二进制可执行文件下载
- [ ] 热键冲突问题
- [ ] 先进后出 连续Ctrl+v 出栈  依赖于热键冲突问题
- [ ] 开启配置 翻译功能
- [ ] 呼出web页面     依赖于热键冲突问题
- [ ] 记录分词(该功能无关紧要)
- [ ] 历史记录展示
- [ ] websocket 实时更新历史记录器
- [ ] 持久化之后可以考虑对历史记录进行dashboard
- [ ] 启动时默认弹出界面
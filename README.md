# Cliphistory
> Golang clipboard recorder. 

## 使用方法

## TODO LIST
- ### 记录模块
    记录数据至sqllite

- ### web 界面
    - 浏览历史记录
    - 更新当前粘贴板  置顶某一条目
    - 记录分词(该功能无关紧要)
    
- ### other idea
    ```
    绑定热键  呼出web页面   面临热键冲突问题  
    启动时默认弹出界面
    直接输入地址获取界面
    持久化之后可以考虑对历史记录进行dashboard
    采用压栈  及  点歌置顶的方式实现
    置顶条目时  调用github.com/atotto/clipboard.WriteAll()
    ```
##  一个Go数据库事务管理库

其他语言：

### **[English](README.md)**

这是一个简单的Go数据库事务管理库。它的目的是创建非侵入式事务管理。 在Go的“sql”库中，有两个数据库处理程序“sql.DB”和
“sql.Tx”。 没有事务管理时你使用“sql.DB”访问数据库； 进行事务管理时，你使用“sql.Tx”。为了共享持久层代码，持久层需要同时
支持两者。 因此，在此库中创建了一个对数据库链接的封装。

要使用它，您需要将该库中的封装作为访问数据库的数据库链接。 除此之外，唯一需要做的就是将要用“EnableTx（）”来调用你想要支持
事务理的函数。

有关如何在实际项目中使用它的示例，请查看 ["servicetmpl1"](https://github.com/jfeng45/servicetmpl1).

### 下载程序

```
go get github.com/jfeng45/glogger
```

### 授权

[MIT](LICENSE.txt) 授权



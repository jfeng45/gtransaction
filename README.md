# A Go Database Trasnaction Managemnet Lib 

Other language: 
### **[中文](README.zh.md)**
 
This is a simple Go SQL database transaction lib. The purpose is to create a non-intrusive transaction
 management in Go. In Go’s “sql” lib, there are two database handlers sql.DB and sql.Tx. When there is no transaction,
  you use "sql.DB" to access database; when there is a transaction, you use "sql.Tx". In order to share the same 
  persistence code between transaction and non-transaction use case, the persistence layer needs to support both. A 
  wrapper is created in this lib for this purpose.  

To use it, you need to use the wrapper in this lib as the database connection to access database. Other than that, the
 only thing you need to do is to wrap the function that you want transaction into "EnableTx()".  

For examples on how to use it in real project, please take a look at ["servicetmpl1"](https://github.com/jfeng45/servicetmpl1).

### Download Code

```
go get github.com/jfeng45/gtransaction
```

### License

[MIT](LICENSE.txt) License



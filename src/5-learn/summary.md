# 第五天

数据在实验环境下使用的连接包是`github.com/lib/pq` 这是用Postgres作为数据库的驱动。官方文档在[这里](https://godoc.org/github.com/lib/pq)。    就使用来说还是很方便的具体可以参考官方给出的例子
```golang
import (
	"database/sql"
	_ "github.com/lib/pq"
}

func main() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	…
}
```

在第五天的数据库实践里，更多的是结合之前几天的内容进行综合的联系。第一个就是go里面的分包用法*大写公有，小写私有*。 自定义 Error的用法 *重写Error() string* 方法。其实官方提供的异常处理可以参考[这里](https://blog.golang.org/error-handling-and-go)

```golang
type Error struct {
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("has error -> %s", e.msg)
}
```

在[getConnectionInfo](./DBUtil/ConnectionInfo.go)里用到一个单例模式，这里需要用到数据库的链接信息，这个消息在工具初始化的时候才实例化的，使用的时候需要先使用`InitConnectionInfo`方法来初始化数据库连接信息。然后使用自己定义的结构体来存储数据库连接信息

```golang
type connectInfo struct {
	Uname string
	DbName string
	Pwd string
	Host string
}
```
然后是一个[核心工具类](./DBUtil/GetConn.go) 这里面封装了上面说到的工具`pq`，使用起来也是很方便的。定义SQL连接信息模板。这里使用的就是刚刚写到的链接细腻的结构体，这里使用到的是`text/template`这个工具类

```golang
	templatePoint, _ := template.New("db").
		Parse("user={{.Uname}} dbname={{.DbName}} password={{.Pwd}} host={{.Host}} sslmode=disable")
  var tpl bytes.Buffer
	if err := sqlConnectionTemplate.ExecuteTemplate(&tpl, "db", connectInfo); err != nil {
		fmt.Println(err.Error())
		return nil, &Error{"sql转换失败"}
	}

	sqlConnectionStr = tpl.String()
```
这里使用的是模板的`ExecuteTemplate`方法，这里面需要一个数据流，这里可以使用`bytes.Buffer` 来接收，然后用他的`地址`来存储流中的信息，最后在使用`String()`方法来输出结果。这样就可以接收输入输出流    
然后封装一个`inteface` 来实现aop功能，实现通用业务封装

```golang
type breakthroughPoint interface {
	execute(*sql.DB) (interface{}, error)
}

func UseConnection(point breakthroughPoint) (interface{}, error) {
	db, error := openConnection()
	if nil != error {
		return nil, &Error{error.Error()}
	}
	rest, err := point.execute(db)
	defer db.Close()
	return rest, err
}
```

然后每个类只要使用`UseConnection()`方法，就能够实现数据db获取，也可用于并发使用。   

接下来就是数据库的CRUD四种操作，优先Query吧，这里封装了一个方法，传入sql返回map结果，可以用于日常使用代码可以查看[这里](./DBUtil/Query.go) 核心是这里

```golang
func (query *queryExecute) rowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, _ := rows.Columns()
	result := make([]map[string]interface{}, 0)
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	for rows.Next() {
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, &Error{err.Error()}
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		result = append(result, m)
	}

	return result, nil
}
```

解析一下
* 查询的记过返回的是一个rows的指针，可以调用`Columns`方法获取返回的列头
* 定义两个slice，一个slice用户存储列的数据，一个slice来存储对应的指针
* 给指针数组赋值
* 循环返回结果，每次调用`Scan`方法，传入指针slice，使用`...`来实现 可变参数传入
* 将获取的数据转换到map之中
* 把每次的map通过`append`方法加入的返回结果集 即可

插入的代码和查询的类型，可以参看[这里](./DBUtil/Insert.go)  

修改和删除的代码可以用同一段可以参看[这里](./DBUtil/Update.go) 不过需要注意的就是 操作之后要调用 res的方法
```golang
func checkRes(updateHistory []interface{}) (interface{}, error) {
	for _, update := range updateHistory {
		val := reflect.ValueOf(update)
		val.MethodByName("RowsAffected").Call(nil)
	}

	return true, nil
}
```

这里是使用反射的方法调用方法来实现修改后的代码提交的。最后在[Main.go](./Main.go)里面进行CRUD的具体操作


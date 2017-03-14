1.reflect.Type 是一个接口

2.想拿到一个结构体的字段名，需要reflect.TypeOf(v).Field(1).Name

  想要给结构体字段赋值，需要reflect.ValueOf(v).Field(i).Set()

3.// 判断一个变量是否为结构体,两种方法

```
a := Input{}
v := reflect.ValueOf(a).Kind()
v := reflect.TypeOf(a).Kind()
```
4.根据a reflect.Value获取a的type

```
aa := Input{}
reflect.ValueOf(aa).Type()
```


3.objT reflect.Type, objV reflect.Value

```
    fieldT := objT.Field(i) // 这是一个StructField
    fieldV := objV.Field(i) // 这是一个Type
```
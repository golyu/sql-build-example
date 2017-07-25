#### sql-build-example

`sql-build-example`是[sql-build](https://github.com/golyu/sql-build)项目的一个示例,和beego的orm结合使用

##### 假设我们的业务需求里,int,int8,int16,in32,int64这样的数据,我们要过滤的是小于0的-1,而不是`sql-build`项目中默认整型都是过滤0的条件,过滤0的条件,我们会用uint...类型

这里就是对sql-build使用的一个小例子

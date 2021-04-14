type StructName struct {

}

结构体名字：
StructName 结构体的名字首字母大写 外可见
structName 

属性名：
Attr 大写能访问
attr

匿名嵌入

S大写结构体 A大写属性 =>嵌入
S A => s  小写的不能再外部访问 嵌入其的自然也不能访问
S a => s

S A => S 结构体名大写 包外可以创建结构体对象 属性也可以访问
S a => S 结构体名大写 包外可以创建结构体对象 属性不可以访问


s A => S 这种比较特殊 小写的匿名 结构体 放到大写的结构体里面 包外可以创建结构体对象
s a => S
s A => s
s a => s 

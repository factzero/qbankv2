package controller

/**
* 引入控制器
* 请把您使用包用 _ "gofly/app/home/XX"导入您编写的包 自动生成路由
* 不是使用则注释掉
* 路由规则：包路径“home/article” + 包中结构体“Cate”转小写+方法名(首字母转小写) 即：http://xx.com/home/article/cate/get_list
 */
import (
	_ "gozero/app/admin/managequestion"
	_ "gozero/app/admin/user"
)

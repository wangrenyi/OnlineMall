# Online Mall
* 1. 线上购物商城,实现安全登陆,商品管理,购物车,热点商品显示
* 2. 技术栈：golang,gin,gorm,jwt
* 3. 数据库：mysql,redis
* 4. 日志记录：logrus,日志分割: rotatelogs



## Error Log
* 1. Model属性值类型不能设置为bool,即数据库字段不能为bit,获取时会导致解析失败. 
*       issues: https://github.com/jinzhu/gorm/issues/1432
* 2. 定义相同request method的url时,前面的地址不能完全重合. 
*       panic: wildcard route '' conflicts with existing children in path
*       eg: Method GET, /goods/list和/goods/:goodsId. 改为: /goods/list,/goods/detail/:goodsId
* 3. win10环境 rotatelogs 配置日志分割,创建软连接失败
*       err: log_symlink: A required privilege is not held by the client

## Online Mall
* 1. 线上购物商城,实现安全登陆,商品管理,购物车,热点商品显示
* 2. 技术栈：golang,gin,gorm,jwt
* 3. 数据库：mysql,redis
* 4. 日志记录：logrus,日志分割: rotatelogs

## Gorm BaseDAO
*       使用golang反射机制,简化常用数据库操作. 详情可参考/onlinemall/repository/base_dao.go,如有更好
*       的实现方式,也可私信给我

## Error Log
* 1. Model属性值类型不能设置为bool,即数据库字段不能为bit,获取时会导致解析失败. 
*       issues: https://github.com/jinzhu/gorm/issues/1432
* 2. 定义相同request method的url时,前面的地址不能完全重合. 
*       panic: wildcard route '' conflicts with existing children in path
*       eg: Method GET, /goods/list和/goods/:goodsId. 改为: /goods/list,/goods/detail/:goodsId
* 3. win10环境 rotatelogs 配置日志分割,创建软连接失败
*       err: log_symlink: A required privilege is not held by the client

## shopping cart实现思路
* 1. 根据用户登录和未登录态两种状态
* 2. 未登录时
*       采取cookie方式保存,存在客户端,有助于缓解服务器压压力. 
* 3. 当用户登录账号后
*       从cookie取数据保存到redis里面 用户登录后，保存购物车数据到redis中.
*       使用redis保存数据在内存中,存取速度快. 高并发时易于拓展,可使用redis集群. 通常做法是把相对比较热的数据使用redis保存,
*       后续可以根据用户购物车记录或者购买记录等用户行为做商品推送等功能,可以使用redis的一些特性,比如取交并集等.
* 4. 保存方案
*       产品: 采用hash表存储,goodsId,price等等；
*       购物车: 采用集合存储,因为集合的唯一性,保证一个ordersId对应多个商品.
*       购物记录: 使用有序集合.  


## 联系方式
*       email: 3168628033@qq.com
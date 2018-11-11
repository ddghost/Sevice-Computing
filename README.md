项目的github地址 
https://github.com/ddghost/Sevice-Computing/tree/cloudgo-io

# 1.概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

# 2.写程序过程中遇到的坑
要用到的库 github.com/codegangsta/negroni 已经更名为 github.com/codegangsta/negroni 参考老师的博客时候要注意修改

在测试静态文件服务时，第一次访问时没有问题，但是访问完第一次后，对静态文件进行修改（增加，修改内容，删除），第二次访问时发现文件并没有发现修改，开始以为没保存，后来发现原来要刷新下页面才能更新，可能与浏览器的缓存有关。


# 3.使用流程
在当前文件夹直接输入 go run main.go 即可运行

访问127.0.0.1/static/ 支持静态文件服务
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181111120538400.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)


![在这里插入图片描述](https://img-blog.csdnimg.cn/20181111121020535.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)


![在这里插入图片描述](https://img-blog.csdnimg.cn/20181111121033803.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)


![在这里插入图片描述](https://img-blog.csdnimg.cn/20181111121105287.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)


访问127.0.0.1，来到主页面，输入用户名和密码
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181111121212360.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)

登陆后显示了提交的表单信息
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181111121233243.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)

访问127.0.0.1/api/unknown,返回错误码 501
![在这里插入图片描述](https://img-blog.csdnimg.cn/201811111213109.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0REZ2hzb3Q=,size_16,color_FFFFFF,t_70)

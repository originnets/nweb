#娱乐使用
## 1. api 
- 	post	/api/v1.0/user/reg		注册

		post 提交:

		{
		"username" : "admin",
		"password" : "123456"
		}
		返回:
		{
		    "code": "0",
		    "meg": "成功"
		}

- 	post	/api/v1.0/user/login	登录

		post 提交:

		{
		"username" : "admin",
		"password" : "123456"
		}
		返回:
		{
		    "code": "0",
		    "meg": "成功"
		}

- 	get		/api/v1.0/user/logout	退出
	
		{
	    "code": "0",
	    "meg": "成功"
		}

- 	get		/api/v1.0/domain/list	获取当前用户所有域名配置
	
		返回
		{
	    "code": "0",
	    "data": {
	        "0": {
	            "id": 5,
	            "log_name": "test2",
	            "port": 80,
	            "root": "/www/web",
	            "server_name": "5.test.com",
	            "status": 1
	        },
	        "1": {
	            "id": 5,
	            "log_name": "test2",
	            "port": 80,
	            "root": "/www/web",
	            "server_name": "5.test.com",
	            "status": 1
	        },
	        "2": {
	            "id": 5,
	            "log_name": "test2",
	            "port": 80,
	            "root": "/www/web",
	            "server_name": "5.test.com",
	            "status": 1
	        },
	        "3": {
	            "id": 5,
	            "log_name": "test2",
	            "port": 80,
	            "root": "/www/web",
	            "server_name": "5.test.com",
	            "status": 1
	        }
	    },
	    "meg": "成功"
		}

- 	post	/api/v1.0/domain/add	添加域名(web配置)

		post提交:
		{
			"server_name" : "5.test.com",
			"port" : "80",
			"root" : "/www/web",
			"logname" : "test2",
			"status" : "1"
		}
		返回
		{
		    "code": "0",
		    "meg": "成功"
		}


- 	get		/api/v1.0/domain/delete/(id)	删除域名和相关配置
		
		返回
		{
		    "code": "0",
		    "meg": "成功"
		}
	
- 	get		/api/v1.0/domain/dis/(id)		关闭域名使用(web配置)
	
		返回
		{
		    "code": "0",
		    "meg": "成功"
		}
	
- 	get		/api/v1.0/domain/rec/(id)	恢复域名使用(web配置)

		返回
			{
			    "code": "0",
			    "meg": "成功"
			}

- 	post	/api/v1.0/domain/change/(id)	更改域名(web配置)
- 	
		post提交:
		{
		"server_name" : "2.test.com",
		"port" : "80",
		"root" : "/www/8080",
		"logname" : "www"
		}
		返回
		{
		    "code": "0",
		    "meg": "成功"
		}

### id:表示获取列表中的id值
	
	
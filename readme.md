1.单机版完成  
    1.1数据持久化
    1.2架构，爬取不同数据只需要提供不同得request结构体
        req := models.Request{
    		Url:        "https://movie.douban.com/subject/25827935/",
    		ParserFunc: parser.ParseMovieInfo,
    	}
    	engine.Run(req)
    1.3去重
    1.403，202解决
    1.5任务队列
2.并发版
    2.1 simpleScheduler
        1.一个输入通道in，接收request，一个输出通道out
        2.worker接收通道里的request，处理输入out
        3.out取出继续向in通道发送request
    2.2 queueScheduler
        1.使用request队列和worker队列实现
        2.一个worker队列里包含多个worker(request输入通道)
        3.一个worker准备好了，就发送到worker chan，接收到worker放入队列
          request chan接收到request也放入队列
        4.当request队列和worker队列都不为空就表示，可以取出一个worker去处理
         一个request，将这个request传入worker
    2.3 ip禁用的解决
    user_agent 伪装和轮换
    不同浏览器的不同版本都有不同的user_agent，是浏览器类型的详细信息，也是浏览器提交Http请求的重要头部信息。我们可以在每次请求的时候提供不同的user_agent，绕过网站检测客户端的反爬虫机制。比如说，可以把很多的user_agent放在一个列表中，每次随机选一个用于提交访问请求。有一个提供各种user_agent的网站：
    
    http://www.useragentstring.com/
    
    最近又看到一个专门提供伪装浏览器身份的开源库，名字取得很直白：
    
    fake-useragent（https://github.com/hellysmile/fake-useragent）
    
    使用代理IP和轮换
    检查ip的访问情况是网站的反爬机制最喜欢也最喜欢用的方式。这种时候就可以更换不同的ip地址来爬取内容。当然，你有很多有公网ip地址的主机或者vps是更好的选择，如果没有的话就可以考虑使用代理，让代理服务器去帮你获得网页内容，然后再转发回你的电脑。代理按透明度可以分为透明代理、匿名代理和高度匿名代理:
    
    透明代理：目标网站知道你使用了代理并且知道你的源IP地址，这种代理显然不符合我们这里使用代理的初衷
    匿名代理：匿名程度比较低，也就是网站知道你使用了代理，但是并不知道你的源IP地址
    高匿代理：这是最保险的方式，目标网站既不知道你使用的代理更不知道你的源IP 
    代理的获取方式可以去购买，当然也可以去自己爬取免费的，这里（http://www.xicidaili.com/nn/）有一个提供免费代理的网站，可以爬下来使用，但是免费的代理通常不够稳定。
    设置访问时间间隔
    很多网站的反爬虫机制都设置了访问间隔时间，一个IP如果短时间内超过了指定的次数就会进入“冷却CD”，所以除了轮换IP和user_agent 
    可以设置访问的时间间间隔长一点，比如没抓取一个页面休眠一个随机时间：
    
    import time，random
    time.sleep(random.random()*3)
    对于一个crawler来说，这是一个比较responsible的做法。 
    因为本来爬虫就可能会给对方网站造成访问的负载压力，所以这种防范既可以从一定程度上防止被封，还可以降低对方的访问压力。   
    
    
    
    2.3去重的封装
  

3.分布式
    使用redis去做任务队列和去重
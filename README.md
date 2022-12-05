# How To Use

- 导入数据库
- 配置nginx
- 配置启动脚本

# mysql创建数据库

```sql
CREATE DATABASE query DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE USER 'query'@'localhost' IDENTIFIED BY 'CoAx54QMh7';

GRANT ALL ON query.* TO 'query'@'localhost';
```

# nginx配置文件

```nginx
server {
  listen 8443 ssl;
  server_name nu1l.online;

  ssl on;
  ssl_certificate      /root/.acme.sh/nu1l.online_ecc/fullchain.cer;
  ssl_certificate_key  /root/.acme.sh/nu1l.online_ecc/nu1l.online.key;

  location / {
        proxy_pass http://localhost:8880;
  }
}
```

# query启动脚本

```shell
pid=$(netstat -antlp|grep 8880|awk '{print $7}'|awk -F"/" '{print $1}');

HOME='/root/go-web';

start(){
        cd $HOME
        nohup ./query &
        echo "start success"
}

stop() {
        kill -9 $pid
        echo "stop success"
}

case $1 in
        start)
                start
        ;;
        stop)
                stop
        ;;
        restart)
                $0 stop
                sleep 1s
                $0 start
        ;;
        *)
        echo "Usage: {start|stop|restart}"
        ;;
esac

exit 0
```
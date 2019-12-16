### 试用docker

```bash
graftcp  docker pull mysql:5.7
```

```bash
d599a449871e: Pull complete 
f287049d3170: Pull complete 
08947732a1b0: Pull complete 
96f3056887f2: Pull complete 
871f7f65f017: Pull complete 
1dd50c4b99cb: Pull complete 
5bcbdf508448: Pull complete 
02a97db830bd: Pull complete 
c09912a99bce: Pull complete 
08a981fc6a89: Pull complete 
818a84239152: Pull complete
Digest: sha256:5779c71a4730da36f013a23a437b5831198e68e634575f487d37a0639470e3a8
Status: Downloaded newer image for mysql:5.7
docker.io/library/mysql:5.7
```

试用mysql的docker,由于自己电脑跑着mysql，所以一开始将docker的3306端口映射到主机的时候出错，试用命令删除这个错误镜像

```bash
docker container prune -f
```

然后执行

```bash
sudo docker run -p 3307:3306 --name mysql2 -e MYSQL_ROOT_PASSWORD=root -d  mysql:5.7
```

查看效果

```bash
docker ps -a
```

```bash
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                 NAMES
15f57ba3e9a1        mysql:5.7           "docker-entrypoint.s…"   12 minutes ago      Up 12 minutes       3306/tcp, 33060/tcp   mysql2
```

#### 创建卷并挂载

将原来的stop然后remove

```bash
docker stop 15f57ba3e9a1
docker rm 15f57ba3e9a1
```

```bash
docker volume create mydb
docker run --name mysql2 -e MYSQL_ROOT_PASSWORD=root -v mydb:/var/lib/mysql -d mysql:5.7
docker run --name myclient --link mysql2:mysql -it mysql:5.7 bash
root@dc2e389ebde8:/# mysql -hmysql -P3306 -uroot -proot
```

```bash
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2
Server version: 5.7.28 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

后来启动的docker能够连接到前面的mysql server,而且使用的端口是3306，也就是docker内的端口

### docker compose

apt安装

```bash
apt-get install docker-compose
```

```bash
root@Linux2:/usr/local# docker-compose version
docker-compose version 1.17.1, build unknown
docker-py version: 2.5.1
CPython version: 2.7.15+
OpenSSL version: OpenSSL 1.1.1  11 Sep 2018

```

```bash
# mkdir comptest && cd comptest
comptest]# vi stack.yml
version: '3.1'
services:
  db:
    image: mysql:5.7
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example
  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080

# graftcp docker-compose -f stack.yml up
```

效果

```bash
adminer_1  | [Mon Dec 16 11:46:50 2019] PHP 7.4.0 Development Server (http://[::]:8080) started
db_1       | 2019-12-16 11:46:49+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.28-1debian9 started.
db_1       | 2019-12-16 11:46:49+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
db_1       | 2019-12-16 11:46:49+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.28-1debian9 started.
db_1       | 2019-12-16 11:46:49+00:00 [Note] [Entrypoint]: Initializing database files
db_1       | 2019-12-16T11:46:49.757414Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
db_1       | 2019-12-16T11:46:50.245695Z 0 [Warning] InnoDB: New log files created, LSN=45790
db_1       | 2019-12-16T11:46:50.369083Z 0 [Warning] InnoDB: Creating foreign key constraint system tables.
db_1       | 2019-12-16T11:46:50.446374Z 0 [Warning] No existing UUID has been found, so we assume that this is the first time that this server has been started. Generating a new UUID: bf444c3a-1ff9-11ea-80c4-0242ac140003.
db_1       | 2019-12-16T11:46:50.449529Z 0 [Warning] Gtid table is not ready to be used. Table 'mysql.gtid_executed' cannot be opened.
db_1       | 2019-12-16T11:46:51.472569Z 0 [Warning] CA certificate ca.pem is self signed.
db_1       | 2019-12-16T11:46:51.752986Z 1 [Warning] root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option.
db_1       | 2019-12-16 11:46:56+00:00 [Note] [Entrypoint]: Database files initialized
db_1       | 2019-12-16 11:46:56+00:00 [Note] [Entrypoint]: Starting temporary server
db_1       | 2019-12-16 11:46:56+00:00 [Note] [Entrypoint]: Waiting for server startup
db_1       | 2019-12-16T11:46:57.076249Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
db_1       | 2019-12-16T11:46:57.077636Z 0 [Note] mysqld (mysqld 5.7.28) starting as process 81 ...
db_1       | 2019-12-16T11:46:57.081167Z 0 [Note] InnoDB: PUNCH HOLE support available
db_1       | 2019-12-16T11:46:57.081185Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
db_1       | 2019-12-16T11:46:57.081190Z 0 [Note] InnoDB: Uses event mutexes
db_1       | 2019-12-16T11:46:57.081195Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
db_1       | 2019-12-16T11:46:57.081199Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.11
db_1       | 2019-12-16T11:46:57.081203Z 0 [Note] InnoDB: Using Linux native AIO
db_1       | 2019-12-16T11:46:57.081451Z 0 [Note] InnoDB: Number of pools: 1
db_1       | 2019-12-16T11:46:57.081552Z 0 [Note] InnoDB: Using CPU crc32 instructions
db_1       | 2019-12-16T11:46:57.082999Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
db_1       | 2019-12-16T11:46:57.093592Z 0 [Note] InnoDB: Completed initialization of buffer pool
db_1       | 2019-12-16T11:46:57.095574Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
db_1       | 2019-12-16T11:46:57.107258Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
db_1       | 2019-12-16T11:46:57.117531Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
db_1       | 2019-12-16T11:46:57.117622Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
db_1       | 2019-12-16T11:46:57.157458Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
db_1       | 2019-12-16T11:46:57.158285Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
db_1       | 2019-12-16T11:46:57.158296Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
db_1       | 2019-12-16T11:46:57.158772Z 0 [Note] InnoDB: 5.7.28 started; log sequence number 2628896
db_1       | 2019-12-16T11:46:57.158981Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
db_1       | 2019-12-16T11:46:57.159219Z 0 [Note] Plugin 'FEDERATED' is disabled.
db_1       | 2019-12-16T11:46:57.160522Z 0 [Note] InnoDB: Buffer pool(s) load completed at 191216 11:46:57
db_1       | 2019-12-16T11:46:57.165487Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
db_1       | 2019-12-16T11:46:57.165502Z 0 [Note] Skipping generation of SSL certificates as certificate files are present in data directory.
db_1       | 2019-12-16T11:46:57.166350Z 0 [Warning] CA certificate ca.pem is self signed.
db_1       | 2019-12-16T11:46:57.166384Z 0 [Note] Skipping generation of RSA key pair as key files are present in data directory.
db_1       | 2019-12-16T11:46:57.169134Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
db_1       | 2019-12-16T11:46:57.177937Z 0 [Note] Event Scheduler: Loaded 0 events
db_1       | 2019-12-16T11:46:57.178220Z 0 [Note] mysqld: ready for connections.
db_1       | Version: '5.7.28'  socket: '/var/run/mysqld/mysqld.sock'  port: 0  MySQL Community Server (GPL)
db_1       | 2019-12-16 11:46:57+00:00 [Note] [Entrypoint]: Temporary server started.
db_1       | Warning: Unable to load '/usr/share/zoneinfo/iso3166.tab' as time zone. Skipping it.
db_1       | Warning: Unable to load '/usr/share/zoneinfo/leap-seconds.list' as time zone. Skipping it.
db_1       | Warning: Unable to load '/usr/share/zoneinfo/zone.tab' as time zone. Skipping it.
db_1       | Warning: Unable to load '/usr/share/zoneinfo/zone1970.tab' as time zone. Skipping it.
db_1       | 
db_1       | 2019-12-16 11:47:00+00:00 [Note] [Entrypoint]: Stopping temporary server
db_1       | 2019-12-16T11:47:00.853694Z 0 [Note] Giving 0 client threads a chance to die gracefully
db_1       | 2019-12-16T11:47:00.853757Z 0 [Note] Shutting down slave threads
db_1       | 2019-12-16T11:47:00.853769Z 0 [Note] Forcefully disconnecting 0 remaining clients
db_1       | 2019-12-16T11:47:00.853782Z 0 [Note] Event Scheduler: Purging the queue. 0 events
db_1       | 2019-12-16T11:47:00.853876Z 0 [Note] Binlog end
db_1       | 2019-12-16T11:47:00.855461Z 0 [Note] Shutting down plugin 'ngram'
db_1       | 2019-12-16T11:47:00.855506Z 0 [Note] Shutting down plugin 'partition'
db_1       | 2019-12-16T11:47:00.855521Z 0 [Note] Shutting down plugin 'BLACKHOLE'
db_1       | 2019-12-16T11:47:00.855535Z 0 [Note] Shutting down plugin 'ARCHIVE'
db_1       | 2019-12-16T11:47:00.855545Z 0 [Note] Shutting down plugin 'PERFORMANCE_SCHEMA'
db_1       | 2019-12-16T11:47:00.855626Z 0 [Note] Shutting down plugin 'MRG_MYISAM'
db_1       | 2019-12-16T11:47:00.855651Z 0 [Note] Shutting down plugin 'MyISAM'
db_1       | 2019-12-16T11:47:00.855680Z 0 [Note] Shutting down plugin 'INNODB_SYS_VIRTUAL'
db_1       | 2019-12-16T11:47:00.855693Z 0 [Note] Shutting down plugin 'INNODB_SYS_DATAFILES'
db_1       | 2019-12-16T11:47:00.855703Z 0 [Note] Shutting down plugin 'INNODB_SYS_TABLESPACES'
db_1       | 2019-12-16T11:47:00.855713Z 0 [Note] Shutting down plugin 'INNODB_SYS_FOREIGN_COLS'
db_1       | 2019-12-16T11:47:00.855723Z 0 [Note] Shutting down plugin 'INNODB_SYS_FOREIGN'
db_1       | 2019-12-16T11:47:00.855733Z 0 [Note] Shutting down plugin 'INNODB_SYS_FIELDS'
db_1       | 2019-12-16T11:47:00.855743Z 0 [Note] Shutting down plugin 'INNODB_SYS_COLUMNS'
db_1       | 2019-12-16T11:47:00.855754Z 0 [Note] Shutting down plugin 'INNODB_SYS_INDEXES'
db_1       | 2019-12-16T11:47:00.855764Z 0 [Note] Shutting down plugin 'INNODB_SYS_TABLESTATS'
db_1       | 2019-12-16T11:47:00.855774Z 0 [Note] Shutting down plugin 'INNODB_SYS_TABLES'
db_1       | 2019-12-16T11:47:00.855784Z 0 [Note] Shutting down plugin 'INNODB_FT_INDEX_TABLE'
db_1       | 2019-12-16T11:47:00.855791Z 0 [Note] Shutting down plugin 'INNODB_FT_INDEX_CACHE'
db_1       | 2019-12-16T11:47:00.855798Z 0 [Note] Shutting down plugin 'INNODB_FT_CONFIG'
db_1       | 2019-12-16T11:47:00.855805Z 0 [Note] Shutting down plugin 'INNODB_FT_BEING_DELETED'
db_1       | 2019-12-16T11:47:00.855812Z 0 [Note] Shutting down plugin 'INNODB_FT_DELETED'
db_1       | 2019-12-16T11:47:00.855819Z 0 [Note] Shutting down plugin 'INNODB_FT_DEFAULT_STOPWORD'
db_1       | 2019-12-16T11:47:00.855826Z 0 [Note] Shutting down plugin 'INNODB_METRICS'
db_1       | 2019-12-16T11:47:00.855833Z 0 [Note] Shutting down plugin 'INNODB_TEMP_TABLE_INFO'
db_1       | 2019-12-16T11:47:00.855841Z 0 [Note] Shutting down plugin 'INNODB_BUFFER_POOL_STATS'
db_1       | 2019-12-16T11:47:00.855848Z 0 [Note] Shutting down plugin 'INNODB_BUFFER_PAGE_LRU'
db_1       | 2019-12-16T11:47:00.855855Z 0 [Note] Shutting down plugin 'INNODB_BUFFER_PAGE'
db_1       | 2019-12-16T11:47:00.855862Z 0 [Note] Shutting down plugin 'INNODB_CMP_PER_INDEX_RESET'
db_1       | 2019-12-16T11:47:00.855869Z 0 [Note] Shutting down plugin 'INNODB_CMP_PER_INDEX'
db_1       | 2019-12-16T11:47:00.855879Z 0 [Note] Shutting down plugin 'INNODB_CMPMEM_RESET'
db_1       | 2019-12-16T11:47:00.855890Z 0 [Note] Shutting down plugin 'INNODB_CMPMEM'
db_1       | 2019-12-16T11:47:00.855901Z 0 [Note] Shutting down plugin 'INNODB_CMP_RESET'
db_1       | 2019-12-16T11:47:00.855912Z 0 [Note] Shutting down plugin 'INNODB_CMP'
db_1       | 2019-12-16T11:47:00.855922Z 0 [Note] Shutting down plugin 'INNODB_LOCK_WAITS'
db_1       | 2019-12-16T11:47:00.855932Z 0 [Note] Shutting down plugin 'INNODB_LOCKS'
db_1       | 2019-12-16T11:47:00.855944Z 0 [Note] Shutting down plugin 'INNODB_TRX'
db_1       | 2019-12-16T11:47:00.855955Z 0 [Note] Shutting down plugin 'InnoDB'
db_1       | 2019-12-16T11:47:00.856267Z 0 [Note] InnoDB: FTS optimize thread exiting.
db_1       | 2019-12-16T11:47:00.856733Z 0 [Note] InnoDB: Starting shutdown...
db_1       | 2019-12-16T11:47:00.957098Z 0 [Note] InnoDB: Dumping buffer pool(s) to /var/lib/mysql/ib_buffer_pool
db_1       | 2019-12-16T11:47:00.957942Z 0 [Note] InnoDB: Buffer pool(s) dump completed at 191216 11:47:00
db_1       | 2019-12-16T11:47:02.187747Z 0 [Note] InnoDB: Shutdown completed; log sequence number 12440769
db_1       | 2019-12-16T11:47:02.192993Z 0 [Note] InnoDB: Removed temporary tablespace data file: "ibtmp1"
db_1       | 2019-12-16T11:47:02.193049Z 0 [Note] Shutting down plugin 'MEMORY'
db_1       | 2019-12-16T11:47:02.193069Z 0 [Note] Shutting down plugin 'CSV'
db_1       | 2019-12-16T11:47:02.193083Z 0 [Note] Shutting down plugin 'sha256_password'
db_1       | 2019-12-16T11:47:02.193095Z 0 [Note] Shutting down plugin 'mysql_native_password'
db_1       | 2019-12-16T11:47:02.193508Z 0 [Note] Shutting down plugin 'binlog'
db_1       | 2019-12-16T11:47:02.196850Z 0 [Note] mysqld: Shutdown complete
db_1       | 
db_1       | 2019-12-16 11:47:02+00:00 [Note] [Entrypoint]: Temporary server stopped
db_1       | 
db_1       | 2019-12-16 11:47:02+00:00 [Note] [Entrypoint]: MySQL init process done. Ready for start up.
db_1       | 
db_1       | 2019-12-16T11:47:03.085553Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
db_1       | 2019-12-16T11:47:03.086981Z 0 [Note] mysqld (mysqld 5.7.28) starting as process 1 ...
db_1       | 2019-12-16T11:47:03.090574Z 0 [Note] InnoDB: PUNCH HOLE support available
db_1       | 2019-12-16T11:47:03.090590Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
db_1       | 2019-12-16T11:47:03.090594Z 0 [Note] InnoDB: Uses event mutexes
db_1       | 2019-12-16T11:47:03.090597Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
db_1       | 2019-12-16T11:47:03.090601Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.11
db_1       | 2019-12-16T11:47:03.090605Z 0 [Note] InnoDB: Using Linux native AIO
db_1       | 2019-12-16T11:47:03.090963Z 0 [Note] InnoDB: Number of pools: 1
db_1       | 2019-12-16T11:47:03.091099Z 0 [Note] InnoDB: Using CPU crc32 instructions
db_1       | 2019-12-16T11:47:03.092587Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
db_1       | 2019-12-16T11:47:03.103085Z 0 [Note] InnoDB: Completed initialization of buffer pool
db_1       | 2019-12-16T11:47:03.104976Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
db_1       | 2019-12-16T11:47:03.116695Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
db_1       | 2019-12-16T11:47:03.131771Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
db_1       | 2019-12-16T11:47:03.131872Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
db_1       | 2019-12-16T11:47:03.217048Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
db_1       | 2019-12-16T11:47:03.220312Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
db_1       | 2019-12-16T11:47:03.220369Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
db_1       | 2019-12-16T11:47:03.221579Z 0 [Note] InnoDB: Waiting for purge to start
db_1       | 2019-12-16T11:47:03.272063Z 0 [Note] InnoDB: 5.7.28 started; log sequence number 12440769
db_1       | 2019-12-16T11:47:03.273054Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
db_1       | 2019-12-16T11:47:03.273917Z 0 [Note] Plugin 'FEDERATED' is disabled.
db_1       | 2019-12-16T11:47:03.285787Z 0 [Note] InnoDB: Buffer pool(s) load completed at 191216 11:47:03
db_1       | 2019-12-16T11:47:03.294689Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
db_1       | 2019-12-16T11:47:03.294729Z 0 [Note] Skipping generation of SSL certificates as certificate files are present in data directory.
db_1       | 2019-12-16T11:47:03.297329Z 0 [Warning] CA certificate ca.pem is self signed.
db_1       | 2019-12-16T11:47:03.297464Z 0 [Note] Skipping generation of RSA key pair as key files are present in data directory.
db_1       | 2019-12-16T11:47:03.298616Z 0 [Note] Server hostname (bind-address): '*'; port: 3306
db_1       | 2019-12-16T11:47:03.298706Z 0 [Note] IPv6 is available.
db_1       | 2019-12-16T11:47:03.298757Z 0 [Note]   - '::' resolves to '::';
db_1       | 2019-12-16T11:47:03.298810Z 0 [Note] Server socket created on IP: '::'.
db_1       | 2019-12-16T11:47:03.302997Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
db_1       | 2019-12-16T11:47:03.331991Z 0 [Note] Event Scheduler: Loaded 0 events
db_1       | 2019-12-16T11:47:03.332772Z 0 [Note] mysqld: ready for connections.
db_1       | Version: '5.7.28'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
```

制作带ifconfig和ping命令的ubuntu容器

写一个Dockerfile然后build

```dockerfile
FROM ubuntu
RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN  apt-get clean
RUN	 apt-get update
RUN	 apt-get install net-tools
RUN	 apt-get install iputils-ping -y
```

```bash
docker build -t unet .
```

```bash
root@Linux2:/home/weltloose/文档/Dockfiles/unet# docker run -it --rm unet bash
root@c775810bb0bb:/# ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:ac:11:00:02  
          inet addr:172.17.0.2  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:17 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:2186 (2.1 KB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

```


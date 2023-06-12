1.查看端口号被哪个进程占用
sudo lsof -i :8888  或者 lsof -Pnl +M -i4 | grep 22
2.根据进程id查看进程
ps -aux|grep 3128778
3.显示所有状态的tcp链接
netstat -at
4.查看指定 ip 的指定端口是否打开
telnet localhost 22
5.linux 根据进程名查看其占用的端口
netstat -nap | grep pid

# top命令分析
top命令提供了实时的对系统处理器的状态监视.它将显示系统中CPU最“敏感”的任务列表.该命令可以按CPU使用.内存使用和执行时间对任务进行排序

指定进程号： top -p 123
指定用户名：top -u root



top数据展示：
第一行展示：
top - 16:54:09 up 17:01,  1 user,  load average: 0.01, 0.04, 0.05
16:54:09 表示当前时间
up 17:01 表示系统已运行的时间
1 user 当前登录用户的数量
load average: 0.01, 0.04, 0.05 相应最近5、10和15分钟内的平均负载。一般来说负载/逻辑cpu数>5就是高负载
load average数据是每隔5秒钟检查一次活跃的进程数，然后按特定算法计算出的数值。如果这个数除以逻辑CPU的数量，结果高于5的时候就表明系统在超负荷运转了。

第二行展示：
Tasks: 100 total,   1 running,  99 sleeping,   0 stopped,   0 zombie
Tasks — 任务（进程），系统现在共有100个进程，其中处于运行中的有1个，99个在休眠（sleep），stoped状态的有0个，zombie状态（僵尸）的有0个。

第三行展示：
%Cpu(s):  0.8 us,  0.8 sy,  0.0 ni, 98.2 id,  0.2 wa,  0.0 hi,  0.0 si,  0.0 st
这里显示不同模式下所占cpu时间百分比，这些不同的cpu时间表示：

us, user：运行(未调整优先级的) 用户进程的CPU时间百分比

sy，system: 运行内核进程的CPU时间百分比

ni，niced：运行已调整优先级的用户进程的CPU时间百分比

wa，IO wait: 用于等待IO完成的CPU时间百分比，也即CPU的空闲时间

hi：处理硬件中断的CPU时间

si: 处理软件中断的CPU时间

st：这个虚拟机被hypervisor偷去的CPU时间（译注：如果当前处于一个hypervisor下的vm，实际上hypervisor也是要消耗一部分CPU处理时间的）。

第四行展示：
KiB Mem :  1881988 total,   363848 free,   355020 used,  1163120 buff/cache
这一行展示物理内存的使用率
物理内存显示如下:全部可用内存、已使用内存、空闲内存、缓冲内存
按m切换展示
KiB Mem : 34.5/1881988  [|||||||||||||||||||||||||||||||||||                                                                 ]

第5行表示内存交换分区
KiB Swap:        0 total,        0 free,        0 used.  1232900 avail Mem
2031612k total — 交换区总量（2031M）

4k used — 使用的交换区总量（4k）

2031608k free — 空闲交换区总量（2031M）

538676k cached — 缓冲的交换区总量（538M）

# 各进程的状态监控

PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+   COMMAND                                                                             
 1136 root      20   0   50088   1112    780 S   1.0  0.1  12:19.54 rshim  
 PID：进程ID，进程的唯一标识符

USER：进程所有者的实际用户名。

PR：进程的调度优先级。这个字段的一些值是'rt'。这意味这这些进程运行在实时态。

NI：进程的nice值（优先级）。越小的值意味着越高的优先级。负值表示高优先级，正值表示低优先级

VIRT：进程使用的虚拟内存。进程使用的虚拟内存总量，单位kb。VIRT=SWAP+RES

RES：驻留内存大小。驻留内存是任务使用的非交换物理内存大小。进程使用的、未被换出的物理内存大小，单位kb。RES=CODE+DATA

SHR：SHR是进程使用的共享内存。共享内存大小，单位kb

S：这个是进程的状态。它有以下不同的值:

D - 不可中断的睡眠态。

R – 运行态

S – 睡眠态

T – 被跟踪或已停止

Z – 僵尸态

%CPU：自从上一次更新时到现在任务所使用的CPU时间百分比。

%MEM：进程使用的可用物理内存百分比。

TIME+：任务启动后到现在所使用的全部CPU时间，精确到百分之一秒。

COMMAND：运行进程所使用的命令。进程名称（命令名/命令行）

# 多核多CPU监控技巧
在top基本视图中，按键盘数字“1”，可监控每个逻辑CPU的状况：

# linux命令是从哪个文件取值的
Linux 的 `top` 命令主要从 `/proc` 虚拟文件系统（procfs）中的一系列文件获取系统和进程信息。这个虚拟文件系统是由内核提供的，并在内存中存储。

以下是 `top` 命令从 `/proc` 文件系统收集数据的一些关键文件：

1. `/proc/meminfo`：此文件包含有关系统内存使用的详细信息，例如总内存量、可用内存量、已用内存量等。
2. `/proc/cpuinfo`：此文件包含有关 CPU 的详细信息和配置，包括型号、速度、内核等。
3. `/proc/loadavg`：此文件包含有关系统负载的信息，如过去 1 分钟、5 分钟和 15 分钟的系统负载平均值。
4. `/proc/uptime`：此文件表示系统自上次启动以来已运行的时间。
5. `/proc/[PID]`：这是一个包含特定进程（由其进程 ID，PID 表示）信息的目录。例如，`/proc/12345` 对应 PID 为 12345 的进程。它包含了许多指向进程详细信息的文件，如下：
   - `/proc/[PID]/stat`：一个包含进程的详细状态信息的文件，如：状态、优先级、虚拟内存大小、运行时间等。
   - `/proc/[PID]/status`：这个文件提供了一种更友好的方式来访问进程状态信息。它主要是从 `/proc/[PID]/stat` 中获取数据，并提供以键值对的方式显示。
   - `/proc/[PID]/cmdline`：这个文件包含启动进程的完整命令行参数字符串。

事实上，`top` 命令将从 `/proc` 虚拟文件系统中收集更多文件和目录中的详细信息。然而，上述列出的文件和目录是获取数据的核心部分。通过读取这些文件中的信息，`top` 可以实时显示关于系统性能和进程的详细统计信息。

# 内存监控的vmstat命令
使用ps命令查看进程的PID：在终端中输入ps命令，可以列出当前系统中所有的进程信息。找到要查看的进程，记录它的PID。

使用vmstat命令查看内存使用情况：在终端中输入vmstat命令，可以查看系统的内存使用情况。其中，si和so列分别表示从磁盘读取和写入内存的数据量，bi和bo列分别表示从磁盘读取和写入数据的块数，free列表示空闲内存的大小，buff和cache列表示缓存的内存大小。

使用pmap命令查看进程的内存映射：在终端中输入pmap命令，后面跟上要查看的进程的PID，可以查看该进程的内存映射信息。其中，RSS列表示进程占用的物理内存大小，Shared列表示多个进程共享的内存大小，Private列表示进程独占的内存大小。

# pmap命令使用 进程内存分析
pmap工具是linux的工具,能够查看进程用了多少内存,还能分析内存用在上面环节,对于一些长期占用内存居高不下的程序可以分析其行为,命令简单,信息简洁。
查看2013进程号的信息
2103:   /usr/local/qcloud/YunJing/YDLive/YDLive
         Address Perm   Offset Device  Inode   Size   Rss   Pss Referenced Anonymous Swap Locked Mapping
    7f963dcb3000 rw-p 000f1000  fd:01   5035      8     8     8          8         8    0      0 libstdc++.so.6.0.19
    7f963dcb5000 rw-p 00000000  00:00      0     84    12    12         12        12    0      0 
Address: 内存开始地址
Kbytes: 占用内存的字节数（KB）
RSS: 保留内存的字节数（KB）
Dirty: 脏页的字节数（包括共享和私有的）（KB）
Mode: 内存的权限：read、write、execute、shared、private (写时复制)
Mapping: 占用内存的文件、或[anon]（分配的内存）、或[stack]（堆栈）
Offset: 文件偏移
Device: 设备名 (major:minor)

1.mapped 表示该进程映射的虚拟地址空间大小，也就是该进程预先分配的虚拟内存大小，即ps出的vsz
2.writeable/private 表示进程所占用的私有地址空间大小，也就是该进程实际使用的内存大小
3.shared 表示进程和其他进程共享的内存大小

在Linux系统中，每个进程都有自己的虚拟地址空间，进程的内存映射区域就是虚拟地址空间中的一部分。进程的内存映射区域可以分为以下几个部分：

代码段（Text Segment）：存放程序的可执行代码，通常是只读的。

数据段（Data Segment）：存放程序的全局变量和静态变量，通常是可读写的。

堆（Heap）：存放动态分配的内存，通常是可读写的。

栈（Stack）：存放函数调用时的局部变量和函数参数，通常是可读写的。

共享库（Shared Libraries）：存放程序所依赖的共享库，通常是只读的。

内存映射文件（Mapped Files）：存放程序所打开的文件，通常是可读写的。

通过pmap命令可以查看进程的内存映射信息，包括每个映射区域的起始地址、结束地址、权限、偏移量、设备号、节点号、文件名等信息。

# IO性能命令
iotop命令可以实时显示系统中正在进行的磁盘I/O操作，以及这些操作所占用的磁盘I/O带宽和I/O请求队列长度等信息。
Total DISK READ :       0.00 B/s | Total DISK WRITE :       0.00 B/s
Actual DISK READ:       0.00 B/s | Actual DISK WRITE:       0.00 B/s
  TID  PRIO  USER     DISK READ  DISK WRITE  SWAPIN     IO>    COMMAND                                                                                  
    1 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % systemd --switched-root --system --deserialize 22
    2 be/4 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kthreadd]
    4 be/0 root        0.00 B/s    0.00 B/s  0.00 %  0.00 % [kworker/0:

进程列表：显示正在进行磁盘I/O操作的进程列表，包括进程的PID、进程名、用户、I/O带宽、I/O请求队列长度等信息。
总体统计信息：显示系统的磁盘I/O带宽、I/O请求队列长度等信息。

# 网络性能命令
netstat是一个常用的Linux系统网络监控工具，可以用来查看系统的网络连接状态、网络接口统计信息等。netstat命令可以显示系统的网络连接状态，包括TCP连接、UDP连接、UNIX域套接字等。

-a：显示所有连接状态，包括监听状态和非监听状态。
-t：显示TCP连接状态。
-u：显示UDP连接状态。
-n：以数字形式显示IP地址和端口号，而不是域名和服务名。
-p：显示进程ID和进程名。
-e：显示扩展信息，包括网络接口统计信息等。
netstat命令的输出包括以下几个部分：

网络接口统计信息：包括每个网络接口的收发数据包数量、错误数量、丢包数量等。
TCP连接状态：包括每个TCP连接的本地IP地址、本地端口号、远程IP地址、远程端口号、连接状态等。
UDP连接状态：包括每个UDP连接的本地IP地址、本地端口号、远程IP地址、远程端口号等。
UNIX域套接字状态：包括每个UNIX域套接字的类型、状态、inode号等。
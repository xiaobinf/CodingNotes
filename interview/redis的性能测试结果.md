测试结果总结
根据上述性能测试结果发现：
大部分情况下得看服务器的性能和配置，机器越牛逼并且配置越高，redis高并发单机就可以上十几万以上；
一般性的配置，redis提供的高并发，至少到上万是没问题的；
大量网络请求的调用，因为网络本身就有开销，所以redis的吞吐量就不一定那么高了。这点可以从单点和读写分离架构可以看出来。
水平扩容redis读节点，提升度吞吐量
根据主从读写架构分离原则，在其他服务器上另外加redis从节点，单个从节点读请求QPS一般在5万左右，两个redis从节点可以承载整个集群读请求QPS可以在10万以上。

测试结果总结
根据上述性能测试结果发现：
大部分情况下得看服务器的性能和配置，机器越牛逼并且配置越高，redis高并发单机就可以上十几万以上；
一般性的配置，redis提供的高并发，至少到上万是没问题的；
大量网络请求的调用，因为网络本身就有开销，所以redis的吞吐量就不一定那么高了。这点可以从单点和读写分离架构可以看出来。
水平扩容redis读节点，提升度吞吐量
根据主从读写架构分离原则，在其他服务器上另外加redis从节点，单个从节点读请求QPS一般在5万左右，两个redis从节点可以承载整个集群读请求QPS可以在10万以上。
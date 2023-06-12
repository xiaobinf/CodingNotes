# 1000W条数据如何优化？
慢SQL优化：
1. 基于mysql自身优化 select *肯定不行，从并发角度优化，收益比较明显
2. IO层面的优化，只筛选需要的列数据
3. 通过explain命令查看sql执行的过程，从查看结果来看 type=ALL就是命中了全表扫描，这种就需要索引优化。（sql优化解释分析器）
4. 查询的字段尽量通过索引覆盖

同一张表中，如果要查相同的偏移量，但是数据量不同
select a from B limit 10000,10
select a from B limit 10000,100
select a from B limit 10000,1000
select a from B limit 10000,10000
结论：数据量越大，耗时越长

select a from B limit 10,10000
select a from B limit 100,10000
select a from B limit 1000,10000
select a from B limit 10000,10000
结论：偏移量越大，耗时越长
优化方式：将偏移量转为id限定的方式，
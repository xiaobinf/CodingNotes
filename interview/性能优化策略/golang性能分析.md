在 pprof 中，`flat`、`cum` 和 `sum` 是针对函数耗时分析报告的三个重要指标。这些指标可以帮助你对应用程序进行性能分析和调优。

- `Flat (flat%)`: 表示函数本身消耗的 CPU 时间或者内存空间。在 `pprof` 报告中，`Flat` 可以是 CPU 时间或者内存空间。`flat%` 表示该函数占调用链总时间或空间的百分比。

- `Cumulative (cum%)`: 表示自身加上所有被调用函数的总和消耗的 CPU 时间或者内存空间。`cum%` 表示该函数及其子函数占调用链总时间或空间的百分比。

- `Sum`: 对 `flat` 或 `cumulutive` 指标进行求和。

`flat`、`cumulative` 和 `sum` 这三个指标通常应该被同时观察和分析。例如，在 pprof 报告中，您可能会发现某个函数的 `Flat` 很高（这意味着该函数自己占用了大量的 CPU 时间或者内存空间），但 `cumulative` 比较低（这意味着该函数虽然消耗很多资源，但是没有被其他函数调用，或者被调用次数较少）。

相反地，您也可能会发现某个函数的 `cumulative` 很高（这意味着该函数及其子函数占用了大量的 CPU 时间或者内存空间），但 `Flat` 较低（这意味着该函数本身的代码开销较小）。

需要注意的是，`flat` 和 `cumulative` 指标并不一定都是好指标或坏指标。具体情况要根据应用程序的实际情况来评估性能。在进行性能分析和调优时，应该同时观察这三个指标，以便更全面、准确地评估应用程序的性能瓶颈和瓶颈原因。


# 性能捕获方式
go tool pprof http://127.0.0.1:8080/debug/pprof/profile -seconds 20

Fetching profile over HTTP from http://127.0.0.1:8080/debug/pprof/profile
-seconds: open -seconds: no such file or directory
20: open 20: no such file or directory
Fetched 1 source profiles out of 3
Saved profile in /Users/alanfu/pprof/pprof.samples.cpu.003.pb.gz
Type: cpu
Time: Jun 15, 2023 at 9:33pm (CST)
Duration: 30.19s, Total samples = 27.51s (91.13%)
Entering interactive mode (type "help" for commands, "o" for options)
# 在线分析数据
(base) SAWYERSUN-MC0:~ alanfu$ go tool pprof -http=127.0.0.1:8081 /Users/alanfu/pprof/pprof.samples.cpu.003.pb.gz
Serving web UI on http://127.0.0.1:8081

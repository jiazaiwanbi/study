# Worker Pool

实现生产者 -> 3 worker -> 汇总：
- producer 发送 1..10
- 3 个 worker 并发处理（如平方）
- collector 汇总并打印

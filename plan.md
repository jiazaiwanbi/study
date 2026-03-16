# Go + React 8 周学习计划（精准版，第一周激进）

> 你的资料已到位：ES6/MDN/React/Node.js/Koa + go-demos。环境已搭好。
> 节奏：每周四上课，共 8 周。这里给出：**课前 4 天游击战 + 8 周周计划**。

## 总节奏（建议强度，可微调）

- 周一~周三：每天 2–3h 预习
- 周四：上课当天 1–2h 复盘
- 周五~周日：每天 3–5h 实操与作业

仓库位置：`/study/go-demos-master/go-demos`

---

# 课前 4 天冲刺（Week 0）

> 目标：把 Go 语法和并发基础 + React/ES6 入门打牢，避免上课听不懂。

### Day -4（今天）

**Go**
- 跑：basic/00-module-package, 01-variables, 02-slice-map
- 练：数组/切片/Map 基础
**JS/React**
- ES6：let/const、箭头函数、模板字符串、解构
- MDN：Array / Map / Set

**代码题**
1) `dedupSort([]int) []int`：去重+排序
2) `countFreq([]string) map[string]int`
3) JS：实现 `unique(arr)` + `countBy(arr)`

---

### Day -3
**Go**
- 跑：basic/03-struct-receiver, 04-closure, 05-interface
- 理解：值接收 vs 指针接收
**JS/React**
- ES6：Promise + async/await
- MDN：Promise, Fetch

**代码题**
1) Task 结构体：`Complete()`、`Bump()`、实现 Stringer
2) 自定义接口：车/自行车 speedUp/speedDown
3) JS：写 `fetchJson(url)` 包装（含错误处理）

---

### Day -2
**Go**
- 跑：basic/06-control, 07-errors, 09-unit-test, 20-testing/basic1
- 修复：basic/20-testing/basic1/mathutils/add_test.go 里的错误期望值
**JS/React**
- React 文档：Getting Started + Main Concepts(Components/Props/State)

**代码题**
1) 错误链：Repo->Service -> errors.Is/As
2) 写 2 个表驱动单测（table-driven tests）
3) React：做一个最小 Counter 组件（useState）

---

### Day -1（课前一天）
**Go**
- 跑：basic/11-goroutine, 12-channel, 13-select（至少 01/02）
- 理解：goroutine 泄漏与 close 通道
**JS/React**
- React：useEffect + 表单受控组件

**代码题**
1) Worker Pool：生产者 -> 3 worker -> 汇总
2) 超时 select：100ms 超时则退出
3) React：表单 + 列表渲染（Todo mini）

---

# 8 周周计划（每周四上课）

## Week 1（激进：Go 并发 + React 基础强化）
**预习（Mon-Wed）**
- Go：basic/11~15（goroutine/channel/select/mutex/context）
- React：组件拆分、props、state、useEffect

**课后实战（Fri-Sun）**
**代码题**
1) `RunTasks(ctx, tasks)`：任何任务失败立即 cancel
2) 心跳检测（参照 basic/19-heartbeat）
3) React Todo：增删改查 + 本地持久化（localStorage）
4) Go API：提供 Todo REST 接口（内存存储+mutex）

交付：`/study/week01/`（Go 后端 + React 前端）

---

## Week 2（Go HTTP / Gin 基础 + React 表单与请求）
**预习**
- zhangsixue2026/week03/go-http/1~5
- gin-basic/router、json、form、query

**代码题**
1) Gin：实现 CRUD（/users）+ 参数校验
2) React：表单 + fetch API 封装（错误统一处理）
3) 联调：React -> Gin API

交付：`/study/week02/`

---

## Week 3（中间件 / 文件上传 / Cookie & JWT）
**预习**
- gin-basic/middleware1~3, file, cookie, cookie-jwt

**代码题**
1) Gin 中间件：记录耗时 + 请求 ID
2) 文件上传接口（保存到本地）
3) Cookie/JWT 登录 & 鉴权
4) React：登录/退出 + 受保护路由

交付：`/study/week03/`

---

## Week 4（测试 + 日志 + 可靠性）
**预习**
- basic/09-unit-test, 10-logger
- go test -race / coverage

**代码题**
1) API 单测 + Mock
2) 覆盖率 >= 60%
3) 统一日志（请求日志 + 错误日志）

交付：`/study/week04/`

---

## Week 5（并发模式进阶）
**预习**
- basic/16-errGroup, 17-pipeline, 18-fanout-fanin

**代码题**
1) 实现可取消 pipeline（3-stage）
2) fan-out 多 worker 处理任务
3) 写一份并发模式对比笔记

交付：`/study/week05/`

---

## Week 6（React 进阶 + 状态管理）
**预习**
- React 官方文档：Hooks 深入、Context
- ES6：模块化、解构、默认参数

**代码题**
1) Context 管理全局状态
2) 自定义 Hook：useFetch/useDebounce
3) 搜索 + 分页列表

交付：`/study/week06/`

---

## Week 7（Node.js / Koa 基础对比）
**预习**
- Node.js API 文档
- Koa 教程（了解中间件模型）

**代码题**
1) Koa 写一个简单 REST API（对比 Gin）
2) 写一份 Gin vs Koa 对比总结

交付：`/study/week07/`

---

## Week 8（综合项目收尾）
**项目目标**
- Go + React 全栈小项目
- 含：登录、列表、CRUD、分页、基础权限

**交付**
- `/study/week08/` 完整项目
- README（运行方式 + 功能说明）

---

# 备注
- 如果你每天可投入时长 >5h，我可以把 Week1/Week2 再细化到**每日粒度**。
- 你希望“第一周更激进”，已在 Week0 + Week1 加重并发/实战强度。

需要我把 Week1/Week2 的代码模板直接生成到 /study 吗？

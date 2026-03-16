# 错误链

实现 Repo -> Service -> Handler 的错误链包装：
- 用 fmt.Errorf("...: %w", err)
- 在 main 里用 errors.Is / errors.As 判断

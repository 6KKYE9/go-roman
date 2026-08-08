# go-roman

零依赖的罗马数字转换命令行工具，支持 1~3999。

## 用法

```bash
go-roman to 1994      # 输出 MCMXCIV
go-roman from MCMXCIV # 输出 1994
```

- `to`：整数转罗马数字
- `from`：罗马数字转整数，非法字符会报错

核心逻辑在 `roman.go`，库函数 `ToRoman` / `FromRoman` 可直接调用。

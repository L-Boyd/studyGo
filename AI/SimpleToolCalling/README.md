# SimpleToolCalling

一个用 Go 实现的命令行 AI 对话程序，演示了大模型（DeepSeek）的 **工具调用（Tool Calling / Function Calling）** 机制。

程序内置模拟的"获取温度"工具（永远返回 25 摄氏度）和"获取时间"工具（永远返回 12:00 PM），大模型在对话中可以根据用户意图自动决定是否调用工具，并把工具返回的结果用自然语言组织后回复用户。

## 功能特性

- 命令行交互式多轮对话，支持上下文
- 接入 DeepSeek 大模型（`deepseek-v4-flash`）
- 实现 OpenAI 兼容的 Tool Calling 协议
- 内置 `get_temperature` 工具，模拟获取城市温度（固定返回 25°C）
- 内置 `get_current_time` 工具，模拟获取当前时间（固定返回 12:00 PM）
- 通过 JSON 配置文件管理 API Key，避免硬编码

## 项目结构

```
SimpleToolCalling/
├── main.go              # 主程序：命令行交互循环 + 工具调用编排
├── client.go            # DeepSeek API 客户端 + 配置加载
├── tools.go             # 工具定义与实现（温度工具、时间工具）
├── config.json          # 实际配置文件（需自行填写，不提交到仓库）
├── config.json.example  # 配置文件模板
└── README.md
```

## 快速开始

### 1. 环境要求

- Go 1.25+
- 一个 DeepSeek API Key（在 [platform.deepseek.com](https://platform.deepseek.com/) 获取）

### 2. 配置

复制配置模板并填入你的 API Key：

```bash
cp config.json.example config.json
```

编辑 `config.json`：

```json
{
    "api_key": "sk-你的真实key",
    "base_url": "https://api.deepseek.com",
    "model": "deepseek-v4-flash"
}
```

| 字段       | 说明                          | 默认值                       |
| ---------- | ----------------------------- | ---------------------------- |
| `api_key`  | DeepSeek API Key（必填）      | -                            |
| `base_url` | API 基础地址                  | `https://api.deepseek.com`   |
| `model`    | 使用的模型名                  | `deepseek-v4-flash`          |

### 3. 运行

在项目根目录 `studyGo` 下执行（程序使用相对路径读取配置）：

```bash
go run ./AI/SimpleToolCalling/
```

### 4. 使用示例

启动后会进入交互式对话：

```
========================================
    DeepSeek 工具调用模拟器
========================================
可用的工具:
  - get_temperature:  获取指定城市的当前温度
  - get_current_time: 获取当前的系统时间
----------------------------------------
输入 'quit' 或 'exit' 退出程序
----------------------------------------

你: 北京现在多少度？

AI:   [正在调用工具: get_temperature]
  [工具返回: {"city":"北京","temperature":25,"unit":"摄氏度"}]

AI: 北京现在的温度是 25 摄氏度。

你: 上海呢？

AI:   [正在调用工具: get_temperature]
  [工具返回: {"city":"上海","temperature":25,"unit":"摄氏度"}]

AI: 上海现在的温度同样是 25 摄氏度。

你: 现在几点了？

AI:   [正在调用工具: get_current_time]
  [工具返回: {"time":"12:00 PM","timezone":"UTC+8"}]

AI: 现在是 12:00 PM（UTC+8 时区）。

你: quit
再见！
```

## 工作原理

整个工具调用流程如下：

```
用户输入
   │
   ▼
┌─────────────────────────────┐
│  构造请求（messages + tools）│
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│     调用 DeepSeek API        │
└──────────────┬──────────────┘
               │
        ┌──────┴───────┐
        │ finish_reason │
        │   是什么？    │
        └──────┬───────┘
               │
      ┌────────┴────────┐
      ▼                 ▼
  tool_calls         stop
      │                 │
      ▼                 ▼
┌──────────┐     直接输出回复
│ 执行工具  │
│ (返回25°C)│
└─────┬────┘
      │
      ▼
┌──────────────────────┐
│ 工具结果加入 messages │
│ 再次调用大模型        │
└──────────┬───────────┘
           │
           ▼
      输出最终回复
```

### 工具调用协议

请求中传给大模型的工具定义（OpenAI 兼容格式）：

```json
{
    "type": "function",
    "function": {
        "name": "get_temperature",
        "description": "获取指定城市的当前温度",
        "parameters": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string",
                    "description": "要查询温度的城市名称"
                }
            },
            "required": ["city"]
        }
    }
}
```

大模型返回的 `tool_calls` 中会包含工具名和参数，程序解析后执行本地工具，再把结果以 `role: "tool"` 的消息回传给大模型，由大模型生成最终的自然语言回复。

## 扩展工具

项目通过 `Tool` 接口实现工具的插件化，新增工具只需三步：

1. 在 `tools.go` 中实现 `Tool` 接口（`Name`、`Description`、`Parameters`、`Execute`）
2. 在 `main.go` 中将新工具加入 `tools` 切片
3. 在 `main.go` 中将新工具定义加入 `toolDefs` 切片

```go
type Tool interface {
    Name() string
    Description() string
    Parameters() map[string]interface{}
    Execute(args map[string]interface{}) (string, error)
}
```

例如可以扩展一个"获取时间"、"发送邮件"等工具。

## 备注

- `get_temperature` 工具是**模拟实现**，无论查询哪个城市都固定返回 25°C，仅用于演示工具调用流程
- `get_current_time` 工具是**模拟实现**，无论何时调用都固定返回 12:00 PM，仅用于演示工具调用流程
- 对话历史保留在内存中，程序退出后清空
- `config.json` 包含敏感信息，建议加入 `.gitignore`，不要提交到代码仓库

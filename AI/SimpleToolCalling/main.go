package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 加载配置
	config, err := LoadConfig("AI/SimpleToolCalling/config.json")
	if err != nil {
		fmt.Println("加载配置失败:", err)
		fmt.Println("请确保 config.json 文件存在并配置了正确的 api_key")
		return
	}

	if config.APIKey == "your-deepseek-api-key-here" || config.APIKey == "" {
		fmt.Println("请先在 config.json 中配置你的 DeepSeek API Key")
		return
	}

	fmt.Println("========================================")
	fmt.Println("    DeepSeek 工具调用模拟器")
	fmt.Println("========================================")
	fmt.Println("可用的工具:")
	fmt.Println("  - get_temperature:  获取指定城市的当前温度")
	fmt.Println("  - get_current_time: 获取当前的系统时间")
	fmt.Println("----------------------------------------")
	fmt.Println("输入 'quit' 或 'exit' 退出程序")
	fmt.Println("----------------------------------------")

	// 初始化工具
	tempTool := &GetTemperatureTool{}
	timeTool := &GetCurrentTimeTool{}
	tools := []Tool{tempTool, timeTool}

	// 构建传给大模型的工具定义
	toolDefs := []ToolDef{
		{
			Type: "function",
			Function: ToolFuncDef{
				Name:        tempTool.Name(),
				Description: tempTool.Description(),
				Parameters:  tempTool.Parameters(),
			},
		},
		{
			Type: "function",
			Function: ToolFuncDef{
				Name:        timeTool.Name(),
				Description: timeTool.Description(),
				Parameters:  timeTool.Parameters(),
			},
		},
	}

	// 初始化客户端
	client := NewDeepSeekClient(config)

	// 对话历史
	messages := []Message{
		{
			Role: "system",
			Content: "你是一个有用的助手。当用户询问温度时，请调用 get_temperature 工具来获取温度信息。" +
				"当用户询问当前时间时，请调用 get_current_time 工具来获取时间信息。" +
				"获取到结果后，请用自然语言告诉用户。",
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for {
		fmt.Print("\n你: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}
		if userInput == "quit" || userInput == "exit" {
			fmt.Println("再见！")
			break
		}

		// 添加用户消息到历史
		messages = append(messages, Message{
			Role:    "user",
			Content: userInput,
		})

		// 调用大模型
		fmt.Print("\nAI: ")
		resp, err := client.Chat(messages, toolDefs)
		if err != nil {
			fmt.Println("调用失败:", err)
			// 移除刚才添加的用户消息，避免历史污染
			messages = messages[:len(messages)-1]
			continue
		}

		if len(resp.Choices) == 0 {
			fmt.Println("未返回有效响应")
			continue
		}

		choice := resp.Choices[0]
		assistantMsg := choice.Message

		// 打印文本内容
		if content, ok := assistantMsg.Content.(string); ok && content != "" {
			fmt.Println(content)
		}

		// 将助手消息加入历史（保留 tool_calls）
		messages = append(messages, assistantMsg)

		// 如果模型要求调用工具
		if choice.FinishReason == "tool_calls" && len(assistantMsg.ToolCalls) > 0 {
			for _, toolCall := range assistantMsg.ToolCalls {
				fmt.Printf("  [正在调用工具: %s]\n", toolCall.Function.Name)

				// 解析工具参数
				var args map[string]interface{}
				if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
					fmt.Println("  [解析工具参数失败:", err, "]")
					continue
				}

				// 执行工具
				var toolResult string
				found := false
				for _, tool := range tools {
					if tool.Name() == toolCall.Function.Name {
						result, err := tool.Execute(args)
						if err != nil {
							fmt.Println("  [工具执行失败:", err, "]")
							toolResult = fmt.Sprintf(`{"error": "%v"}`, err)
						} else {
							toolResult = result
							fmt.Printf("  [工具返回: %s]\n", toolResult)
						}
						found = true
						break
					}
				}

				if !found {
					fmt.Printf("  [未找到工具: %s]\n", toolCall.Function.Name)
					toolResult = fmt.Sprintf(`{"error": "工具 %s 不存在"}`, toolCall.Function.Name)
				}

				// 将工具结果加入历史
				messages = append(messages, Message{
					Role:       "tool",
					Content:    toolResult,
					ToolCallID: toolCall.ID,
				})
			}

			// 再次调用大模型，让它根据工具结果生成最终回复
			fmt.Print("\nAI: ")
			resp2, err := client.Chat(messages, toolDefs)
			if err != nil {
				fmt.Println("调用失败:", err)
				continue
			}

			if len(resp2.Choices) > 0 {
				finalMsg := resp2.Choices[0].Message
				if content, ok := finalMsg.Content.(string); ok && content != "" {
					fmt.Println(content)
				}
				messages = append(messages, finalMsg)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入错误:", err)
	}
}

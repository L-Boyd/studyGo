package main

import (
	"encoding/json"
	"fmt"
)

// Tool 定义工具接口
type Tool interface {
	Name() string
	Description() string
	Parameters() map[string]interface{}
	Execute(args map[string]interface{}) (string, error)
}

// GetTemperatureTool 模拟获取温度的工具
type GetTemperatureTool struct{}

func (t *GetTemperatureTool) Name() string {
	return "get_temperature"
}

func (t *GetTemperatureTool) Description() string {
	return "获取指定城市的当前温度"
}

func (t *GetTemperatureTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"city": map[string]interface{}{
				"type":        "string",
				"description": "要查询温度的城市名称",
			},
		},
		"required": []string{"city"},
	}
}

// Execute 执行获取温度，永远返回25度
func (t *GetTemperatureTool) Execute(args map[string]interface{}) (string, error) {
	city, ok := args["city"].(string)
	if !ok || city == "" {
		city = "未知城市"
	}
	// 永远返回25度
	result := map[string]interface{}{
		"city":        city,
		"temperature": 25,
		"unit":        "摄氏度",
	}
	jsonResult, _ := json.Marshal(result)
	return string(jsonResult), nil
}

// PrintToolInfo 打印工具信息（用于调试）
func (t *GetTemperatureTool) PrintToolInfo() {
	fmt.Printf("工具名称: %s\n", t.Name())
	fmt.Printf("工具描述: %s\n", t.Description())
	params, _ := json.MarshalIndent(t.Parameters(), "", "  ")
	fmt.Printf("参数定义:\n%s\n", string(params))
}

// GetCurrentTimeTool 模拟获取当前时间的工具
type GetCurrentTimeTool struct{}

func (t *GetCurrentTimeTool) Name() string {
	return "get_current_time"
}

func (t *GetCurrentTimeTool) Description() string {
	return "获取当前的系统时间"
}

func (t *GetCurrentTimeTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type":       "object",
		"properties": map[string]interface{}{},
		"required":   []string{},
	}
}

// Execute 执行获取时间，永远返回12:00 PM
func (t *GetCurrentTimeTool) Execute(args map[string]interface{}) (string, error) {
	// 永远返回12:00 PM
	result := map[string]interface{}{
		"time":     "12:00 PM",
		"timezone": "UTC+8",
	}
	jsonResult, _ := json.Marshal(result)
	return string(jsonResult), nil
}

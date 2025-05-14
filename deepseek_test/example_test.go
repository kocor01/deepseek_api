package deepseek_test

import (
	"fmt"
	"github.com/kocor01/deepseek_api"
	"testing"
)

func TestDeepSeekAPI(t *testing.T) {
	// 从环境变量获取API密钥
	apiKey := "DEEPSEEK_API_KEY"

	// 创建客户端
	client := deepseek_api.NewClient(deepseek_api.Config{
		APIKey:     apiKey,
		MaxRetries: 3,
		Timeout:    30,
	})

	// 使用QuickChat便捷方法
	response, err := client.QuickChat(
		"你是个乐于助人的助手",
		"2025年5月ai最新资讯top3",
	)
	if err != nil {
		t.Fatalf("Error in QuickChat: %v", err)
	}

	fmt.Println("QuickChat Assistant response:")
	fmt.Println(response)

	// 使用QuickChatWebSearch便捷方法
	response, err = client.QuickChatWebSearch(
		"你是个乐于助人的助手",
		"2025年5月ai最新资讯top3",
	)
	if err != nil {
		t.Fatalf("Error in QuickChatWebSearch: %v", err)
	}

	fmt.Println("QuickChatWebSearch Assistant response:")
	fmt.Println(response)

	// 使用完整的Chat方法
	fullReq := deepseek_api.Request{
		Messages: []deepseek_api.Message{
			{Role: "system", Content: "你是个乐于助人的助手"},
			{Role: "user", Content: "你好"},
		},
		Model:       "deepseek-chat",
		MaxTokens:   1024,
		Temperature: 0.7,
		WebSearch:   true,
	}

	fullResp, err := client.Chat(fullReq)
	if err != nil {
		t.Fatalf("Error in Chat: %v", err)
	}

	fmt.Println("\nDetailed response:")
	fmt.Printf("ID: %s\n", fullResp.ID)
	fmt.Printf("Model: %s\n", fullResp.Model)
	if len(fullResp.Choices) > 0 {
		fmt.Printf("Assistant: %s\n", fullResp.Choices[0].Message.Content)
	}
	fmt.Printf("Token usage: Prompt=%d, Completion=%d, Total=%d\n",
		fullResp.Usage.PromptTokens,
		fullResp.Usage.CompletionTokens,
		fullResp.Usage.TotalTokens,
	)
}

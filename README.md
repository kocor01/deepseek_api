# 创建请求客户端
```
// API密钥
apiKey := "DEEPSEEK_API_KEY"

// 创建客户端
client := deepseek_api.NewClient(deepseek_api.Config{
    APIKey:     apiKey,
    MaxRetries: 3,
    Timeout:    30,
})
```

# 调用方法：

## QuickChat()：简单快捷的单次对话
```
// 使用QuickChat便捷方法
response, err := client.QuickChat(
    "你是个乐于助人的助手",
    "你好",
)
if err != nil {
    t.Fatalf("Error in QuickChat: %v", err)
}

fmt.Println("Assistant response:")
fmt.Println(response)
```

## QuickChatWebSearch()：简单快捷的单次对话，启用联网搜索
```
// 使用QuickChat便捷方法
response, err := client.QuickChatWebSearch(
    "你是个乐于助人的助手",
    "你好",
)
if err != nil {
    t.Fatalf("Error in QuickChat: %v", err)
}

fmt.Println("Assistant response:")
fmt.Println(response)
```

## Chat()：完整的请求控制
```
// 使用完整的Chat方法
fullReq := deepseek_api.Request{
    Messages: []deepseek_api.Message{
        {Role: "system", Content: "你是个乐于助人的助手"},
        {Role: "user", Content: "你好"},
    },
    MaxTokens:   1024,
    Temperature: 0.7,
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
```
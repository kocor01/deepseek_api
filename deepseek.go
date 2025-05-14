package deepseek_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://api.deepseek.com"
	defaultTimeout = 30 * time.Second
)

// Client 表示DeepSeek API客户端
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	maxRetries int
}

// NewClient 创建一个新的DeepSeek客户端
func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}
	if config.Timeout == 0 {
		config.Timeout = int(defaultTimeout.Seconds())
	}

	return &Client{
		apiKey:  config.APIKey,
		baseURL: config.BaseURL,
		httpClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
		maxRetries: config.MaxRetries,
	}
}

// Chat 发送聊天请求并返回响应
func (c *Client) Chat(req Request) (*Response, error) {
	// 设置默认值
	if req.Model == "" {
		req.Model = "deepseek-chat"
	}
	if req.MaxTokens == 0 {
		req.MaxTokens = 2048
	}
	if req.Temperature == 0 {
		req.Temperature = 1
	}
	if req.TopP == 0 {
		req.TopP = 1
	}

	// 准备请求体
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	// 创建HTTP请求
	httpReq, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/chat/completions", c.baseURL),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// 设置请求头
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Authorization", "Bearer "+c.apiKey)

	// 发送请求(带重试逻辑)
	var resp *http.Response
	for i := 0; i <= c.maxRetries; i++ {
		resp, err = c.httpClient.Do(httpReq)
		if err == nil {
			break
		}
		if i == c.maxRetries {
			return nil, fmt.Errorf("max retries exceeded, last error: %w", err)
		}
		time.Sleep(time.Duration(i+1) * time.Second) // 指数退避
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var apiResp Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	// 检查API错误
	if apiResp.Error.Message != "" {
		return nil, fmt.Errorf("API error: %s (type: %s)", apiResp.Error.Message, apiResp.Error.Type)
	}

	return &apiResp, nil
}

// QuickChat 快速发送单条消息的便捷方法
func (c *Client) QuickChat(systemPrompt, userMessage string) (string, error) {
	req := Request{
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userMessage},
		},
	}

	resp, err := c.Chat(req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return resp.Choices[0].Message.Content, nil
}

// QuickChatWebSearch 快速发送单条消息的便捷方法,启用联网搜索
func (c *Client) QuickChatWebSearch(systemPrompt, userMessage string) (string, error) {
	req := Request{
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userMessage},
		},
		WebSearch: true,
	}

	resp, err := c.Chat(req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return resp.Choices[0].Message.Content, nil
}

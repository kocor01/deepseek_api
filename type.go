package deepseek_api

// Message 表示对话中的一条消息
type Message struct {
	Role    string `json:"role"`    // 角色: system/user/assistant
	Content string `json:"content"` // 消息内容
}

// Request 表示API请求结构
type Request struct {
	Messages         []Message `json:"messages"`          // 对话消息列表
	Model            string    `json:"model"`             // 使用的模型
	MaxTokens        int       `json:"max_tokens"`        // 最大token数
	Temperature      float64   `json:"temperature"`       // 温度参数
	TopP             float64   `json:"top_p"`             // Top-p采样
	FrequencyPenalty float64   `json:"frequency_penalty"` // 频率惩罚
	PresencePenalty  float64   `json:"presence_penalty"`  // 存在惩罚
	Stop             []string  `json:"stop,omitempty"`    // 停止词
	WebSearch        bool      `json:"web_search"`        // 是否启用联网搜索
}

// Response 表示API响应结构
type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

// Config 包含API客户端配置
type Config struct {
	APIKey     string // DeepSeek API密钥
	BaseURL    string // API基础URL(可选)
	MaxRetries int    // 最大重试次数(可选)
	Timeout    int    // 超时时间(秒)(可选)
}

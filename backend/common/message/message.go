package message

const (
	MessageTypeText  = 1
	MessageTypeImage = 2
	MessageTypeFile  = 3
	MessageTypeVoice = 4
	MessageTypeVideo = 5
)

type LoginRequest struct {
	AccountID int64  `json:"account_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Nickname string `json:"nickname,omitempty" `
}

type SendMessageRequest struct {
	SenderID    int64  `json:"sender_id" binding:"required"`
	ReceiverID  int64  `json:"receiver_id" binding:"required"`
	Content     string `json:"content" binding:"required"`
	MessageType int    `json:"message_type" binding:"required,oneof=1 2 3 4 5"`
	FileSize    int64  `json:"file_size,omitempty"`
	FileName    string `json:"file_name,omitempty"`
}

type BaseResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

type LoginResponse struct {
	AccountID int64  `json:"account_id"`
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"` // Token过期时间
}

type RegisterResponse struct {
	AccountID int64  `json:"account_id"` // 系统生成的用户ID,对外展示
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type SendMessageResponse struct {
	MessageID int64 `json:"message_id"`
	Timestamp int64 `json:"timestamp"`
	Status    int   `json:"status"` // 1-发送中 2-已送达 3-发送失败
}

type PageInfo struct {
	Page      int `json:"page"`       // 当前是第几页（从1开始）
	PageSize  int `json:"page_size"`  // 每页显示多少条
	Total     int `json:"total"`      // 总记录数
	TotalPage int `json:"total_page"` // 总页数
}

type ListResponse struct {
	BaseResponse
	PageInfo *PageInfo `json:"page_info,omitempty"`
}

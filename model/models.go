package model

import "time"

type Prodduct struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type User struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"colun:password_hash"`
	CreateTime   time.Time `json:"create_time" gorm:"autoCreateTime"`
}

type Ohlcv struct {
	ID        int64   `json:"id" gorm:"primaryKey"`
	Symbol    string  `json:"symbol"`
	Timeframe string  `json:"timeframe"`
	Timestamp string  `json:"timestamp"`
	Open      float64 `json:"open" gorm:"type:decimal(18,8)"`
	High      float64 `json:"high" gorm:"type:decimal(18,8)"`
	Low       float64 `json:"low" gorm:"type:decimal(18,8)"`
	Close     float64 `json:"close" gorm:"type:decimal(18,8)"`
	Volume    float64 `json:"volume" gorm:"type:decimal(18,8)"`
}

type Order struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	UserID      int64     `json:"user_id"`
	Symbol      string    `json:"symbol"`
	Timeframe   string    `json:"timeframe"`
	Side        string    `json:"side" gorm:"type:enum('BUY','SELL')"`
	EntryPrice  float64   `json:"entry_price" gorm:"type:decimal(18,8)"`
	StopLoss    float64   `json:"stop_loss" gorm:"type:decimal(18,8)"`
	TakeProfit  float64   `json:"take_profit" gorm:"type:decimal(18,8)"`
	OpenTime    time.Time `json:"open_time" gorm:"autoCreateTime"`
	ExitPrice   float64   `json:"exit_price" gorm:"type:decimal(18,8)"`
	CloseTime   time.Time `json:"close_time" gorm:"autoCreateTime"`
	CloseReason string    `json:"close_reason" gorm:"type:enum('TP_HIT','SL_HIT','OPPOSITE_SIGNAL')"`
	Result      string    `json:"result" gorm:"type:enum('WIN','LOSE','PENDING');default:PENDING"`
}

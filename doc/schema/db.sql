CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ohlcv (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    symbol VARCHAR(20) NOT NULL,
    timeframe VARCHAR(20) NOT NULL,
    timestamp BIGINT NOT NULL,
    open DECIMAL(18,8),
    high DECIMAL(18,8),
    low DECIMAL(18,8),
    close DECIMAL(18,8),
    volume DECIMAL(18,8),
    UNIQUE KEY uk_ohlcv (symbol, timeframe,timestamp)
);

CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    symbol VARCHAR(20) NOT NULL,
    timeframe VARCHAR(10) NOT NULL,
    
    side ENUM('BUY', 'SELL') NOT NULL,
    entry_price DECIMAL(18,8) NOT NULL,
    
    stop_loss DECIMAL(18,8) NOT NULL,
    take_profit DECIMAL(18,8) NOT NULL,
    
    open_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    exit_price DECIMAL(18,8),
    close_time TIMESTAMP,
    close_reason ENUM('TP_HIT', 'SL_HIT', 'OPPOSITE_SIGNAL') DEFAULT NULL,
    
    result ENUM('WIN', 'LOSE', 'PENDING') DEFAULT 'PENDING',
    
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_user_result (user_id, result)
);
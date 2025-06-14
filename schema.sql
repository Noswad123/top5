CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE lists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    title TEXT,
    category TEXT,
    validated BOOLEAN DEFAULT FALSE,
    auto_validate BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE list_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    list_id INTEGER,
    rank INTEGER,
    content TEXT,
    validation_id INTEGER,
    FOREIGN KEY (list_id) REFERENCES lists(id)
);

CREATE TABLE validation_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    item_id INTEGER,
    category TEXT,
    is_valid BOOLEAN,
    confidence REAL,
    explanation TEXT,
    model_version TEXT,
    validated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE override_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    item_id INTEGER,
    requested_by INTEGER,
    reason TEXT,
    status TEXT CHECK (status IN ('pending', 'approved', 'rejected')) DEFAULT 'pending',
    reviewed_by INTEGER,
    reviewed_at DATETIME
);

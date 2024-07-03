ALTER TABLE blog_categories ADD CONSTRAINT UC_name UNIQUE (name);

CREATE TABLE blogs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    content TEXT,
    author_id INT NOT NULL,
    category VARCHAR(255),
    status ENUM('Pending', 'Active', 'Inactive') DEFAULT 'Pending',
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (author_id) REFERENCES users(id),
    CONSTRAINT fk_category FOREIGN KEY (category) REFERENCES blog_categories(name)
);

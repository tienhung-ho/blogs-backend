ALTER TABLE  blog_categories MODIFY COLUMN status ENUM('Pending', 'Active', 'Inactive') DEFAULT 'Pending';

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status ON users(status);


CREATE INDEX idx_blog_categories_name ON blog_categories(name);
CREATE INDEX idx_blog_categories_status ON blog_categories(status);
CREATE INDEX idx_blog_categories_parentcategory ON blog_categories(parentcategory);


CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_status ON roles(status);


CREATE INDEX idx_permissions_name ON permissions(name);
CREATE INDEX idx_permissions_status ON permissions(status);

CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);


CREATE INDEX idx_accounts_username ON accounts(username);
CREATE INDEX idx_accounts_email ON accounts(email);
CREATE INDEX idx_accounts_role_id ON accounts(role_id);
CREATE INDEX idx_accounts_status ON accounts(status);


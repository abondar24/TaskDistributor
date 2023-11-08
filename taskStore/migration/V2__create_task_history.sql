CREATE TABLE IF NOT EXISTS task (
                                    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
                                    task_id VARCHAR(128) NOT NULL,
                                    status VARCHAR(16) NOT NULL,
                                    updated_at DATE NOT NULL
);
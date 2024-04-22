CREATE DATABASE backend_dev;
USE backend_dev;

CREATE TABLE stories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    story blob(65536) NOT NULL
);

INSERT INTO stories (title, story)
VALUES (
    'Robin Hood',
    LOAD_FILE('/var/lib/mysql-files/stories/story_1.txt')
);


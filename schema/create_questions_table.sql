CREATE TABLE questions (
  id INT PRIMARY KEY,
  game_id INT,
  category VARCHAR(255),
  round TEXT,
  prompt TEXT,
  answer TEXT,
  monetary_value INT
);

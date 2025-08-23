-- +goose Up
CREATE TABLE user_skills (
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  skill_id UUID REFERENCES skills(id) ON DELETE CASCADE,
  UNIQUE(user_id, skill_id)
);

-- +goose Down
DROP TABLE user_skills;

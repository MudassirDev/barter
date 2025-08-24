-- name: CreateSkill :one
INSERT INTO skills (
  id, title, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: CreateUserSkill :one
INSERT INTO user_skills (
  user_id, skill_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetSkillByTitle :one
SELECT * FROM skills
WHERE title = ?;

-- name: GetSkillsByUserID :many
SELECT *
FROM user_skills
INNER JOIN skills ON skills.id = user_skills.skill_id
WHERE user_skills.user_id = ?;

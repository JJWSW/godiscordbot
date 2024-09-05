-- name: GetSchedule :many
SELECT * FROM schedule;

-- name: GetScheduleById :one
SELECT * FROM schedule WHERE id = $1;

-- name: GetScheduleToday :one
WITH previous_time AS (
    SELECT *
    FROM schedule
    WHERE running_time BETWEEN NOW() - INTERVAL '1 hour 30 minutes' AND NOW()
    ORDER BY running_time DESC
    LIMIT 1
),
next_time AS (
    SELECT *
    FROM schedule
    WHERE running_time > NOW()
    ORDER BY running_time ASC
    LIMIT 1
)
SELECT *
FROM previous_time
UNION ALL
SELECT *
FROM next_time
WHERE NOT EXISTS (SELECT 1 FROM previous_time)
LIMIT 1;

-- name: CreateSchedule :execresult
INSERT INTO schedule (episode, title, guest, description, running_time) VALUES ($1, $2, $3, $4, $5);
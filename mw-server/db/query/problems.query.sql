-- name: CreateProblem :one
INSERT INTO problems(
    created_at,
    updated_at,
    description,
    is_done,
    owner_uuid,
    day_report_uuid
) VALUES (
    @created_at,
    @updated_at,
    @description,
    @is_done,
    @owner_uuid,
    @day_report_uuid
) RETURNING *, (SELECT name FROM users WHERE uuid = @owner_uuid) AS owner_name;

-- name: GetListProblemsByDayReportId :many
SELECT * FROM problems
WHERE problems.day_report_uuid = @day_report_uuid
ORDER BY created_at;

-- name: UpdateProblem :one
UPDATE problems
SET
updated_at = coalesce(sqlc.narg('updated_at'), updated_at),
description = coalesce(sqlc.narg('description'), description),
is_done = coalesce(sqlc.narg('is_done'), is_done)
WHERE problems.uuid = sqlc.arg('uuid')
RETURNING *, (SELECT name FROM users WHERE problems.owner_uuid = users.uuid) AS owner_name;

-- name: DeleteProblem :exec
DELETE FROM problems
WHERE uuid = @problem_uuid;

-- name: GetProblemsByDayReportUuids :many
SELECT *
FROM problems
WHERE problems.day_report_uuid = ANY($1::UUID[]);

-- name: GetIsUserHavingPermissionsForProblem :one
SELECT
    ways.uuid as way_uuid,
    ways.name as way_name,
    EXISTS (
        SELECT 1
        FROM mentor_users_ways
        WHERE mentor_users_ways.way_uuid = ways.uuid
        AND mentor_users_ways.user_uuid = @user_uuid
    ) OR ways.owner_uuid = @user_uuid AS is_permission_given
FROM ways
INNER JOIN day_reports ON ways.uuid = day_reports.way_uuid
INNER JOIN problems ON problems.day_report_uuid = day_reports.uuid
WHERE problems.uuid = @problem_uuid;

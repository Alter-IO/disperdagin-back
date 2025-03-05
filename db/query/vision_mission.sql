-- name: FindVisionMissionByID :one
SELECT
    id,
    vision,
    mission,
    author,
    created_at,
    updated_at
FROM
    vision_mission
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindLatestVisionMission :one
SELECT
    id,
    vision,
    mission,
    author,
    created_at,
    updated_at
FROM
    vision_mission
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC
LIMIT 1;

-- name: FindAllVisionMissions :many
SELECT
    id,
    vision,
    mission,
    author,
    created_at
FROM
    vision_mission
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertVisionMission :exec
INSERT INTO vision_mission(id, vision, mission, author, created_at)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateVisionMission :execrows
UPDATE
    vision_mission
SET
    vision = sqlc.arg(vision),
    mission = sqlc.arg(mission),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteVisionMission :execrows
UPDATE
    vision_mission
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;
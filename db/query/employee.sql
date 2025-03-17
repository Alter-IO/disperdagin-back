-- name: FindEmployeeByID :one
SELECT
    id,
    name,
    position,
    address,
    employee_id,
    birthplace,
    birthdate,
    photo,
    status,
    author,
    created_at,
    updated_at
FROM
    employees
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllEmployees :many
SELECT
    id,
    name,
    position,
    employee_id,
    status,
    author,
    created_at
FROM
    employees
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindEmployeesByPosition :many
SELECT
    id,
    name,
    position,
    employee_id,
    status,
    author,
    created_at
FROM
    employees
WHERE
    position = sqlc.arg(position)
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindActiveEmployees :many
SELECT
    id,
    name,
    position,
    employee_id,
    author,
    created_at
FROM
    employees
WHERE
    status = 1
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertEmployee :exec
INSERT INTO employees(id, name, position, address, employee_id, birthplace, birthdate, photo, status, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(name),
    sqlc.arg(position),
    sqlc.arg(address),
    sqlc.arg(employee_id),
    sqlc.arg(birthplace),
    sqlc.arg(birthdate),
    sqlc.arg(photo),
    sqlc.arg(status),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateEmployee :execrows
UPDATE
    employees
SET
    name = sqlc.arg(name),
    position = sqlc.arg(position),
    address = sqlc.arg(address),
    employee_id = sqlc.arg(employee_id),
    birthplace = sqlc.arg(birthplace),
    birthdate = sqlc.arg(birthdate),
    photo = sqlc.arg(photo),
    status = sqlc.arg(status),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteEmployee :execrows
UPDATE
    employees
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;
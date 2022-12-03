-- name: GetTicket :one
SELECT * FROM ticket
WHERE username=$1 AND ticket_uid=$2 LIMIT 1;

-- name: ListTickets :many
SELECT * FROM ticket
WHERE username=$1;

-- name: CreateTicket :one
INSERT INTO ticket(ticket_uid,username,flight_number,price,status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM ticket
WHERE username=$1 AND ticket_uid=$2;
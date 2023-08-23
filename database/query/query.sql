-- name: GetAllBooks :one
select * from "Books"
where id = $1 LIMIT 1;

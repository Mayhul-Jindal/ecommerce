-- name: GetTagsByBookId :many
select tag_name from "Tags" as T
where T.id IN (
    select unnest(tags_array) from "Books" as B
    where B.id = $1
);

-- name: GetBooksByTags :many



-- name: GetBooksByFuzzy :many
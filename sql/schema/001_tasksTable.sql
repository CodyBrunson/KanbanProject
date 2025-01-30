-- +goose Up
CREATE TABLE tasks (
id UUID PRIMARY KEY,
title TEXT NOT NULL,
description TEXT NOT NULL,
status TEXT NOT NULL,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
completed_at TIMESTAMP,
deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE tasks;
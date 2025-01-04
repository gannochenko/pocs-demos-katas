CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users"
(
    "id"            uuid DEFAULT uuid_generate_v4(),
    "sup"           varchar(255) NOT NULL,
    "email"         varchar(255) NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT now(),
    "updated_at"    timestamptz NOT NULL DEFAULT now(),
    "deleted_at"    timestamptz,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX ix_users_sup ON "users" ("sup");

CREATE TABLE IF NOT EXISTS "images"
(
    "id"            uuid DEFAULT uuid_generate_v4(),
    "url"           varchar(255),
    "original_url"  varchar(255) NOT NULL,
    "created_by"    uuid NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT now(),
    "updated_at"    timestamptz NOT NULL DEFAULT now(),
    "deleted_at"    timestamptz,
    "is_processed"  bool NOT NULL DEFAULT FALSE,
    "is_failed"     bool NOT NULL DEFAULT FALSE,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("created_by") REFERENCES users("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS ix_images_created_by on images (created_by);

CREATE TABLE IF NOT EXISTS "image_processing_queues"
(
    "id"                uuid DEFAULT uuid_generate_v4(),
    "image_id"          uuid NOT NULL,
    "created_by"        uuid NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT now(),
    "updated_at"        timestamptz NOT NULL DEFAULT now(),
    "deleted_at"        timestamptz,
    "completed_at"      timestamptz,
    "operation_id"      uuid,
    "is_failed"         bool NOT NULL DEFAULT FALSE,
    "failure_reason"    text,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("created_by") REFERENCES users("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("image_id") REFERENCES images("id") ON DELETE CASCADE ON UPDATE CASCADE
);

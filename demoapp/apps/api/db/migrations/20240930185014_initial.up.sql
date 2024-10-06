CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "categories"
(
    "id"         uuid DEFAULT uuid_generate_v4(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "name"       text NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "tags"
(
    "id"         uuid DEFAULT uuid_generate_v4(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "name"       text NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "pets"
(
    "id"          uuid DEFAULT uuid_generate_v4(),
    "created_at"  timestamptz NOT NULL DEFAULT now(),
    "updated_at"  timestamptz NOT NULL DEFAULT now(),
    "deleted_at"  timestamptz,
    "name"        text NOT NULL,
    "status"      text NOT NULL,
    "photo_urls"  TEXT[],
    "category_id" uuid,
    PRIMARY KEY ("id"),
    CONSTRAINT fk_pets_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "pet_tags"
(
    "pet_id"         uuid,
    "tag_id"         uuid,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    PRIMARY KEY ("pet_id", "tag_id"),
    CONSTRAINT fk_pet_tags_pet FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    CONSTRAINT fk_pet_tags_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "customers"
(
    "id"         uuid DEFAULT uuid_generate_v4(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "username"   text NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "orders"
(
    "id"         uuid DEFAULT uuid_generate_v4(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "pet_id"     uuid NOT NULL,
    "quantity"   integer NOT NULL DEFAULT 1,
    "ship_date"  timestamptz,
    "status"     varchar(16) NOT NULL,
    "complete"   boolean NOT NULL DEFAULT FALSE,
    PRIMARY KEY ("id"),
    CONSTRAINT fk_pet FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "addresses"
(
    "id"         uuid DEFAULT uuid_generate_v4(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "street"     text NOT NULL,
    "city"       text NOT NULL,
    "state"      text NOT NULL,
    "zip"        text NOT NULL,
    PRIMARY KEY ("id")
);

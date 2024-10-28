CREATE TABLE IF NOT EXISTS items (
    "user_id"           INTEGER     NOT NULL,
    "name"              TEXT        NOT NULL,
    "quantity"          INTEGER     NOT NULL    DEFAULT(0),
    "reservation_id"    INTEGER     NULL,

    UNIQUE("user_id", "name")
)

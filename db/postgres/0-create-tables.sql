CREATE TABLE IF NOT EXISTS profile (
     id         uuid            PRIMARY KEY,
     username   VARCHAR (128)   UNIQUE NOT NULL,
     email      VARCHAR (128)   UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS image (
    id          uuid            PRIMARY KEY,
    path        VARCHAR (128)   UNIQUE NOT NULL,
    width       SMALLINT        NOT NULL,
    height      SMALLINT        NOT NULL
);

CREATE TABLE IF NOT EXISTS template (
    id          uuid            PRIMARY KEY,
    slug        VARCHAR (128)   UNIQUE NOT NULL,
    name        VARCHAR (128)   NOT NULL,
    nsfw        BOOLEAN         NOT NULL,
    created_on  TIMESTAMP       NOT NULL,
    user_id     uuid,
    image_id    uuid            NOT NULL,

    FOREIGN KEY (image_id) REFERENCES image (id),
    FOREIGN KEY (user_id) REFERENCES  profile (id)
);

CREATE TABLE IF NOT EXISTS template_default_text (
    template_id uuid            NOT NULL,
    index       SMALLINT        NOT NULL,
    text        VARCHAR (1024)  NOT NULL,

    FOREIGN KEY (template_id) REFERENCES template (id)
);

CREATE TABLE IF NOT EXISTS template_text_style (
    template_id         uuid            NOT NULL,
    index               SMALLINT        NOT NULL,
    x                   SMALLINT        NOT NULL,
    y                   SMALLINT        NOT NULL,
    width               SMALLINT        NOT NULL,
    font_family         VARCHAR (1024)  NOT NULL,
    font_size           SMALLINT        NOT NULL,
    font_color          VARCHAR(7)      NOT NULL,
    stroke_size         SMALLINT,
    stroke_color        VARCHAR(7),
    rotation_degrees    SMALLINT,

    FOREIGN KEY (template_id) REFERENCES template (id)
);

CREATE TABLE IF NOT EXISTS meme (
    id          uuid            PRIMARY KEY,
    path        VARCHAR (128)   UNIQUE NOT NULL,
    width       SMALLINT        NOT NULL,
    height      SMALLINT        NOT NULL,
    nsfw        BOOLEAN         NOT NULL,
    created_on  TIMESTAMP       NOT NULL,
    user_id     uuid,
    template_id uuid            NOT NULL,

    FOREIGN KEY (template_id) REFERENCES template (id),
    FOREIGN KEY (user_id) REFERENCES profile (id)
);

CREATE TABLE IF NOT EXISTS meme_text (
    meme_id     uuid            NOT NULL,
    index       SMALLINT        NOT NULL,
    text        VARCHAR (1024),

    FOREIGN KEY (meme_id) REFERENCES meme (id)
);
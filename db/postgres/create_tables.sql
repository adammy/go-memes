CREATE TABLE IF NOT EXISTS profile (
     id         uuid            PRIMARY KEY,
     username   VARCHAR (128)   UNIQUE NOT NULL,
     email      VARCHAR (128)   UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS base_image (
    id          uuid            PRIMARY KEY,
    img_path    VARCHAR (128)   UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS template (
    id          uuid            PRIMARY KEY,
    slug        VARCHAR (128)   UNIQUE NOT NULL,
    name        VARCHAR (128)   NOT NULL,
    width       SMALLINT        NOT NULL,
    height      SMALLINT        NOT NULL,
    created_on  TIMESTAMP       NOT NULL,
    user_id     uuid,
    image_id    uuid            NOT NULL,

    FOREIGN KEY (image_id) REFERENCES base_image (id),
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
    img_path    VARCHAR (128)   UNIQUE NOT NULL,
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

INSERT INTO profile (id, username, email) VALUES
('986a28ad-91d8-4454-8fae-1d1bfa2a14a9', 'adammy', 'adam@adammy.com'),
('d4e301d5-0cfa-4915-9e35-5084bc3a6b66', 'jesus', 'jesus@gmail.com');

INSERT INTO base_image (id, img_path) VALUES
('4cd223c2-6022-4bbf-aff1-e254548d5f71', 'assets/templates/yall-got-any-more-of-that.png'),
('41190721-022b-4cf5-ba45-b6f288ad78ce', 'assets/templates/two-buttons.png');

INSERT INTO template (id, slug, name, width, height, created_on, user_id, image_id) VALUES
('bc64af60-ef90-42c0-89cf-538297295432', 'yall-got-any-more-of-that', 'Y''all Got Any More Of That', 600, 471, NOW(), '986a28ad-91d8-4454-8fae-1d1bfa2a14a9', '4cd223c2-6022-4bbf-aff1-e254548d5f71'),
('5c71bce4-1b1d-4811-8f82-19744e769173', 'two-buttons', 'Two Buttons', 500, 756, NOW(), 'd4e301d5-0cfa-4915-9e35-5084bc3a6b66', '41190721-022b-4cf5-ba45-b6f288ad78ce');

INSERT INTO template_default_text (template_id, index, text) VALUES
('bc64af60-ef90-42c0-89cf-538297295432', 0, 'Y''all Got Any More of That'),
('bc64af60-ef90-42c0-89cf-538297295432', 1, 'Text'),
('5c71bce4-1b1d-4811-8f82-19744e769173', 0, 'Button Text'),
('5c71bce4-1b1d-4811-8f82-19744e769173', 1, 'Button Text'),
('5c71bce4-1b1d-4811-8f82-19744e769173', 2, 'Text Denoting Decider');

INSERT INTO template_text_style (template_id, index, x, y, width, font_family, font_size, font_color, stroke_size, stroke_color, rotation_degrees) VALUES
('bc64af60-ef90-42c0-89cf-538297295432', 0, 10, 10, 500, 'Impact', 40, '#FFFFFF', 4, '#000000', null),
('bc64af60-ef90-42c0-89cf-538297295432', 1, 10, 421, 580, 'Impact', 40, '#FFFFFF', 4, '#000000', null),
('5c71bce4-1b1d-4811-8f82-19744e769173', 0, 80, 110, 100, 'Arial', 20, '#000000', null, null, -10),
('5c71bce4-1b1d-4811-8f82-19744e769173', 1, 245, 80, 100, 'Arial', 20, '#000000', null, null, -10),
('5c71bce4-1b1d-4811-8f82-19744e769173', 2, 20, 675, 460, 'Impact', 40, '#FFFFFF', 4, '#000000', null);

INSERT INTO meme (id, img_path, created_on, user_id, template_id) VALUES
('7d161ffe-2a25-42d1-bb97-001065cc302f', 'assets/memes/7d161ffe-2a25-42d1-bb97-001065cc302f.png', NOW(), '986a28ad-91d8-4454-8fae-1d1bfa2a14a9', 'bc64af60-ef90-42c0-89cf-538297295432'),
('c41e3228-cdad-423b-9421-bad0a05739e0', 'assets/memes/c41e3228-cdad-423b-9421-bad0a05739e0.png', NOW(), 'd4e301d5-0cfa-4915-9e35-5084bc3a6b66', '5c71bce4-1b1d-4811-8f82-19744e769173');

INSERT INTO meme_text (meme_id, index, text) VALUES
('7d161ffe-2a25-42d1-bb97-001065cc302f', 0, 'Y''ALL GOT ANY MORE OF THEM'),
('7d161ffe-2a25-42d1-bb97-001065cc302f', 1, 'MONKEY JPEGs'),
('c41e3228-cdad-423b-9421-bad0a05739e0', 0, 'Well-Written Java'),
('c41e3228-cdad-423b-9421-bad0a05739e0', 1, 'Poorly-Written Kotlin'),
('c41e3228-cdad-423b-9421-bad0a05739e0', 2, 'JASON');

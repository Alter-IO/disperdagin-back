INSERT INTO roles(id, name, created_at)
VALUES ('superadmin', 'Superadmin', now()),
       ('admin', 'Admin', now()),
       ('guest', 'Guest', now());

INSERT INTO users(id, role_id, username, password, created_at)
VALUES ('01JPC83RZFN19CZRTRC6516XNM', 'superadmin', 'superadmin', '$2a$11$XewsGTuepOQl4VR2aU4fl.0VEJz.3BWn3v.GcTNfsGpleHyvQ7HEW', now()),
       ('01JPC878XKF1P3NJ8Q8ZZ5QHB2', 'admin', 'admin', '$2a$11$XewsGTuepOQl4VR2aU4fl.0VEJz.3BWn3v.GcTNfsGpleHyvQ7HEW', now()),
       ('01JPC879KBZEF9ES45616C5B05', 'guest', 'guest', '$2a$11$XewsGTuepOQl4VR2aU4fl.0VEJz.3BWn3v.GcTNfsGpleHyvQ7HEW', now());
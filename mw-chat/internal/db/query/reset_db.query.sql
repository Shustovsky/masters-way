-- name: RemoveEverything :exec
DO $$ DECLARE
    obj_name text;
    obj_type text;
    -- Exclude functions related to extension uuid-ossp
    exclude_functions text[] := ARRAY[
        'uuid_nil',
        'uuid_ns_dns',
        'uuid_ns_url',
        'uuid_ns_oid',
        'uuid_ns_x500',
        'uuid_generate_v1',
        'uuid_generate_v1mc',
        'uuid_generate_v3',
        'uuid_generate_v4',
        'uuid_generate_v5'
    ];
BEGIN
    -- Drop all tables
    FOR obj_name IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public' AND tablename != 'schema_migrations') LOOP
        EXECUTE 'DROP TABLE IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;

    -- Drop all sequences
    FOR obj_name IN (SELECT sequencename FROM pg_sequences WHERE schemaname = 'public') LOOP
        EXECUTE 'DROP SEQUENCE IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;

    -- Drop all views
    FOR obj_name IN (SELECT viewname FROM pg_views WHERE schemaname = 'public') LOOP
        EXECUTE 'DROP VIEW IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;

    -- Drop all materialized views
    FOR obj_name IN (SELECT matviewname FROM pg_matviews WHERE schemaname = 'public') LOOP
        EXECUTE 'DROP MATERIALIZED VIEW IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;

   -- Drop all functions except 'uuid_nil'
    FOR obj_name, obj_type IN
        SELECT proname, 'FUNCTION'
        FROM pg_proc
        JOIN pg_namespace ON pg_proc.pronamespace = pg_namespace.oid
        WHERE nspname = 'public' AND proname NOT IN (SELECT unnest(exclude_functions))
    LOOP
        EXECUTE 'DROP FUNCTION IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;

    -- Drop all types except 'schema_migrations' and 'schema_migrations[]'
    FOR obj_name IN (SELECT typname FROM pg_type WHERE typnamespace = 'public'::regnamespace AND typname NOT IN ('schema_migrations', 'schema_migrations[]')) LOOP
        BEGIN
            EXECUTE 'DROP TYPE IF EXISTS ' || obj_name || ' CASCADE;';
        EXCEPTION
            WHEN others THEN
                RAISE NOTICE 'Could not drop type %', obj_name;
        END;
    END LOOP;

    -- Drop all domains
    FOR obj_name IN (SELECT domain_name FROM information_schema.domains WHERE domain_schema = 'public') LOOP
        EXECUTE 'DROP DOMAIN IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;

    -- Drop all foreign tables
    FOR obj_name IN (SELECT foreign_table_name FROM information_schema.foreign_tables WHERE foreign_table_schema = 'public') LOOP
        EXECUTE 'DROP FOREIGN TABLE IF EXISTS ' || obj_name || ' CASCADE;';
    END LOOP;
END $$;


-- name: RegenerateDbData :exec
DO $$
BEGIN
INSERT INTO "rooms" ("uuid", "name", "type")
VALUES
    ('78bdf878-3b83-4f97-8d2e-928c132a10cd', NULL, 'private'),
    ('7c3a2511-c938-4a60-a9db-406e18bfeec0', NULL, 'private'),
    ('b7a3664c-f5ed-4cb0-aa2e-b2c758d22b55', NULL, 'private'),
    ('4f85694e-ff29-478f-90e9-1581577dfa84', NULL, 'private'),
    ('e57fc491-69f7-4b30-9979-78879c8873bf', NULL, 'private'),
    ('897f4a0f-fe31-4036-8358-f89a19c9bda6', NULL, 'private'),
    ('85f610df-9f86-4c55-8ee1-02485d42defb', NULL, 'private');

INSERT INTO "users_rooms" ("user_uuid", "room_uuid", "user_role", "is_room_blocked")
VALUES
    ('d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', '78bdf878-3b83-4f97-8d2e-928c132a10cd', 'regular', false),
    ('3d922e8a-5d58-4b82-9a3d-83e2e73b3f91', '78bdf878-3b83-4f97-8d2e-928c132a10cd', 'regular', false),

    ('d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', '7c3a2511-c938-4a60-a9db-406e18bfeec0', 'regular', false),
    ('1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1', '7c3a2511-c938-4a60-a9db-406e18bfeec0', 'regular', false),

    ('110f00b8-19e4-4cf4-a5f1-77b298bf0876', 'b7a3664c-f5ed-4cb0-aa2e-b2c758d22b55', 'regular', false),
    ('e93f8494-0c5c-420d-a68e-5d43903a80f2', 'b7a3664c-f5ed-4cb0-aa2e-b2c758d22b55', 'regular', false),

    ('b6eb9dd1-944c-4d1d-bc09-7c9933c53bab', '4f85694e-ff29-478f-90e9-1581577dfa84', 'regular', false),
    ('b51f6b20-9404-403b-8d48-1c0ab7d51340', '4f85694e-ff29-478f-90e9-1581577dfa84', 'regular', false),

    ('d63d2f89-6412-4324-8587-7061bf02dca4', 'e57fc491-69f7-4b30-9979-78879c8873bf', 'regular', false),
    ('c31384a6-b811-4a1f-befa-95dd53e3f4b9', 'e57fc491-69f7-4b30-9979-78879c8873bf', 'regular', false),

    ('d63d2f89-6412-4324-8587-7061bf02dca4', '897f4a0f-fe31-4036-8358-f89a19c9bda6', 'regular', false),
    ('5a31e3cb-7e9a-41e5-9a3b-1f1e5d6b7c3e', '897f4a0f-fe31-4036-8358-f89a19c9bda6', 'regular', false),

    ('d63d2f89-6412-4324-8587-7061bf02dca4', '85f610df-9f86-4c55-8ee1-02485d42defb', 'regular', false),
    ('8a3d1fe1-42da-499a-bf64-248297fd670a', '85f610df-9f86-4c55-8ee1-02485d42defb', 'regular', false);

INSERT INTO "messages" ("uuid", "owner_uuid", "room_uuid", "text", "created_at")
VALUES
    ('7939af01-e785-445d-b79d-70a433979c7b', 'd2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', '78bdf878-3b83-4f97-8d2e-928c132a10cd', 'Test message 1', '2024-08-09 01:00:00'),
    ('91be5d99-eddf-4949-bf15-b4cee3989fa6', '3d922e8a-5d58-4b82-9a3d-83e2e73b3f91', '78bdf878-3b83-4f97-8d2e-928c132a10cd', 'Test message 2', '2024-08-09 01:00:00'),

    ('88a6d503-a03b-412c-bfab-06649e49d62d', 'd2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', '78bdf878-3b83-4f97-8d2e-928c132a10cd', 'Test message 3', '2024-08-09 01:00:00'),
    ('6cea59ef-f0d4-4d8c-aa12-e48a746c93d0', '3d922e8a-5d58-4b82-9a3d-83e2e73b3f91', '78bdf878-3b83-4f97-8d2e-928c132a10cd', 'Test message 4', '2024-08-09 01:00:00'),

    ('aa9b5ca2-27af-494e-b1cb-20e9ec9d9ee8', 'd2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', '7c3a2511-c938-4a60-a9db-406e18bfeec0', 'Test message 5', '2024-08-09 01:00:00'),
    ('516394b3-0748-4e5e-9c5f-6d9e8e6ba0b1', '1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1', '7c3a2511-c938-4a60-a9db-406e18bfeec0', 'Test message 6', '2024-08-09 01:00:00'),

    ('3b5503c0-1ffa-4df7-8a3c-56535ce67422', '110f00b8-19e4-4cf4-a5f1-77b298bf0876', 'b7a3664c-f5ed-4cb0-aa2e-b2c758d22b55', 'Test message 7', '2024-08-09 01:00:00'),
    ('f6713833-4bc9-4a7a-b9e1-0bb797a52ef0', 'e93f8494-0c5c-420d-a68e-5d43903a80f2', 'b7a3664c-f5ed-4cb0-aa2e-b2c758d22b55', 'Test message 8', '2024-08-09 01:00:00'),

    ('8f7d6113-527f-44d9-a1db-53721653ba89', 'b6eb9dd1-944c-4d1d-bc09-7c9933c53bab', '4f85694e-ff29-478f-90e9-1581577dfa84', 'Test message 9', '2024-08-09 01:00:00'),
    ('b0532e85-84dc-4ab8-ae4d-b2c418c9f849', 'b51f6b20-9404-403b-8d48-1c0ab7d51340', '4f85694e-ff29-478f-90e9-1581577dfa84', 'Test message 10', '2024-08-09 01:00:00'),

    ('667f090a-0095-4884-bddc-c99c3fcc628d', 'd63d2f89-6412-4324-8587-7061bf02dca4', 'e57fc491-69f7-4b30-9979-78879c8873bf', 'Test message 11', '2024-08-09 01:00:00'),
    ('4f5d0404-d6bf-47a6-9f56-b7e3bb88a605', 'c31384a6-b811-4a1f-befa-95dd53e3f4b9', 'e57fc491-69f7-4b30-9979-78879c8873bf', 'Test message 12', '2024-08-09 01:00:00'),

    ('62a0c033-7206-4d21-9f7f-2246a4cdb0c2', 'd63d2f89-6412-4324-8587-7061bf02dca4', '897f4a0f-fe31-4036-8358-f89a19c9bda6', 'Test message 13', '2024-08-09 01:00:00'),
    ('dd02d5b3-5377-4f73-81a0-8119a208ea58', '5a31e3cb-7e9a-41e5-9a3b-1f1e5d6b7c3e', '897f4a0f-fe31-4036-8358-f89a19c9bda6', 'Test message 14', '2024-08-09 01:00:00'),

    ('b2ab09d5-7dcc-434d-9714-bb56cc71a2a8', 'd63d2f89-6412-4324-8587-7061bf02dca4', '85f610df-9f86-4c55-8ee1-02485d42defb', 'Test message 15', '2024-08-09 01:00:00'),
    ('c19faf46-6ddc-41ef-9260-16bf19372cc9', '8a3d1fe1-42da-499a-bf64-248297fd670a', '85f610df-9f86-4c55-8ee1-02485d42defb', 'Test message 16', '2024-08-09 01:00:00');

INSERT INTO "message_status" ("message_uuid", "receiver_uuid", "is_read", "updated_at")
VALUES
    ('7939af01-e785-445d-b79d-70a433979c7b', '3d922e8a-5d58-4b82-9a3d-83e2e73b3f91', false, '2024-08-09 01:05:00'),
    ('91be5d99-eddf-4949-bf15-b4cee3989fa6', 'd2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', false, '2024-08-09 01:05:00'),

    ('88a6d503-a03b-412c-bfab-06649e49d62d', '3d922e8a-5d58-4b82-9a3d-83e2e73b3f91', false, '2024-08-09 01:05:00'),
    ('6cea59ef-f0d4-4d8c-aa12-e48a746c93d0', 'd2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', false, '2024-08-09 01:05:00'),

    ('aa9b5ca2-27af-494e-b1cb-20e9ec9d9ee8', '1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1', false, '2024-08-09 01:05:00'),
    ('516394b3-0748-4e5e-9c5f-6d9e8e6ba0b1', 'd2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2', false, '2024-08-09 01:05:00'),

    ('3b5503c0-1ffa-4df7-8a3c-56535ce67422', 'e93f8494-0c5c-420d-a68e-5d43903a80f2', false, '2024-08-09 01:05:00'),
    ('f6713833-4bc9-4a7a-b9e1-0bb797a52ef0', '110f00b8-19e4-4cf4-a5f1-77b298bf0876', false, '2024-08-09 01:05:00'),

    ('8f7d6113-527f-44d9-a1db-53721653ba89', 'b51f6b20-9404-403b-8d48-1c0ab7d51340', false, '2024-08-09 01:05:00'),
    ('b0532e85-84dc-4ab8-ae4d-b2c418c9f849', 'b6eb9dd1-944c-4d1d-bc09-7c9933c53bab', false, '2024-08-09 01:05:00'),

    ('667f090a-0095-4884-bddc-c99c3fcc628d', 'c31384a6-b811-4a1f-befa-95dd53e3f4b9', false, '2024-08-09 01:05:00'),
    ('4f5d0404-d6bf-47a6-9f56-b7e3bb88a605', 'd63d2f89-6412-4324-8587-7061bf02dca4', false, '2024-08-09 01:05:00'),

    ('62a0c033-7206-4d21-9f7f-2246a4cdb0c2', '5a31e3cb-7e9a-41e5-9a3b-1f1e5d6b7c3e', false, '2024-08-09 01:05:00'),
    ('dd02d5b3-5377-4f73-81a0-8119a208ea58', 'd63d2f89-6412-4324-8587-7061bf02dca4', false, '2024-08-09 01:05:00'),

    ('b2ab09d5-7dcc-434d-9714-bb56cc71a2a8', '8a3d1fe1-42da-499a-bf64-248297fd670a', false, '2024-08-09 01:05:00'),
    ('c19faf46-6ddc-41ef-9260-16bf19372cc9', 'd63d2f89-6412-4324-8587-7061bf02dca4', false, '2024-08-09 01:05:00');

END $$;

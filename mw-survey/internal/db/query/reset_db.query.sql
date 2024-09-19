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

END $$;

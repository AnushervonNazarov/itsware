CREATE TABLE device_profiles (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL,
    description TEXT,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE device_profile_type AS (
    id INT,
    name CITEXT,
    description TEXT,
    tenant_id INT,
    created_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION create_device_profile(dprofile device_profile_type, user_id INT) 
RETURNS device_profile_type AS $$
DECLARE
    new_device_profile device_profile_type;
BEGIN
    IF check_permission(user_id) THEN
        INSERT INTO device_profiles (name, description, tenant_id)
        VALUES (dprofile.name, dprofile.description, dprofile.tenant_id)
        RETURNING device_profiles.id, device_profiles.name, device_profiles.description, device_profiles.tenant_id, device_profiles.created_at
        INTO new_device_profile;

        RETURN new_device_profile;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Get Device Profiles By ID

CREATE OR REPLACE FUNCTION get_device_profile_by_id(p_id int)
RETURNS device_profile_type AS $$
DECLARE
    new_device_profile device_profile_type;
BEGIN
    SELECT id, name, description, tenant_id, created_at
    INTO new_device_profile
    FROM device_profiles
    WHERE id = p_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_device_profile;
END;
$$ LANGUAGE plpgsql;

-- Get Device Profiles

CREATE OR REPLACE FUNCTION get_device_profiles(user_id INT)
RETURNS SETOF device_profile_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY
        SELECT id, name, description, tenant_id, created_at
          FROM device_profiles;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Device Profile

CREATE OR REPLACE FUNCTION update_device_profile(dprofile device_profile_type, user_id INT)
RETURNS device_profile_type AS $$
DECLARE
    new_device_profile device_profile_type;
BEGIN
    IF check_permission(user_id) THEN
        UPDATE device_profiles
           SET name = dprofile.name,
               description = dprofile.description,
               tenant_id = dprofile.tenant_id
        WHERE id = dprofile.id
        RETURNING device_profiles.id, device_profiles.name, device_profiles.description, device_profiles.tenant_id, device_profiles.created_at
        INTO new_device_profile;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

        RETURN new_device_profile;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Delete Device Profile

CREATE OR REPLACE FUNCTION delete_device_profile(p_id INT, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM device_profiles WHERE id = p_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
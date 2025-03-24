-- Active: 1739515049434@@localhost@5432@postgres
-- Log: Script execution started

CREATE SCHEMA tenants;
CREATE SCHEMA users;
CREATE SCHEMA teams;
CREATE SCHEMA cabinets;
CREATE SCHEMA devices;
CREATE SCHEMA device_profiles;
CREATE SCHEMA roles;

-------------------------------
-- Create Tables
-------------------------------
CREATE TABLE tenants (
    id SERIAL PRIMARY KEY,
    name CITEXT,
    is_support BOOLEAN,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_modified_on TIMESTAMP
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name CITEXT
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name CITEXT NOT NULL,
    last_name CITEXT NOT NULL,
    email CITEXT UNIQUE NOT NULL,
    phone TEXT NOT NULL,
    password CITEXT NOT NULL,
    role_id INT REFERENCES roles(id) ON DELETE SET NULL,
    tenant_id INT REFERENCES tenants(id),
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id)
);
CREATE INDEX idx_users_tenant_id ON users(tenant_id);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_teams_tenant_id ON teams(tenant_id);

CREATE TABLE device_statuses (
    id SERIAL PRIMARY KEY,
    name CITEXT
);


CREATE TABLE cabinets (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL UNIQUE,
    location CITEXT,
    description TEXT,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_cabinets_tenant_id ON cabinets(tenant_id);


CREATE TABLE device_profiles (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL,
    description TEXT,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_device_profiles_tenant_id ON device_profiles(tenant_id);

CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL,
    description TEXT,
    device_status_id INT REFERENCES device_statuses(id) ON DELETE SET NULL,
    serial_number TEXT UNIQUE NOT NULL,
    checked_out_on TIMESTAMP,
    checked_out_by INT REFERENCES users(id) ON DELETE SET NULL,
    cabinet_id INT REFERENCES cabinets(id) ON DELETE SET NULL,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    device_profile_id INT REFERENCES device_profiles(id) ON DELETE SET NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_devices_cabinet_id ON devices(cabinet_id);
CREATE INDEX idx_devices_tenant_id ON devices(tenant_id);
CREATE INDEX idx_devices_device_profile_id ON devices(device_profile_id);


CREATE TABLE team_users (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    created_on TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL,
    PRIMARY KEY (team_id, user_id)
);


CREATE TABLE team_cabinets (
    cabinet_id INT NOT NULL REFERENCES cabinets(id) ON DELETE CASCADE,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    created_on TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL,
    PRIMARY KEY (team_id, cabinet_id)
);

CREATE TABLE cabinet_user_access ( 
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    cabinet_id INT NOT NULL REFERENCES cabinets(id) ON DELETE CASCADE,
    created_on TIMESTAMP,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    last_modified_on TIMESTAMP,
    last_modified_by INT REFERENCES users(id) ON DELETE SET NULL,
    PRIMARY KEY (user_id, cabinet_id)
);

-------------------------------
-- Drop Types if They Exist
-------------------------------
DROP TYPE IF EXISTS device_profiles.device_profile CASCADE;
DROP TYPE IF EXISTS devices.device CASCADE;
DROP TYPE IF EXISTS teams.team CASCADE;
DROP TYPE IF EXISTS tenants.tenant CASCADE;
DROP TYPE IF EXISTS cabinets.cabinet CASCADE;
DROP TYPE IF EXISTS users."user" CASCADE;
DROP TYPE IF EXISTS team_cabinet CASCADE;
DROP TYPE IF EXISTS cabinet_user_access CASCADE;
DROP TYPE IF EXISTS team_user CASCADE;


-------------------------------
-- Create Types
-------------------------------

CREATE TYPE tenants.tenant AS (
    id INT,
    name CITEXT,
    is_support BOOLEAN,
    created_on TIMESTAMP,
    last_modified_on TIMESTAMP
);

CREATE TYPE users."user" AS (
    id INT,
    first_name CITEXT,
    last_name CITEXT,
    email CITEXT,
    phone TEXT,
    password CITEXT,
    role_id INT,
    tenant_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE teams.team AS (
    id INT,
    name CITEXT,
    tenant_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE team_user AS (
    user_id INT,
    team_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE cabinets.cabinet AS (
    id INT,
    name CITEXT,
    location CITEXT,
    description TEXT,
    tenant_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE team_cabinet AS (
    cabinet_id INT,
    team_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE cabinet_user_access AS (
    user_id INT,
    cabinet_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE device_profiles.device_profile AS (
    id INT,
    name CITEXT,
    description TEXT,
    tenant_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);

CREATE TYPE devices.device AS (
    id INT,
    name CITEXT,
    description TEXT,
    device_status_id INT,
    serial_number TEXT,
    checked_out_on TIMESTAMP,
    checked_out_by INT,
    cabinet_id INT,
    tenant_id INT,
    device_profile_id INT,
    created_on TIMESTAMP,
    created_by INT,
    last_modified_on TIMESTAMP,
    last_modified_by INT
);
-------------------------------
-- CRUD Functions Using Composite Types
-------------------------------

CREATE OR REPLACE FUNCTION check_support_admin()
RETURNS VOID AS $$
DECLARE
    s_user_role TEXT;
BEGIN
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role != 'support_admin' THEN
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION permission_denied()
RETURNS VOID AS $$
BEGIN
    RAISE EXCEPTION 'Permission denied';
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION has_support_read_access()
RETURNS BOOLEAN AS $$
DECLARE
    s_user_role TEXT;
BEGIN
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role NOT IN ('support_admin', 'support_view') THEN
        RETURN FALSE;
    ELSE
        RETURN TRUE;
    END IF;
END;
$$ LANGUAGE plpgsql;  

-- Create Tenant
CREATE OR REPLACE FUNCTION tenants.create(p_tenant tenants.tenant)
RETURNS tenants.tenant AS $$
DECLARE
    new_tenant tenants.tenant;
BEGIN
    
    PERFORM check_support_admin();

    INSERT INTO tenants (name, is_support, created_on)
    VALUES (p_tenant.name, p_tenant.is_support, NOW())
    RETURNING id, name, is_support, created_on, last_modified_on
    INTO new_tenant;

    RETURN new_tenant;
END;
$$ LANGUAGE plpgsql;


-- Get Tenant By ID

CREATE OR REPLACE FUNCTION tenants.get_one(p_tenant_id INT)
RETURNS tenants.tenant AS $$
DECLARE
    new_tenant tenants.tenant;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT id, name, is_support, created_on, last_modified_on
    INTO new_tenant
    FROM tenants
    WHERE id = p_tenant_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    IF NOT (has_support_read_access() OR (s_user_role = 'administrator' AND s_tenant_id = new_tenant.id)) THEN
        PERFORM permission_denied();
    END IF;
    RETURN new_tenant;
END;
$$ LANGUAGE plpgsql;

-- Get Tenants

CREATE OR REPLACE FUNCTION tenants.get_all()
RETURNS SETOF tenants.tenant AS $$
DECLARE
    s_user_id INT;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF has_support_read_access() THEN
        RETURN QUERY
        SELECT id, name, is_support, created_on, last_modified_on
        FROM tenants;
    ELSEIF s_user_role = 'administrator' THEN
        RETURN QUERY
        SELECT id, name, is_support, created_on, last_modified_on
        FROM tenants
        WHERE tenant_id = s_tenant_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Tenant

CREATE OR REPLACE FUNCTION tenants.update(p_tenant tenants.tenant)
RETURNS tenants.tenant AS $$
DECLARE
    new_tenant tenants.tenant;
    s_tenant_id INT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;

    PERFORM check_support_admin();

    UPDATE tenants
        SET name = p_tenant.name,
            is_support = p_tenant.is_support,
            last_modified_on = NOW()
        WHERE id = p_tenant.id
    RETURNING id, name, is_support, created_on, last_modified_on
    INTO new_tenant;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_tenant;
END;
$$ LANGUAGE plpgsql;

-- Delete Tenant

CREATE OR REPLACE FUNCTION tenants.delete(p_tenant_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;

    PERFORM check_support_admin();

    DELETE FROM tenants WHERE id = p_tenant_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

END;
$$ LANGUAGE plpgsql;

-- Create User

CREATE OR REPLACE FUNCTION users.create(p_user_input users."user") 
RETURNS users."user" AS $$
DECLARE
    new_user users."user";
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_user_input.tenant_id = s_tenant_id) THEN
        INSERT INTO users (first_name, last_name, password, email, phone, role_id, tenant_id, created_on, created_by)
        VALUES (p_user_input.first_name, p_user_input.last_name, p_user_input.password, p_user_input.email, p_user_input.phone, p_user_input.role_id, p_user_input.tenant_id, NOW(), s_user_id)
        RETURNING id, first_name, last_name, password, email, phone, role_id, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        INTO new_user;
    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_user;
END;
$$ LANGUAGE plpgsql;

-- Get User By ID

CREATE OR REPLACE FUNCTION users.get_one(p_user_id INT)
RETURNS users."user" AS $$
DECLARE
    new_user users."user";
    s_user_id INT;
    s_user_role TEXT;
    s_tenant_id INT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT id, first_name, last_name, password, email, phone, role_id, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_user
    FROM users
    WHERE id = p_user_id;   

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    IF NOT (s_user_id = new_user.id OR has_support_read_access() 
    OR (s_user_role = 'administrator' AND s_tenant_id = new_user.tenant_id)) THEN
        PERFORM permission_denied();
    END IF;

    RETURN new_user;
END;
$$ LANGUAGE plpgsql;

-- Get Users

CREATE OR REPLACE FUNCTION users.get_all()
RETURNS SETOF users."user" AS $$
DECLARE
    s_user_id INT;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF has_support_read_access() THEN
        RETURN QUERY 
        SELECT id, first_name, last_name, email, phone, password, role_id, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM users;
    ELSEIF s_user_role = 'administrator' THEN
        RETURN QUERY
        SELECT id, first_name, last_name, email, phone, password, role_id, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM users
        WHERE tenant_id = s_tenant_id;
    ELSE
        RETURN QUERY
        SELECT id, first_name, last_name, email, phone, password, role_id, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM users
        WHERE id = s_user_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update User

CREATE OR REPLACE FUNCTION users.update(p_user_input users."user")
RETURNS users."user" AS $$
DECLARE
    new_user users."user";
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_user_input.tenant_id = s_tenant_id) THEN
    UPDATE users
       SET first_name = p_user_input.first_name,
           last_name = p_user_input.last_name,
           email = p_user_input.email,
           phone = p_user_input.phone,
           last_modified_on = NOW(),
           last_modified_by = s_user_id
    WHERE id = p_user_input.id
    RETURNING id, first_name, last_name, email, phone, password, role_id, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_user;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_user;
END;
$$ LANGUAGE plpgsql;

-- Delete User

CREATE OR REPLACE FUNCTION users.delete(p_user_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' THEN
        DELETE FROM users WHERE id = p_user_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSIF s_user_role = 'administrator' THEN
        DELETE FROM users WHERE id = p_user_id AND tenant_id = s_tenant_id;        
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;
    
    ELSE
        PERFORM permission_denied();
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Add User to Team

CREATE OR REPLACE FUNCTION add_user_to_team(p_user team_user)
RETURNS team_user AS $$
DECLARE 
    new_team_users team_user;
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
    user_tenant_id INT;
    team_tenant_id INT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);
    
    SELECT tenant_id INTO user_tenant_id FROM users WHERE id = p_user.user_id;
    SELECT tenant_id INTO team_tenant_id FROM teams WHERE id = p_user.user_id;

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND user_tenant_id = s_tenant_id AND team_tenant_id = s_tenant_id) THEN
    INSERT INTO team_users (user_id, team_id, created_on, created_by)
    VALUES (p_user.user_id, p_user.team_id, NOW(), s_user_id)
    ON CONFLICT DO NOTHING
    RETURNING user_id, team_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_team_users;

    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_team_users;
END;
$$ LANGUAGE plpgsql;

-- Remove User from Team

CREATE OR REPLACE FUNCTION remove_user_from_team(p_user_id INT, p_team_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
    user_tenant_id INT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT tenant_id INTO user_tenant_id FROM users WHERE id = s_tenant_id;

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND user_tenant_id = s_tenant_id) THEN
    DELETE FROM team_users WHERE team_users.team_id = p_team_id AND team_users.user_id = p_user_id;
    
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record does not exist';
    END IF;

    ELSE
        PERFORM permission_denied();
    END IF;

END;
$$ LANGUAGE plpgsql;

-- Create Team

CREATE OR REPLACE FUNCTION teams.create(p_team teams.team) 
RETURNS teams.team AS $$
DECLARE
    new_team teams.team;
    s_user_role TEXT;
    s_tenant_id INT;
    s_user_id INT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_team.tenant_id = s_tenant_id) THEN
    INSERT INTO teams (name, tenant_id, created_on, created_by)
    VALUES (p_team.name, p_team.tenant_id, NOW(), s_user_id)
    RETURNING id, name, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_team;

    ELSE
        PERFORM permission_denied();
    END IF;
    
    RETURN new_team;

END;
$$ LANGUAGE plpgsql;

-- Get Team By ID

CREATE OR REPLACE FUNCTION teams.get_one(p_team_id INT)
RETURNS teams.team AS $$
DECLARE
    new_team teams.team;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT id, name, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_team
    FROM teams
    WHERE id = p_team_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    IF NOT (has_support_read_access() OR (s_user_role = 'administrator' AND s_tenant_id = new_team.tenant_id)) THEN
        PERFORM permission_denied();
    END IF;

    RETURN new_team;
END;
$$ LANGUAGE plpgsql;

-- Get Teams

CREATE OR REPLACE FUNCTION teams.get_all()
RETURNS SETOF teams.team AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF has_support_read_access() THEN
        RETURN QUERY 
        SELECT id, name, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM teams;
    ELSEIF s_user_role = 'administrator' THEN
        RETURN QUERY
        SELECT id, name, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM teams
        WHERE tenant_id = s_tenant_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Teams 

CREATE OR REPLACE FUNCTION teams.update(p_team teams.team)
RETURNS teams.team AS $$
DECLARE
    new_team teams.team;
    s_user_id INT;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_team.tenant_id = s_tenant_id) THEN
    UPDATE teams
       SET name = p_team.name,
           last_modified_on = NOW(),
           last_modified_by = s_user_id
    WHERE id = p_team.id
    RETURNING id, name, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_team;
  
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_team;
END;
$$ LANGUAGE plpgsql;

-- Delete Team 

CREATE OR REPLACE FUNCTION teams.delete(p_team_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' THEN
        DELETE FROM teams WHERE id = p_team_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;
    ELSIF s_user_role = 'administrator' THEN
        DELETE FROM teams WHERE id = p_team_id AND tenant_id = s_tenant_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;
    ELSE
        PERFORM permission_denied();
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Create Cabinet

CREATE OR REPLACE FUNCTION cabinets.create(p_cabinet cabinets.cabinet)
RETURNS cabinets.cabinet AS $$
DECLARE
    new_cabinet cabinets.cabinet;
    s_user_id INT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;

    PERFORM check_support_admin();

    INSERT INTO cabinets (name, location, description, tenant_id, created_on, created_by)
    VALUES (p_cabinet.name, p_cabinet.location, p_cabinet.description, p_cabinet.tenant_id, NOW(), s_user_id)
    RETURNING id, name, location, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_cabinet;

    RETURN new_cabinet;

END;
$$ LANGUAGE plpgsql;

-- Get Cabinet By ID

CREATE OR REPLACE FUNCTION cabinets.get_one(p_cabinet_id INT)
RETURNS cabinets.cabinet AS $$
DECLARE
    new_cabinet cabinets.cabinet;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT id, name, location, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_cabinet
    FROM cabinets
    WHERE id = p_cabinet_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    IF NOT (has_support_read_access() OR (s_user_role = 'administrator' AND s_tenant_id = new_cabinet.tenant_id)) THEN
        PERFORM permission_denied();
    END IF;

    RETURN new_cabinet;
END;
$$ LANGUAGE plpgsql;

-- Get Cabinets

CREATE OR REPLACE FUNCTION cabinets.get_all()
RETURNS SETOF cabinets.cabinet AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);
        
    IF has_support_read_access() THEN
        RETURN QUERY
        SELECT id, name, location, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM cabinets;
    ELSEIF s_user_role = 'administrator' THEN
        RETURN QUERY
        SELECT id, name, location, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM cabinets
        WHERE tenant_id = s_tenant_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Cabinet

CREATE OR REPLACE FUNCTION cabinets.update(p_cabinet cabinets.cabinet)
RETURNS cabinets.cabinet AS $$
DECLARE
    new_cabinet cabinets.cabinet;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    PERFORM check_support_admin();
    
    UPDATE cabinets
        SET name = p_cabinet.name, 
            location = p_cabinet.location, 
            description = p_cabinet.description,
            last_modified_on = NOW(),
            last_modified_by = s_user_id
    WHERE cabinets.id = p_cabinet.id
    RETURNING id, name, location, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_cabinet;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_cabinet;

END;
$$ LANGUAGE plpgsql;

-- Delete Cabinet

CREATE OR REPLACE FUNCTION cabinets.delete(p_cabinet_id INT) 
RETURNS VOID AS $$
BEGIN

    PERFORM check_support_admin();

    DELETE FROM cabinets WHERE id = p_cabinet_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Add Cabinet to Team

CREATE OR REPLACE FUNCTION add_cabinet_to_team(p_cabinet team_cabinet)
RETURNS team_cabinet AS $$
DECLARE
    new_team_cabinet team_cabinet;
    s_tenant_id INT;
    s_user_role TEXT;
    s_user_id INT;
    cabinet_tenant_id INT;
    team_tenant_id INT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);
    
    SELECT tenant_id INTO cabinet_tenant_id FROM cabinets WHERE id = s_tenant_id;
    SELECT tenant_id INTO team_tenant_id FROM teams WHERE id = s_tenant_id;

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND cabinet_tenant_id = s_tenant_id AND team_tenant_id = s_tenant_id) THEN
        INSERT INTO team_cabinets (cabinet_id, team_id, created_on, created_by)
        VALUES (p_cabinet.cabinet_id, p_cabinet.team_id, NOW(), s_user_id)
        ON CONFLICT DO NOTHING
        RETURNING cabinet_id, team_id, created_on, created_by, last_modified_on, last_modified_by
        INTO new_team_cabinet;
    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_team_cabinet;
END;
$$ LANGUAGE plpgsql;

-- Remove Cabinet from Team

CREATE OR REPLACE FUNCTION remove_cabinet_from_team(p_cabinet_id INT, p_team_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
    cabinet_tenant_id INT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT tenant_id INTO cabinet_tenant_id FROM cabinets WHERE id = s_tenant_id;

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND cabinet_tenant_id = s_tenant_id) THEN
    DELETE FROM team_cabinets WHERE team_cabinets.cabinet_id = p_cabinet_id AND team_cabinets.team_id = p_team_id;
    
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record does not exist';
    END IF;

    ELSE
        PERFORM permission_denied();
    END IF;

END;
$$ LANGUAGE plpgsql;

SELECT remove_cabinet_from_team(1,2);

CREATE OR REPLACE FUNCTION cabinet_user_access_add(input cabinet_user_access)
RETURNS cabinet_user_access AS $$
DECLARE
    new_access cabinet_user_access;
BEGIN

    PERFORM check_support_admin();

    INSERT INTO cabinet_user_access (user_id, cabinet_id, created_on, created_by)
    VALUES (input.user_id, input.cabinet_id, NOW(), input.created_by)
    --ON CONFLICT DO NOTHING
    RETURNING user_id, cabinet_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_access;

    RETURN new_access;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION cabinet_user_access_get_one(p_user_id INT, p_cabinet_id INT)
RETURNS cabinet_user_access AS $$
DECLARE
    access cabinet_user_access;
    s_user_role TEXT;
BEGIN
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT user_id, cabinet_id, created_on, created_by, last_modified_on, last_modified_by
    INTO access
    FROM cabinet_user_access
    WHERE user_id = p_user_id AND cabinet_id = p_cabinet_id AND s_user_role IN ('administrator', 'support_admin', 'support_view');

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN access;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION cabinet_user_access_get_all()
RETURNS SETOF cabinet_user_access AS $$
DECLARE
    s_user_role TEXT;
BEGIN
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    RETURN QUERY
    SELECT user_id, cabinet_id, created_on, created_by, last_modified_on, last_modified_by
    FROM cabinet_user_access
    WHERE s_user_role IN ('administrator', 'support_admin', 'support_view');
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION cabinet_user_access_remove(p_user_id INT, p_cabinet_id INT)
RETURNS VOID AS $$
BEGIN

    PERFORM check_support_admin();

    DELETE FROM cabinet_user_access WHERE user_id = p_user_id AND cabinet_id = p_cabinet_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Create Device Profile

CREATE OR REPLACE FUNCTION device_profiles.create(p_device_profile device_profiles.device_profile) 
RETURNS device_profiles.device_profile AS $$
DECLARE
    new_device_profile device_profiles.device_profile;
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);
    
    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_device_profile.tenant_id = s_tenant_id) THEN
        INSERT INTO device_profiles (name, description, tenant_id, created_on, created_by)
        VALUES (p_device_profile.name, p_device_profile.description, p_device_profile.tenant_id, NOW(), s_user_id)
        RETURNING id, name, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        INTO new_device_profile;

    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_device_profile;

END;
$$ LANGUAGE plpgsql;

-- Get Device Profiles By ID

CREATE OR REPLACE FUNCTION device_profiles.get_one(p_device_profile_id INT)
RETURNS device_profiles.device_profile AS $$
DECLARE
    new_device_profile device_profiles.device_profile;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT id, name, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_device_profile
    FROM device_profiles
    WHERE id = p_device_profile_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    IF NOT (has_support_read_access() OR (s_user_role = 'administrator' AND s_tenant_id = new_device_profile.tenant_id)) THEN
        PERFORM permission_denied();
    END IF;

    RETURN new_device_profile;
END;
$$ LANGUAGE plpgsql;

-- Get Device Profiles

CREATE OR REPLACE FUNCTION device_profiles.get_all()
RETURNS SETOF device_profiles.device_profile AS $$
DECLARE
    s_user_id INT;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);
    
    IF has_support_read_access() THEN
        RETURN QUERY
        SELECT id, name, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM device_profiles;
    ELSEIF s_user_role = 'administrator' THEN
        RETURN QUERY
        SELECT id, name, description, tenant_id, created_on, created_by, last_modified_on, last_modified_by
        FROM device_profiles
        WHERE tenant_id = s_tenant_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Device Profile

CREATE OR REPLACE FUNCTION device_profiles.update(p_device_profile device_profiles.device_profile)
RETURNS device_profiles.device_profile AS $$
DECLARE
    new_device_profile device_profiles.device_profile;
    s_user_id INT;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_device_profile.tenant_id = s_tenant_id) THEN
        UPDATE device_profiles
           SET name = p_device_profile.name,
               description = p_device_profile.description,
               last_modified_on = NOW(),
               last_modified_by = s_user_id
        WHERE id = p_device_profile.id
        RETURNING device_profiles.id, device_profiles.name, device_profiles.description, device_profiles.tenant_id, device_profiles.created_on, device_profiles.created_by, device_profiles.last_modified_on, device_profiles.last_modified_by
        INTO new_device_profile;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_device_profile;

END;
$$ LANGUAGE plpgsql;

-- Delete Device Profile

CREATE OR REPLACE FUNCTION device_profiles.delete(p_device_profile_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' THEN
        DELETE FROM device_profiles WHERE id = p_device_profile_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;
    ELSIF s_user_role = 'administrator' THEN
        DELETE FROM device_profiles WHERE id = p_device_profile_id AND tenant_id = s_tenant_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;
    ELSE
        PERFORM permission_denied();
    END IF;

END;
$$ LANGUAGE plpgsql;

-- Create Device

CREATE OR REPLACE FUNCTION devices.create(p_device devices.device)
RETURNS devices.device AS $$
DECLARE
    new_device devices.device;
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_device.tenant_id = s_tenant_id) THEN
        INSERT INTO devices (name, description, device_status_id, serial_number, cabinet_id, tenant_id, device_profile_id, created_on, created_by)
        VALUES (p_device.name, p_device.description, p_device.device_status_id, p_device.serial_number, p_device.cabinet_id, p_device.tenant_id, p_device.device_profile_id, NOW(), s_user_id)
        RETURNING id ,name, description, device_status_id, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_on, created_by, last_modified_on, last_modified_by
        INTO new_device;

    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_device;

END;
$$ LANGUAGE plpgsql;

-- Get Device By ID


CREATE OR REPLACE FUNCTION devices.get_one(p_device_id INT)
RETURNS devices.device AS $$
DECLARE
    new_device devices.device;
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    SELECT id, name, description, device_status_id, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_on, created_by, last_modified_on, last_modified_by
    INTO new_device
    FROM devices
    WHERE id = p_device_id;

    UPDATE devices
        SET checked_out_by = s_user_id,
            checked_out_on = NOW()
        WHERE id = p_device_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    IF NOT (has_support_read_access() OR (s_user_role = 'administrator' AND s_tenant_id = new_device.tenant_id)) THEN
        PERFORM permission_denied();
    END IF;

    RETURN new_device;
END;
$$ LANGUAGE plpgsql;

-- Get Devices

CREATE OR REPLACE FUNCTION devices.get_all()
RETURNS SETOF devices.device AS $$
DECLARE
    s_tenant_id INT;
    s_user_id INT;
    s_user_role TEXT;
BEGIN
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF has_support_read_access() THEN
        RETURN QUERY 
        SELECT id, name, description, device_status_id, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_on, created_by, last_modified_on, last_modified_by
        FROM devices;

        UPDATE devices
        SET checked_out_by = s_user_id,
            checked_out_on = NOW();
    ELSEIF s_user_role = 'administrator' THEN
        RETURN QUERY 
        SELECT id, name, description, device_status_id, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_on, created_by, last_modified_on, last_modified_by
        FROM devices
        WHERE tenant_id = s_tenant_id;
    
        UPDATE devices
        SET checked_out_by = s_user_id,
            checked_out_on = NOW();
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Device

CREATE FUNCTION devices.update(p_device devices.device)
RETURNS devices.device AS $$
DECLARE
    new_device devices.device;
    s_user_id INT;
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN 
    s_user_id := current_setting('myapp.session.user_id', TRUE)::INT;
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' OR (s_user_role = 'administrator' AND p_device.tenant_id = s_tenant_id) THEN
        UPDATE devices
        SET name = p_device.name,
            description = p_device.description,
            device_status_id = p_device.device_status_id,
            serial_number = p_device.serial_number,
            cabinet_id = p_device.cabinet_id,
            device_profile_id = p_device.device_profile_id,
            last_modified_on = NOW(),
            last_modified_by = s_user_id
        WHERE id = p_device.id
        RETURNING id, name, description, device_status_id, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_on, created_by, last_modified_on, last_modified_by
        INTO new_device;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;  
    
    ELSE
        PERFORM permission_denied();
    END IF;

    RETURN new_device;
END;
$$ LANGUAGE plpgsql;

-- Delete Device

CREATE OR REPLACE FUNCTION devices.delete(p_device_id INT)
RETURNS VOID AS $$
DECLARE
    s_tenant_id INT;
    s_user_role TEXT;
BEGIN
    s_tenant_id := current_setting('myapp.session.tenant_id', TRUE)::INT;
    s_user_role := current_setting('myapp.session.user_role', TRUE);

    IF s_user_role = 'support_admin' THEN
        DELETE FROM devices WHERE id = p_device_id;
        IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;
    ELSIF s_user_role = 'administrator' THEN
        DELETE FROM devices WHERE id = p_device_id AND tenant_id = s_tenant_id;
        IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;
    ELSE
        PERFORM permission_denied();
    END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TABLE cabinets (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL UNIQUE,
    location CITEXT,
    description TEXT,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cabinets_team (
    team_id INT REFERENCES teams(id) ON DELETE SET NULL,
    cabinet_id INT REFERENCES cabinets(id) ON DELETE SET NULL,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP,
    PRIMARY KEY (team_id, cabinet_id)
);

CREATE TYPE cabinet_type AS (
    id INT,
    name CITEXT,
    location CITEXT,
    description TEXT,
    tenant_id INT,
    created_by INT,
    created_at TIMESTAMP
);

CREATE TYPE cabinets_team_type AS (
    team_id INT,
    cabinet_id INT,
    created_by INT,
    created_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION create_cabinet(cabinet cabinet_type, user_id INT)
RETURNS cabinet_type AS $$
DECLARE
    new_cabinet cabinet_type;
BEGIN
    IF check_permission(user_id) THEN
        INSERT INTO cabinets (name, location, description, tenant_id, created_by)
        VALUES (cabinet.name, cabinet.location, cabinet.description, cabinet.tenant_id, cabinet.created_by)
        RETURNING cabinets.id, cabinets.name, cabinets.location, cabinets.description, cabinets.tenant_id, cabinets.created_by, cabinets.created_at
        INTO new_cabinet;

        RETURN new_cabinet;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Get Cabinet By ID

CREATE OR REPLACE FUNCTION get_cabinet_by_id(c_id INT)
RETURNS cabinet_type AS $$
DECLARE
    new_cabinet cabinet_type;
BEGIN
    SELECT id, name, location, description, tenant_id, created_by, created_at
    INTO new_cabinet
    FROM cabinets
    WHERE id = c_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_cabinet;
END;
$$ LANGUAGE plpgsql;

-- Get Cabinets

CREATE OR REPLACE FUNCTION get_cabinets(user_id INT)
RETURNS SETOF cabinet_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY
            SELECT id, name, location, description, tenant_id, created_by, created_at
            FROM cabinets;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Cabinet

CREATE OR REPLACE FUNCTION update_cabinet(cabinet cabinet_type, user_id INT)
RETURNS cabinet_type AS $$
DECLARE
    new_cabinet cabinet_type;
BEGIN
    IF check_permission(user_id) THEN
        UPDATE cabinets
        SET 
            name = cabinet.name, 
            location = cabinet.location, 
            description = cabinet.description,
            tenant_id = cabinet.tenant_id,
            created_by = cabinet.created_by
        WHERE cabinets.id = cabinet.id
        RETURNING cabinets.id, cabinets.name, cabinets.location, cabinets.description, cabinets.tenant_id, cabinets.created_by, cabinets.created_at
        INTO new_cabinet;

            IF NOT FOUND THEN
                RAISE EXCEPTION 'record not found';
            END IF;

        RETURN new_cabinet;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Delete Cabinet

CREATE OR REPLACE FUNCTION delete_cabinet(c_id INT, user_id INT) 
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM cabinets WHERE id = c_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Add Cabinet to Team

CREATE OR REPLACE FUNCTION add_cabinet_to_team(cabinet cabinets_team_type, user_id INT)
RETURNS cabinets_team_type AS $$
DECLARE
    new_cabinets_team cabinets_team_type;
BEGIN
    IF check_permission(user_id) THEN
        IF EXISTS (SELECT 1 FROM cabinets_team WHERE cabinets_team.team_id = cabinet.team_id AND cabinets_team.cabinet_id = cabinet.cabinet_id) THEN
            RAISE EXCEPTION 'Cabinet with ID % is already in team with ID %', cabinet.cabinet_id, cabinet.team_id;
        ELSE
            INSERT INTO cabinets_team (team_id, cabinet_id, created_by, created_at)
            VALUES (cabinet.team_id, cabinet.cabinet_id, cabinet.created_by, NOW())
            RETURNING team_id, cabinet_id, created_by, created_at
            INTO new_cabinets_team;

            RETURN new_cabinest_team;
        END IF;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Remove Cabinet from Team

CREATE OR REPLACE FUNCTION remove_cabinet_from_team(cabinet cabinets_team_type, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM cabinets_team WHERE cabinets_team.team_id = cabinet.team_id AND cabinets_team.cabinet_id = cabinet.cabinet_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
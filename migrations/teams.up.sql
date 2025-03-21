CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_team (
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP,
    PRIMARY KEY (team_id, user_id)
);

CREATE TYPE team_type AS (
    id INT,
    name CITEXT,
    tenant_id INT,
    created_at TIMESTAMP
);

CREATE TYPE users_team_type AS (
    team_id INT,
    user_id INT,
    created_by INT,
    created_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION create_team(team team_type, user_id INT) 
RETURNS team_type AS $$
DECLARE
    new_team team_type;
BEGIN
    IF check_permission(user_id) THEN
        INSERT INTO teams (name, tenant_id)
        VALUES (team.name, team.tenant_id)
        RETURNING teams.id, teams.name, teams.tenant_id, teams.created_at
        INTO new_team;

        RETURN new_team;

    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Get Team By ID

CREATE OR REPLACE FUNCTION get_team_by_id(team_id INT)
RETURNS team_type AS $$
DECLARE
    new_team team_type;
BEGIN
    SELECT id, name, tenant_id, created_at
    INTO new_team
    FROM teams
    WHERE id = team_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_team;
END;
$$ LANGUAGE plpgsql;

-- Get Teams

CREATE OR REPLACE FUNCTION get_teams()
RETURNS SETOF team_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY
        SELECT id, name, tenant_id, created_at
        FROM teams;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Teams 

CREATE OR REPLACE FUNCTION update_team(team team_type, user_id INT)
RETURNS team_type AS $$
DECLARE
    new_team team_type;
BEGIN
    IF check_permission(user_id) THEN
        UPDATE teams
           SET name = team.name,
               tenant_id = team.tenant_id
        WHERE id = team.id
        RETURNING teams.id, teams.name, teams.tenant_id, teams.created_at
        INTO new_team;
  
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

        RETURN new_team;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Delete Team 

CREATE OR REPLACE FUNCTION delete_team(t_id INT, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM teams WHERE id = t_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Add User to Team

CREATE OR REPLACE FUNCTION add_user_to_team(input users_team_type, u_id INT)
RETURNS users_team_type AS $$
DECLARE 
    new_users_team users_team_type;
BEGIN
    IF check_permission(u_id) THEN
        IF EXISTS (SELECT 1 FROM users_team WHERE users_team.team_id = input.team_id AND users_team.user_id = input.user_id) THEN
            RAISE EXCEPTION 'User with ID % is already in team with ID %', input.user_id, input.team_id;
        ELSE
            INSERT INTO users_team (team_id, user_id, created_by, created_at)
            VALUES (input.team_id, input.user_id, input.created_by, NOW())
            RETURNING team_id, user_id, created_by, created_at
            INTO new_users_team;

            RETURN new_users_team;
        END IF;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Remove User from Team

CREATE OR REPLACE FUNCTION remove_user_from_team(input users_team_type, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM users_team WHERE users_team.team_id = input.team_id AND users_team.user_id = input.user_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
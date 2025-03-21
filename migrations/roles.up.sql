CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name CITEXT
);

CREATE TYPE role_type AS (
    id INT,
    name CITEXT
);

CREATE OR REPLACE FUNCTION check_permission(user_id INT)
RETURNS BOOLEAN AS $$
BEGIN
    RETURN (
        SELECT CASE 
            WHEN roles.name = 'admin' THEN TRUE
            ELSE FALSE
        END
        FROM users
        JOIN roles ON users.role_id = roles.id
        WHERE users.id = user_id
    );
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION create_role(role role_type)
RETURNS role_type AS $$
DECLARE
    new_role role_type;
BEGIN
    INSERT INTO roles (name)
    VALUES (role.name)
    RETURNING roles.id, roles.name
    INTO new_role;
    RETURN new_role;
END;
$$ LANGUAGE plpgsql;

-- Get role By ID

CREATE OR REPLACE FUNCTION get_role_by_id(t_id INT)
RETURNS role_type AS $$
DECLARE
    new_role role_type;
BEGIN
    SELECT id, name
    INTO new_role
    FROM roles
    WHERE id = t_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;
    RETURN new_role;
END;
$$ LANGUAGE plpgsql;

-- Get roles

CREATE OR REPLACE FUNCTION get_roles(user_id INT)
RETURNS SETOF role_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY
        SELECT id, name
        FROM roles;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update role

CREATE OR REPLACE FUNCTION update_role(role role_type, user_id INT)
RETURNS role_type AS $$
DECLARE
    new_role role_type;
BEGIN
    IF check_permission(user_id) THEN
        UPDATE roles
            SET name = role.name
            WHERE id = role.id
        RETURNING roles.id, roles.name
        INTO new_role;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

        RETURN new_role;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Delete role

CREATE OR REPLACE FUNCTION delete_role(r_id INT, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM roles WHERE id = r_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
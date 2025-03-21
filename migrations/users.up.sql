CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name CITEXT NOT NULL,
    last_name CITEXT NOT NULL,
    email CITEXT UNIQUE NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    role_id INT REFERENCES roles(id) ON DELETE SET NULL,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    invited_by INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE user_type AS (
    id INT,
    first_name CITEXT,
    last_name CITEXT,
    email CITEXT,
    phone TEXT,
    role_id INT,
    tenant_id INT,
    invited_by INT,
    created_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION create_user(user_input user_type, user_id INT) 
RETURNS user_type AS $$
DECLARE
    new_user user_type;
BEGIN
    IF check_permission(user_id) THEN
        INSERT INTO users (first_name, last_name, email, phone, role_id, tenant_id, invited_by)
        VALUES (user_input.first_name, user_input.last_name, user_input.email, user_input.phone, user_input.role_id, user_input.tenant_id, user_input.invited_by)
        RETURNING users.id, users.first_name, users.last_name, users.email, users.phone, users.role_id, users.tenant_id, users.invited_by, users.created_at
        INTO new_user;

        RETURN new_user;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Get User By ID

CREATE OR REPLACE FUNCTION get_user_by_id(u_id INT)
RETURNS user_type AS $$
DECLARE
    new_user user_type;
BEGIN
    SELECT id, first_name, last_name, email, phone, role_id, tenant_id, invited_by, created_at
    INTO new_user
    FROM users
    WHERE id = u_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_user;
END;
$$ LANGUAGE plpgsql;

-- Get Users

CREATE OR REPLACE FUNCTION get_users(user_id INT)
RETURNS SETOF user_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY 
        SELECT id, first_name, last_name, email, phone, role_id, tenant_id, invited_by, created_at
        FROM users;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update User
CREATE OR REPLACE FUNCTION update_user(user_input user_type, user_id INT)
RETURNS user_type AS $$
DECLARE
    new_user user_type;
BEGIN
    IF check_permission(user_id) THEN
        UPDATE users
           SET first_name = user_input.first_name,
               last_name = user_input.last_name,
               email = user_input.email,
               phone = user_input.phone,
               role_id = user_input.role_id,
               tenant_id = user_input.tenant_id,
               invited_by = user_input.invited_by
        WHERE users.id = user_input.id
        RETURNING users.id, users.first_name, users.last_name, users.email, users.phone, users.role_id, users.tenant_id, users.invited_by, users.created_at
        INTO new_user;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_user;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Delete User

CREATE OR REPLACE FUNCTION delete_user(u_id INT, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM users WHERE id = u_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'User with ID % not found', u_id;
        END IF;

    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
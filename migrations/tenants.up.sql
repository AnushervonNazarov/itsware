CREATE TABLE tenants (
    id SERIAL PRIMARY KEY,
    name CITEXT,
    is_support BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE tenant_type AS (
    id INT,
    name CITEXT,
    is_support BOOLEAN,
    created_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION create_tenant(tenant tenant_type, user_id INT)
RETURNS tenant_type AS $$
DECLARE
    new_tenant tenant_type;
BEGIN
    IF check_permission(user_id) THEN
        INSERT INTO tenants (name, is_support)
        VALUES (tenant.name, tenant.is_support)
        RETURNING tenants.id, tenants.name, tenants.is_support, tenants.created_at
        INTO new_tenant;

        RETURN new_tenant;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Get Tenant By ID

CREATE OR REPLACE FUNCTION get_tenant_by_id(t_id INT)
RETURNS tenant_type AS $$
DECLARE
    new_tenant tenant_type;
BEGIN
    SELECT id, name, is_support, created_at
    INTO new_tenant
    FROM tenants
    WHERE id = t_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_tenant;
END;
$$ LANGUAGE plpgsql;

-- Get Tenants

CREATE OR REPLACE FUNCTION get_tenants(user_id INT)
RETURNS SETOF tenant_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY
        SELECT id, name, is_support, created_at
        FROM tenants;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Tenant

CREATE OR REPLACE FUNCTION update_tenant(tenant tenant_type, user_id INT)
RETURNS tenant_type AS $$
DECLARE
    new_tenant tenant_type;
BEGIN
    IF check_permission(user_id) THEN
        UPDATE tenants
            SET name = tenant.name,
                is_support = tenant.is_support
            WHERE id = tenant.id
        RETURNING tenants.id, tenants.name, tenants.is_support, tenants.created_at
        INTO new_tenant;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

        RETURN new_tenant;
    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Delete Tenant

CREATE OR REPLACE FUNCTION delete_tenant(t_id INT, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM tenants WHERE id = t_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE    
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
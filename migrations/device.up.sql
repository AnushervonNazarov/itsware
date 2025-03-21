CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL,
    description TEXT,
    status CITEXT CHECK (status IN ('active', 'inactive', 'checked_out')) DEFAULT 'active',
    serial_number TEXT UNIQUE NOT NULL,
    checked_out_on TIMESTAMP,
    checked_out_by INT REFERENCES users(id) ON DELETE SET NULL,
    cabinet_id INT REFERENCES cabinets(id) ON DELETE SET NULL,
    tenant_id INT REFERENCES tenants(id) ON DELETE SET NULL,
    device_profile_id INT REFERENCES device_profiles(id) ON DELETE SET NULL,
    created_by INT REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE device_type AS (
    id INT,
    name CITEXT,
    description TEXT,
    status CITEXT,
    serial_number TEXT,
    checked_out_on TIMESTAMP,
    checked_out_by INT,
    cabinet_id INT,
    tenant_id INT,
    device_profile_id INT,
    created_by INT,
    created_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION create_device(device device_type, user_id INT)
RETURNS device_type AS $$
DECLARE
    new_device device_type;
BEGIN
    IF check_permission(user_id) THEN
        INSERT INTO devices (name, description, status, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_by)
        VALUES (device.name, device.description, device.status, device.serial_number, device.checked_out_on, device.checked_out_by, device.cabinet_id, device.tenant_id, device.device_profile_id, device.created_by)
        RETURNING devices.id ,devices.name, devices.description, devices.status, devices.serial_number, devices.checked_out_on, devices.checked_out_by, devices.cabinet_id, devices.tenant_id, devices.device_profile_id, devices.created_by, devices.created_at
        INTO new_device;

        RETURN new_device;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Get Device By ID


CREATE OR REPLACE FUNCTION get_device_by_id(p_id int)
RETURNS device_type AS $$
DECLARE
    new_device device_type;
BEGIN
    SELECT id, name, description, status, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_by, created_at
    INTO new_device
    FROM devices
    WHERE id = p_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'record not found';
    END IF;

    RETURN new_device;
END;
$$ LANGUAGE plpgsql;

-- Get Devices

CREATE OR REPLACE FUNCTION get_devices(user_id INT)
RETURNS SETOF device_type AS $$
BEGIN
    IF check_permission(user_id) THEN
        RETURN QUERY 
        SELECT id, name, description, status, serial_number, checked_out_on, checked_out_by, cabinet_id, tenant_id, device_profile_id, created_by, created_at
        FROM devices;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Update Device

CREATE FUNCTION update_device(device device_type, user_id INT)
RETURNS device_type AS $$
DECLARE
    new_device device_type;
BEGIN 
    IF check_permission(user_id) THEN
        UPDATE devices
        SET name = device.name,
            description = device.description,
            status = device.status,
            serial_number = device.serial_number,
            checked_out_on = device.checked_out_on,
            checked_out_by = device.checked_out_by,
            cabinet_id = device.cabinet_id,
            tenant_id = device.tenant_id,
            device_profile_id = device.device_profile_id,
            created_by = device.created_by
        WHERE devices.id = device.id
        RETURNING devices.id, devices.name, devices.description, devices.status, devices.serial_number, devices.checked_out_on, devices.checked_out_by, devices.cabinet_id, devices.tenant_id, devices.device_profile_id, devices.created_by, devices.created_at
        INTO new_device;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;  

        RETURN new_device;
    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;  
END;
$$ LANGUAGE plpgsql;

-- Delete Device

CREATE OR REPLACE FUNCTION delete_device(p_id INT, user_id INT)
RETURNS VOID AS $$
BEGIN
    IF check_permission(user_id) THEN
        DELETE FROM devices WHERE id = p_id;
        IF NOT FOUND THEN
            RAISE EXCEPTION 'record not found';
        END IF;

    ELSE
        RAISE EXCEPTION 'Permission denied';
    END IF;
END;
$$ LANGUAGE plpgsql;
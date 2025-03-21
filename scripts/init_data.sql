
INSERT INTO roles (name) VALUES ('support_admin');
INSERT INTO roles (name) VALUES ('support_view');
INSERT INTO roles (name) VALUES ('administrator');

INSERT INTO device_statuses (name) VALUES ('active');
INSERT INTO device_statuses (name) VALUES ('inactive');
INSERT INTO device_statuses (name) VALUES ('checked out');

INSERT INTO tenants (name, is_support, created_on) VALUES ('ItsWare', TRUE, NOW());
INSERT INTO users (first_name, last_name, password, email, phone, role_id, tenant_id, created_on) VALUES ('support_admin', 'user', 'q123', 'support@itsware.com', '', 1, 1, NOW());
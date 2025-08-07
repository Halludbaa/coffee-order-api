
CREATE TABLE IF NOT EXISTS admins (
    id          uuid        PRIMARY KEY DEFAULT gen_random_uuid(),
    admin_id    VARCHAR(50) NOT NULL,
    password    VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS baristas (
    id          uuid        PRIMARY KEY DEFAULT gen_random_uuid(),
    barista_id  VARCHAR(80) NOT NULL,
    password    VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    joined_at   TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS branches (
    id          uuid         PRIMARY KEY DEFAULT gen_random_uuid(),
    address     VARCHAR(255) DEFAULT NULL,
    password    VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id          uuid         PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    price       DECIMAL(9, 0) DEFAULT 0,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS branch_products (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id  VARCHAR(100) NOT NULL,
    branch_id   VARCHAR(100) NOT NULL,
    stock       DECIMAL(6, 0) DEFAULT 0,
    hot_status  BOOLEAN DEFAULT TRUE,
    iced_status BOOLEAN DEFAULT TRUE,
    available   BOOLEAN DEFAULT TRUE,
    CONSTRAINT product_id_fk FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
    CONSTRAINT branch_id_fk FOREIGN KEY(branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- NEXT ORDER SCHEMA in mongo
-- schema.sql
-- Cafe Order Management System (Indonesia / IDR Ready)
-- Status in English (safe for backend), UI can translate to Bahasa

-- 1. Stores
CREATE TABLE IF NOT EXISTS stores (
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    location      TEXT,
    address       TEXT,
    phone         VARCHAR(20),
    email         VARCHAR(100),
    store_slug    VARCHAR(50) UNIQUE NOT NULL,
    is_active     BOOLEAN DEFAULT true,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_stores_slug ON stores(store_slug);
CREATE INDEX IF NOT EXISTS idx_stores_active ON stores(is_active);

-- 2. Users
CREATE TABLE IF NOT EXISTS users (
    id                SERIAL PRIMARY KEY,
    full_name         VARCHAR(100) NOT NULL,
    email             VARCHAR(100) UNIQUE,
    role              VARCHAR(20) NOT NULL CHECK (role IN ('barista', 'manager', 'admin')),
    store_id          INT REFERENCES stores(id) ON DELETE CASCADE,
    pin_hash          BYTEA,
    is_active         BOOLEAN DEFAULT true,
    must_reset_pin    BOOLEAN DEFAULT true,
    created_at        TIMESTAMPTZ DEFAULT NOW(),
    updated_at        TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_store ON users(store_id);
CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);

-- 3. Categories
CREATE TABLE IF NOT EXISTS categories (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    icon        VARCHAR(50),
    color       VARCHAR(7) DEFAULT '#4B3621',
    is_active   BOOLEAN DEFAULT true,
    created_at  TIMESTAMPTZ DEFAULT NOW()
);

-- 4. Store Categories (Visibility)
CREATE TABLE IF NOT EXISTS store_categories (
    id             SERIAL PRIMARY KEY,
    store_id       INT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    category_id    INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    name           VARCHAR(50),
    is_visible     BOOLEAN DEFAULT true,
    sort_order     INT DEFAULT 0,
    UNIQUE(store_id, category_id)
);

CREATE INDEX IF NOT EXISTS idx_store_categories ON store_categories(store_id, category_id);

-- 5. Menu Items
CREATE TABLE IF NOT EXISTS menu_items (
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    description   TEXT,
    base_price    DECIMAL(12,0) NOT NULL CHECK (base_price >= 0),
    category_id   INT NOT NULL REFERENCES categories(id),
    is_active     BOOLEAN DEFAULT true,
    image_url     TEXT,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_menu_items_category ON menu_items(category_id);
CREATE INDEX IF NOT EXISTS idx_menu_items_active ON menu_items(is_active);

-- 6. Store Menu (Pricing)
CREATE TABLE IF NOT EXISTS store_menu (
    id               SERIAL PRIMARY KEY,
    store_id         INT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    menu_item_id     INT NOT NULL REFERENCES menu_items(id) ON DELETE CASCADE,
    price_override   DECIMAL(12,0),
    is_available     BOOLEAN DEFAULT true,
    sort_order       INT DEFAULT 0,
    created_at       TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(store_id, menu_item_id)
);

CREATE INDEX IF NOT EXISTS idx_store_menu ON store_menu(store_id, menu_item_id, is_available);

-- 7. Customization Groups
CREATE TABLE IF NOT EXISTS customization_groups (
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(50) NOT NULL,
    store_id     INT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    is_required  BOOLEAN DEFAULT false,
    sort_order   INT DEFAULT 0,
    created_at   TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_customization_groups_store ON customization_groups(store_id);

-- 8. Customization Options
CREATE TABLE IF NOT EXISTS customization_options (
    id                 SERIAL PRIMARY KEY,
    group_id           INT NOT NULL REFERENCES customization_groups(id) ON DELETE CASCADE,
    label              VARCHAR(50) NOT NULL,
    additional_price   DECIMAL(12,0) DEFAULT 0,
    is_available       BOOLEAN DEFAULT true,
    sort_order         INT DEFAULT 0,
    created_at         TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_customization_options_group ON customization_options(group_id);

-- 9. Menu Item Customizations
CREATE TABLE IF NOT EXISTS menu_item_customizations (
    id             SERIAL PRIMARY KEY,
    menu_item_id   INT NOT NULL REFERENCES menu_items(id) ON DELETE CASCADE,
    group_id       INT NOT NULL REFERENCES customization_groups(id) ON DELETE CASCADE,
    is_default     BOOLEAN DEFAULT false,
    UNIQUE(menu_item_id, group_id)
);

CREATE INDEX IF NOT EXISTS idx_menu_item_customizations ON menu_item_customizations(menu_item_id, group_id);

-- 10. Orders
CREATE TABLE IF NOT EXISTS orders (
    id             SERIAL PRIMARY KEY,
    store_id       INT NOT NULL REFERENCES stores(id),
    order_number   VARCHAR(20) NOT NULL UNIQUE,
    status         VARCHAR(20) NOT NULL DEFAULT 'pending'
                 CHECK (status IN ('pending', 'preparing', 'ready', 'completed', 'cancelled')),
    total          DECIMAL(12,0) NOT NULL,
    customer_note  TEXT,
    created_at     TIMESTAMPTZ DEFAULT NOW(),
    updated_at     TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_orders_store_status ON orders(store_id, status);
CREATE INDEX IF NOT EXISTS idx_orders_created ON orders(created_at);

-- 11. Order Items
CREATE TABLE IF NOT EXISTS order_items (
    id               SERIAL PRIMARY KEY,
    order_id         INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id     INT NOT NULL REFERENCES menu_items(id),
    quantity         INT NOT NULL DEFAULT 1 CHECK (quantity >= 1),
    unit_price       DECIMAL(12,0) NOT NULL,
    customizations   JSONB,
    note             TEXT,
    created_at       TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_order_items_order ON order_items(order_id);
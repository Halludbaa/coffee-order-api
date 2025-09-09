-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ===================================
-- 1. Stores
-- ===================================
INSERT INTO stores (id, name, location, address, phone, email, store_slug, is_active)
VALUES 
(
    'd1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22',
    'BrewCo Downtown',
    'Jakarta Pusat',
    'Jl. Sudirman No. 123, Jakarta',
    '021-1234-5678',
    'downtown@brewco.id',
    'downtown',
    true
),
(
    'e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33',
    'BrewCo Airport',
    'Soekarno-Hatta',
    'Terminal 3, Tangerang',
    '021-8765-4321',
    'airport@brewco.id',
    'airport',
    true
);

-- ===================================
-- 2. Categories
-- ===================================
INSERT INTO categories (id, name, icon, color, is_active)
VALUES 
('c1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 'Coffee', 'coffee', '#4B3621', true),
('d2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22', 'Tea', 'leaf', '#D69E2E', true),
('e3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44', 'Pastry', 'bread', '#8B5E3C', true);

-- ===================================
-- 3. Store Categories (Visibility & Naming)
-- ===================================
INSERT INTO store_categories (store_id, category_id, name, is_visible, sort_order)
VALUES 
-- Downtown
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'c1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 'Coffee', true, 1),
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'd2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22', 'Tea', true, 2),
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'e3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44', 'Kue', true, 3),

-- Airport
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'c1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 'Coffee', true, 1),
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'd2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22', 'Tea', true, 2),
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'e3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44', 'Pastry', false, 3); -- Pastry not available

-- ===================================
-- 4. Menu Items (Global)
-- ===================================
INSERT INTO menu_items (id, name, description, base_price, category_id, is_active, image_url)
VALUES 
(
    'm1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11',
    'Latte',
    'Espresso dengan susu steamed',
    25000,
    'c1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11',
    true,
    'https://example.com/images/latte.jpg'
),
(
    'm2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22',
    'Espresso',
    'Single shot espresso pekat',
    18000,
    'c1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11',
    true,
    'https://example.com/images/espresso.jpg'
),
(
    'm3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44',
    'Green Tea',
    'Teh hijau segar',
    15000,
    'd2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22',
    true,
    'https://example.com/images/green-tea.jpg'
),
(
    'm4f8d8f4-6e5f-7i4d-0i6d-4h8e1f0g9d55',
    'Croissant',
    'Pastry lembut dengan mentega',
    12000,
    'e3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44',
    true,
    'https://example.com/images/croissant.jpg'
);

-- ===================================
-- 5. Store Menu (Store-Specific Pricing & Availability)
-- ===================================
INSERT INTO store_menu (store_id, menu_item_id, price_override, is_available, sort_order)
VALUES 
-- Downtown Cafe
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'm1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 25000, true, 1),
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'm2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22', 18000, true, 2),
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'm3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44', 15000, true, 3),
('d1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', 'm4f8d8f4-6e5f-7i4d-0i6d-4h8e1f0g9d55', 12000, true, 4),

-- Airport Cafe (different prices, no croissant)
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'm1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 28000, true, 1), -- premium price
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'm2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22', 20000, true, 2),
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'm3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44', 17000, true, 3),
('e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', 'm4f8d8f4-6e5f-7i4d-0i6d-4h8e1f0g9d55', NULL, false, 4); -- not available

-- ===================================
-- 6. Customization Groups
-- =================================--
INSERT INTO customization_groups (id, name, store_id, is_required, sort_order)
VALUES 
('cg1-1234-5678-90ab', 'Jenis Susu', 'd1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', false, 1),
('cg2-2345-6789-01bc', 'Suhu', 'd1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', true, 2),
('cg3-3456-7890-12cd', 'Tambahan Shot', 'd1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22', false, 3);

-- Airport uses same groups
INSERT INTO customization_groups (id, name, store_id, is_required, sort_order)
VALUES 
('cg4-4567-8901-23de', 'Jenis Susu', 'e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', false, 1),
('cg5-5678-9012-34ef', 'Suhu', 'e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33', true, 2);

-- ===================================
-- 7. Customization Options
-- =================================--
INSERT INTO customization_options (group_id, label, additional_price, is_available, sort_order)
VALUES 
-- Downtown - Milk
('cg1-1234-5678-90ab', 'Susu Sapi', 0, true, 1),
('cg1-1234-5678-90ab', 'Susu Almond (+2k)', 2000, true, 2),
('cg1-1234-5678-90ab', 'Susu Oat (+3k)', 3000, true, 3),

-- Downtown - Temperature
('cg2-2345-6789-01bc', 'Panas', 0, true, 1),
('cg2-2345-6789-01bc', 'Es', 0, true, 2),

-- Downtown - Shots
('cg3-3456-7890-12cd', 'Extra Shot (+1.5k)', 1500, true, 1),

-- Airport - Milk
('cg4-4567-8901-23de', 'Whole Milk', 0, true, 1),
('cg4-4567-8901-23de', 'Almond Milk (+2k)', 2000, true, 2),

-- Airport - Temperature
('cg5-5678-9012-34ef', 'Hot', 0, true, 1),
('cg5-5678-9012-34ef', 'Iced', 0, true, 2);

-- ===================================
-- 8. Menu Item Customizations
-- =================================--
INSERT INTO menu_item_customizations (menu_item_id, group_id, is_default)
VALUES 
('m1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 'cg1-1234-5678-90ab', true), -- Latte - Milk
('m1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 'cg2-2345-6789-01bc', true), -- Latte - Temp
('m1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a11', 'cg3-3456-7890-12cd', false), -- Latte - Extra Shot

('m2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b22', 'cg2-2345-6789-01bc', true), -- Espresso - Temp

('m3f7c7f3-5d4e-6h3c-9h5c-3g7d0e9f8c44', 'cg5-5678-9012-34ef', true); -- Green Tea - Temp (Airport)

-- ===================================
-- 9. Users (with bcrypt-hashed PINs)
-- =================================--
-- PIN 1357 -> bcrypt hash (cost 10): $2a$10$YsG.e3w6Yv.zOZ9bYq1tUeO6q2v3w4x5y6z7A8B9C0D1E2F3G4H5I
-- You can generate with: bcrypt.GenerateFromPassword([]byte("1357"), 10)

INSERT INTO users (id, full_name, email, role, store_id, pin_hash, is_active, must_reset_pin)
VALUES 
(
    'u1a2b3c4-d5e6-7f8g-9h0i-1j2k3l4m5n67',
    'Admin Utama',
    'admin@brewco.id',
    'admin',
    NULL, -- admin not tied to one store
    '$2a$10$YsG.e3w6Yv.zOZ9bYq1tUeO6q2v3w4x5y6z7A8B9C0D1E2F3G4H5I',
    true,
    false
),
(
    'u2b3c4d5-e6f7-8g9h-0i1j-2k3l4m5n6o88',
    'Manager Downtown',
    'manager-downtown@brewco.id',
    'manager',
    'd1f5b5e1-3b2c-4f1a-9f3a-1e5b8c7d6a22',
    '$2a$10$YsG.e3w6Yv.zOZ9bYq1tUeO6q2v3w4x5y6z7A8B9C0D1E2F3G4H5I',
    true,
    false
),
(
    'u3c4d5e6-f7g8-9h0i-1j2k-3l4m5n6o7p99',
    'Barista Airport',
    'barista-airport@brewco.id',
    'barista',
    'e2f6c6f2-4c3d-5g2b-8g4b-2f6c9d8e7b33',
    '$2a$10$YsG.e3w6Yv.zOZ9bYq1tUeO6q2v3w4x5y6z7A8B9C0D1E2F3G4H5I',
    true,
    false
);

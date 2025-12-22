-- Создание типа ENUM для ролей администраторов
-- super_admin - полные права доступа ко всем функциям системы
-- admin - стандартные права администратора
CREATE TYPE admin_role AS ENUM ('super_admin', 'admin');

-- Таблица для хранения учетных записей администраторов системы
CREATE TABLE admins
(
    id            SERIAL PRIMARY KEY,                           -- Уникальный идентификатор администратора
    user_name     VARCHAR(100) UNIQUE NOT NULL,                 -- Имя пользователя (логин), должно быть уникальным
    password_hash VARCHAR(255)        NOT NULL,                 -- Хеш пароля (не храним пароли в открытом виде)
    role          admin_role          NOT NULL DEFAULT 'admin', -- Роль администратора, по умолчанию 'admin'
    created_at    TIMESTAMP                    DEFAULT NOW(),   -- Дата и время создания учетной записи
    created_by    INTEGER REFERENCES admins (id)                -- ID администратора, который создал эту запись (может быть NULL для первого админа)
);

-- Индекс для быстрого поиска администратора по имени пользователя (используется при авторизации)
CREATE INDEX idx_admins_user_name ON admins (user_name);
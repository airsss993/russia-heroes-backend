-- Включаем расширение PostGIS для работы с геоданными
CREATE EXTENSION IF NOT EXISTS postgis;

-- Создание типа ENUM для статусов заявок
-- pending - заявка ожидает модерации
-- approved - заявка одобрена
-- rejected - заявка отклонена
CREATE TYPE request_status AS ENUM ('pending', 'approved', 'rejected');

-- Таблица для хранения заявок на размещение сайтов от пользователей
CREATE TABLE site_requests
(
    id               SERIAL PRIMARY KEY,                                     -- Уникальный идентификатор заявки
    user_email       VARCHAR(255)              NOT NULL,                     -- Email пользователя, подавшего заявку
    user_name        VARCHAR(255)              NOT NULL,                     -- Имя пользователя, подавшего заявку
    site_title       VARCHAR(500)              NOT NULL,                     -- Название сайта/события
    site_description TEXT,                                                   -- Описание сайта/события
    location         GEOGRAPHY(Point, 4326),                                 -- Геолокация сайта (WGS84, используется в GPS)
    event_date       DATE                      NOT NULL,                     -- Дата события
    event_type_id    INTEGER                   NOT NULL REFERENCES event_types (id) ON DELETE RESTRICT, -- Тип события
    archive_path     VARCHAR(500),                                           -- Путь к архиву с файлами сайта
    extracted_path   VARCHAR(500),                                           -- Путь к распакованным файлами сайта
    preview_image_url VARCHAR(500),                                          -- URL превью-изображения сайта
    status           request_status            NOT NULL DEFAULT 'pending',   -- Статус модерации заявки
    rejection_reason TEXT,                                                   -- Причина отклонения заявки (заполняется при отклонении)
    submitted_at     TIMESTAMP                          DEFAULT NOW(),       -- Дата и время подачи заявки
    reviewed_at      TIMESTAMP,                                              -- Дата и время модерации заявки
    reviewed_by      INTEGER REFERENCES admins (id) ON DELETE SET NULL,     -- ID администратора, проверившего заявку

    CONSTRAINT check_rejection_reason CHECK (
        (status = 'rejected' AND rejection_reason IS NOT NULL AND rejection_reason != '')
            OR (status != 'rejected')
        ),
    CONSTRAINT check_review_fields CHECK (
        (reviewed_at IS NOT NULL AND reviewed_by IS NOT NULL)
            OR (reviewed_at IS NULL AND reviewed_by IS NULL)
        )
);

-- Индексы для быстрого поиска по статусу заявки
CREATE INDEX idx_site_requests_status ON site_requests (status);

-- Индекс для быстрого поиска по email пользователя
CREATE INDEX idx_site_requests_user_email ON site_requests (user_email);

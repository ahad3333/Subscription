
CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE subscription (
    id UUID PRIMARY KEY,
    title_ru VARCHAR,
    title_en VARCHAR,
    title_uz VARCHAR,
    price NUMERIC,
    duration_type VARCHAR,
    duration NUMERIC,
    is_active boolean,
    free_trial NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    delete_at  TIMESTAMP
);

CREATE TABLE subscription_user (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    subscription_id UUID NOT NULL REFERENCES subscription(id),
    free_trial_start_date TIMESTAMP,
    free_trial_end_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    status VARCHAR,
    send_notification boolean
);

CREATE TABLE user_subscription_history (
    id UUID PRIMARY KEY,
    user_subscription_id UUID NOT NULL REFERENCES subscription_user(id),
    status VARCHAR,
    start_date TIMESTAMP,
    end_date  TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    price NUMERIC
);


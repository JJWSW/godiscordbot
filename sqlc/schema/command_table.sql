CREATE TABLE command (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    command varchar NOT NULL,
    message varchar NOT NULL,
    target  varchar NOT NULL,
    args TEXT[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
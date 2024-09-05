CREATE TABLE schedule (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    episode INTEGER NOT NULL,
    title TEXT NOT NULL,
    guest TEXT[],
    description TEXT NOT NULL,
    running_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
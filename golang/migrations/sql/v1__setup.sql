CREATE TABLE power_plant (
    id SERIAL PRIMARY KEY,
    energy_in_kwh INTEGER NOT NULL,
    date_time TIMESTAMP WITH TIME ZONE
);
CREATE TABLE pcs (
		timestamp TIMESTAMPTZ NOT NULL,
		name TEXT NOT NULL,
		value DOUBLE PRECISION NOT NULL,
		PRIMARY KEY(timestamp, name)
);
SELECT create_hypertable('pcs', by_range('timestamp'));

package timescale

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/entity"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/domain/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TimescaleReadingRepository struct {
	wrapper repository.ReadingRepository
	conn    *pgxpool.Pool
}

func NewTimescaleReadingRepository(
	wrapper repository.ReadingRepository,
	conn *pgxpool.Pool,
) TimescaleReadingRepository {
	return TimescaleReadingRepository{
		wrapper,
		conn,
	}
}

func (r TimescaleReadingRepository) FindByInterval(
	ctx context.Context,
	interval entity.Interval,
) []entity.Reading {

	statement := `
		WITH time_series AS (
			SELECT generate_series(
				$1::timestamptz,
				$2::timestamptz,
				'1 hour'::interval
			) AS time
		)
		SELECT
			time_series.time,
			pcs.name,
			pcs.value
		FROM
			time_series
		LEFT JOIN
			pcs
		ON
			time_series.time = pcs.timestamp
		ORDER BY
			time_series.time;
	`

	rows, _ := r.conn.Query(
		ctx,
		statement,
		interval.StartTime(),
		interval.StopTime(),
	)
	output, err := pgx.CollectRows(
		rows, func (row pgx.CollectableRow) (entity.Reading, error) {
		var timestamp time.Time
		var name *string
		var value *float64
		err := rows.Scan(&timestamp, &name, &value)
		if err != nil {
			return entity.Reading{}, err
		}
		if name == nil || value == nil {
			return entity.Reading{}, fmt.Errorf("No name or value exist")
		}
		return entity.NewReading(timestamp, *name, *value), nil
	})
	if err != nil {
		return r.handleCacheMiss(ctx, interval)
	}
	return output
}

func (r TimescaleReadingRepository) handleCacheMiss(
	ctx context.Context,
	interval entity.Interval,
) []entity.Reading {
	wrapped := r.wrapper.FindByInterval(ctx, interval)
	if err := r.InsertInBatch(ctx, wrapped); err != nil {
		log.Println(err)
	}
	return wrapped
}

func (r TimescaleReadingRepository) InsertInBatch(
	ctx context.Context,
	readings []entity.Reading,
) error {
	queryInsertTimeseriesData := `
		INSERT INTO pcs (timestamp, name, value) VALUES ($1, $2, $3)
		ON CONFLICT (timestamp, name)
		DO UPDATE SET value = EXCLUDED.value;
	`
	batch := &pgx.Batch{}
	//load insert statements into batch queue
	for _, reading := range readings {
		batch.Queue(
			queryInsertTimeseriesData,
			reading.Timestamp(),
			reading.Name(),
			reading.Value(),
		)
	}

	return r.conn.SendBatch(ctx, batch).Close()
}

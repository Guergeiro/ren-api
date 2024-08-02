package timescale

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func beforeEach(ctx context.Context) (*postgres.PostgresContainer, error) {
	return postgres.RunContainer(ctx,
		testcontainers.WithImage("timescale/timescaledb:latest-pg14"),
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		postgres.WithInitScripts("../../../assets/postgres-entrypoint.sql"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
}

func afterEach(
	ctx context.Context,
	container *postgres.PostgresContainer,
) func() {
	return func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalln("failed to terminate container: ", err)
		}
	}
}

func TestFoo(t *testing.T) {
	ctx := context.Background()
	container, err := beforeEach(ctx)
	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(afterEach(ctx, container))

	// Do some integration testing with database
}

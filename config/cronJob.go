package config

import (
	"github.com/go-co-op/gocron"
	"log"
)

type Scheduler struct {
	cron     *gocron.Scheduler
	pgClient *PostgresClient
}

func NewScheduler(cron *gocron.Scheduler, pgClient *PostgresClient) *Scheduler {
	return &Scheduler{
		cron:     cron,
		pgClient: pgClient,
	}
}

// StartCronJobs cron job for update pawn orders
func (s *Scheduler) StartCronJobs() {

	_, cronErr := s.cron.Every(1).Day().Do(func() {
		s.pgClient.conn.Exec("CALL update_orders()")
		log.Print("calling stored procedure for update orders...")
	})

	if cronErr != nil {
		log.Print(cronErr)
	}

	//running job async
	s.cron.StartAsync()
}

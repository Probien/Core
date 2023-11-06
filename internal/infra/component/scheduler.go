package component

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/postgres"
	"github.com/go-co-op/gocron"
	"log"
)

type Scheduler struct {
	cron     *gocron.Scheduler
	pgClient *postgres.Client
}

func NewScheduler(cron *gocron.Scheduler, pgClient *postgres.Client) *Scheduler {
	return &Scheduler{
		cron:     cron,
		pgClient: pgClient,
	}
}

// StartCronJobs cron job for update pawn orders
func (s *Scheduler) StartCronJobs() {

	_, cronErr := s.cron.Every(1).Day().Do(func() {
		s.pgClient.Conn.Exec("CALL update_orders()")
		log.Print("updating orders...")
	})

	if cronErr != nil {
		log.Print(cronErr)
	}

	//running job async
	s.cron.StartAsync()
}

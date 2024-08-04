package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	eventHandler "gotik/internal/handler/event"
	"gotik/internal/provider/database"
	eventRepository "gotik/internal/repository/event"
	eventUsecase "gotik/internal/usecase/event"

	ticketHandler "gotik/internal/handler/ticket"
	ticketRepository "gotik/internal/repository/ticket"
	ticketUsecase "gotik/internal/usecase/ticket"

	tdHandler "gotik/internal/handler/transactiondetail"
	tdRepository "gotik/internal/repository/transactiondetail"
	tdUsecase "gotik/internal/usecase/transactiondetail"

	userHandler "gotik/internal/handler/user"
	userRepository "gotik/internal/repository/user"
	userUsecase "gotik/internal/usecase/user"
)

func autowired(db *sql.DB) (eventH eventHandler.EventHandler, ticketH ticketHandler.TicketHandler, tdH tdHandler.TransactionDetailHandler, userH userHandler.UserHandler) {
	ticketRepo := ticketRepository.NewTicketRepository(db)
	ticketUsecase := ticketUsecase.NewTicketUsecase(ticketRepo)
	ticketH = ticketHandler.NewTicketHandler(ticketUsecase)

	userRepo := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userH = userHandler.NewUserHandler(userUsecase)

	tdRepo := tdRepository.NewTransactionDetailRepository(db)
	tdUsecase := tdUsecase.NewTransactionDetailUsecase(tdRepo)
	tdH = tdHandler.NewTransactionDetailHandler(tdUsecase)

	eventRepo := eventRepository.NewEventRepository(db)
	eventUsecase := eventUsecase.NewEventUsecase(eventRepo, userRepo, tdRepo, ticketRepo)
	eventH = eventHandler.NewEventHandler(eventUsecase)

	return
}

type Routes struct {
	Server *gin.Engine
}

func NewRoutes() *Routes {
	db, err := database.NewDB()
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to database")
	}

	eventH, ticketH, tdH, userH := autowired(db)

	r := &Routes{
		Server: gin.New(),
	}

	// Routes
	r.Server.GET("/hello", func(c *gin.Context) {
		log.Info().Msg("Hello")
		c.Writer.Write([]byte("Hello World"))
	})

	api := r.Server.Group("/api")
	r.Event(api, eventH)
	r.Ticket(api, ticketH)
	r.User(api, userH)
	r.TransactionDetail(api, tdH)

	return r
}

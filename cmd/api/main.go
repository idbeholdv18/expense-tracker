package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/auth"
	"github/idbeholdv18/expense-tracker/internal/domain"
	expensetypes "github/idbeholdv18/expense-tracker/internal/expense_types"
	"github/idbeholdv18/expense-tracker/internal/expenses"
	"github/idbeholdv18/expense-tracker/internal/security/password"
	"github/idbeholdv18/expense-tracker/internal/token"
	"github/idbeholdv18/expense-tracker/internal/user"

	// "github/idbeholdv18/expense-tracker/internal/provider"

	transport_errors "github/idbeholdv18/expense-tracker/internal/transport/http/errors"
	"github/idbeholdv18/expense-tracker/internal/transport/http/middleware"
	"log"
	"net/http"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	port := flag.Int("port", 8080, "API server port")

	flag.Parse()

	db, err := sql.Open("pgx", "postgres://idbeholdv:idbeholdv@localhost:5433/expense_tracker?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to Postgres:", err)
	}

	fmt.Println("Connected to Postgres")

	userRepo := user.NewUserRepository(db)
	expensesRepo := expenses.NewExpensesRepository(db)
	expenseTypesRepo := expensetypes.NewExepenseTypesRepository(db)

	authService := &auth.AuthService{
		Repo: userRepo,
		Hasher: &password.BcryptHasher{
			Cost: bcrypt.DefaultCost,
		},
	}

	expensesService := &expenses.ExpenseService{
		Repo: expensesRepo,
	}

	tokenService := &token.TokenService{
		Secret: []byte("secret"),
	}

	expenseTypesService := &expensetypes.ExpenseTypesService{
		Repo: expenseTypesRepo,
	}

	authHandler := auth.AuthHandler{
		Auth:  authService,
		Token: tokenService,
	}

	expensesHandler := expenses.ExpenseHandler{
		Service: expensesService,
	}

	expenseTypesHandler := expensetypes.ExpenseTypesHandler{
		Service: expenseTypesService,
	}

	cors := middleware.CorsMiddleware(&middleware.CorsConfig{
		AllowedOrigin: "https://localhost:3000",
	})

	domain.RegisterErrors(transport_errors.Register)
	auth.RegisterErrors(transport_errors.Register)
	token.RegisterErrors(transport_errors.Register)
	user.RegisterErrors(transport_errors.Register)
	expenses.RegisterErrors(transport_errors.Register)
	expensetypes.RegisterErrors(transport_errors.Register)
	middleware.RegisterJwtErrors(transport_errors.Register)

	jwt := middleware.JwtMiddleware(tokenService)

	http.Handle("/api/v1/login", middleware.ErrorMiddleware(cors(authHandler.HandleLogin())))
	http.Handle("/api/v1/register", middleware.ErrorMiddleware(cors(authHandler.HandleRegister())))
	http.Handle("/api/v1/expenses", middleware.ErrorMiddleware(cors(jwt(expensesHandler.HandleExpense()))))
	http.Handle("/api/v1/expense-types", middleware.ErrorMiddleware(cors(jwt(expenseTypesHandler.HandleExpenseTypes()))))

	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(*port), "cert/cert.pem", "cert/key.pem", nil))
}

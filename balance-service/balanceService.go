package main

import (
	"context"
	"database/sql"

	repo "github.com/asciiu/gomo/balance-service/db/sql"
	bp "github.com/asciiu/gomo/balance-service/proto/balance"
)

type BalanceService struct {
	DB *sql.DB
}

func (service *BalanceService) GetUserBalance(ctx context.Context, req *bp.GetUserBalanceRequest, res *bp.BalanceResponse) error {
	balance, error := repo.FindBalance(service.DB, req)

	if error == nil {
		res.Status = "success"
		res.Data = &bp.UserBalanceData{
			Balance: balance,
		}
	} else {
		res.Status = "error"
		res.Message = error.Error()
	}

	return error
}
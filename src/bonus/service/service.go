package service

import (
	"bonus/model"
	"bonus/repository"
	"context"
)

type Service interface {
	GetPrivilege(ctx context.Context, username string) *model.PrivilegeInfoResponse
	UpdateBalanceAndHistory(ctx context.Context, username string, history model.BalanceHistory) error
}

type serviceImpl struct {
	repo repository.Repo
}

func NewService(repo repository.Repo) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) GetPrivilege(ctx context.Context, username string) *model.PrivilegeInfoResponse {

	privilege, err := s.repo.GetPrivilege(ctx, username)

	if err != nil {
		privilege, err = s.repo.CreatePrivilege(ctx, username)
		if err != nil {
			return nil
		}
	}

	res := model.PrivilegeInfoResponse{
		Balance: privilege.Balance,
		History: []model.BalanceHistory{},
		Status:  model.PrivilegeInfoResponseStatus(privilege.Status),
	}

	histories, err := s.repo.ListPrivilegeHistories(ctx, privilege.ID)

	if err != nil {
		return nil
	}

	for _, h := range histories {
		res.History = append(res.History, model.BalanceHistory{
			BalanceDiff:   h.BalanceDiff,
			Date:          h.Datetime,
			OperationType: model.BalanceHistoryOperationType(h.OperationType),
			TicketUid:     h.TicketUid,
		})
	}

	return &res
}

func (s *serviceImpl) UpdateBalanceAndHistory(ctx context.Context, username string, history model.BalanceHistory) error {
	privilege, err := s.repo.GetPrivilege(ctx, username)

	if err != nil {
		return err
	}

	err = s.repo.UpdatePrivilegeBalance(ctx, repository.UpdatePrivilegeBalanceParams{
		Username: username,
		Balance:  privilege.Balance + history.BalanceDiff,
	})

	if err != nil {
		return err
	}

	_, err = s.repo.CreatePrivilegeHistory(ctx, repository.CreatePrivilegeHistoryParams{
		PrivilegeID:   privilege.ID,
		TicketUid:     history.TicketUid,
		BalanceDiff:   history.BalanceDiff,
		OperationType: string(history.OperationType),
	})

	return err
}

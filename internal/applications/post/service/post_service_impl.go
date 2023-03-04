package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/post/dto"
	"myapp/internal/applications/post/repository/db"
	"myapp/internal/applications/transaction"
)

type PostServiceImpl struct {
	repository  db.PostRepository
	transaction *transaction.TxService
}

func NewPostServiceImpl(repository db.PostRepository, transaction *transaction.TxService) *PostServiceImpl {
	return &PostServiceImpl{repository: repository, transaction: transaction}
}

func (s *PostServiceImpl) Create(ctx context.Context, request dto.PostRequest) (*ent.Post, error) {

	post, err := s.repository.Create(ctx, request)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return post, nil
}

func (s *PostServiceImpl) Update(ctx context.Context, request dto.PostRequest, id int) (*ent.Post, error) {
	post, err := s.repository.Update(ctx, request, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return post, nil
}

func (s *PostServiceImpl) SoftDelete(ctx context.Context, id int) (*ent.Post, error) {

	var dataDeleted = &ent.Post{}

	//start transaction:
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {
		dataDel, errDeleted := s.softDeleteTrx(ctx, id, tx)
		dataDeleted = dataDel
		return errDeleted
	}); err != nil {
		return nil, err
	}

	return dataDeleted, nil
}

func (s *PostServiceImpl) softDeleteTrx(ctx context.Context, id int, tx *ent.Tx) (*ent.Post, error) {

	data, err := s.repository.SoftDelete(ctx, id, tx.Client())
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	//for testing transaction, data not update deleted_at
	if true {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return data, nil
}

func (s *PostServiceImpl) Delete(ctx context.Context, id int) (*ent.Post, error) {

	exist, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	_, err = s.repository.Delete(ctx, exist.ID)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return exist, nil
}

func (s *PostServiceImpl) GetById(ctx context.Context, id int) (*ent.Post, error) {

	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}

func (s *PostServiceImpl) GetAll(ctx context.Context) ([]*ent.Post, error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}

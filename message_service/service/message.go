package service

import (
	"contact_service/genproto/contact_service"
	"contact_service/pkg/helper"
	"contact_service/pkg/logger"
	"contact_service/storage"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
)

type contactService struct {
	logger  logger.Logger
	storage storage.StorageI
	contact_service.UnimplementedContactServiceServer
}

func NewContactService(db *sqlx.DB, log logger.Logger) *contactService {
	return &contactService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *contactService) Create(ctx context.Context, req *contact_service.Contact) (*contact_service.ContactId, error) {
	fmt.Println(s.storage)
	id, err := s.storage.Contact().Create(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while create contact", req, codes.Internal)
	}

	return &contact_service.ContactId{
		Id: id,
	}, nil
}

func (s *contactService) GetAll(ctx context.Context, req *contact_service.UserId) (*contact_service.GetAllContactResponse, error) {
	resp, err := s.storage.Contact().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all contacts", req, codes.Internal)
	}

	return resp, nil
}

func (s *contactService) Get(ctx context.Context, req *contact_service.ContactUserId) (*contact_service.Contact, error) {
	contact, err := s.storage.Contact().Get(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting contact", req, codes.Internal)
	}

	return contact, nil
}

func (s *contactService) Update(ctx context.Context, req *contact_service.Contact) (*contact_service.ContactUpdate, error) {
	resp, err := s.storage.Contact().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating contact", req, codes.Internal)
	}

	return &contact_service.ContactUpdate{
		Name:  resp.Name,
		Phone: resp.Phone,
	}, nil
}

func (s *contactService) Delete(ctx context.Context, req *contact_service.ContactUserId) (*contact_service.ContactDelete, error) {
	resp, err := s.storage.Contact().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting contact", req, codes.Internal)
	}

	return &contact_service.ContactDelete{
		Name:  resp,
	}, nil
}

package main

import (
	"context"
	"database/sql"
	"log"

	repo "github.com/asciiu/gomo/key-service/db/sql"
	kp "github.com/asciiu/gomo/key-service/proto/key"
	micro "github.com/micro/go-micro"
)

type KeyService struct {
	DB     *sql.DB
	KeyPub micro.Publisher
}

func (service *KeyService) AddKey(ctx context.Context, req *kp.KeyRequest, res *kp.KeyResponse) error {
	apiKey, error := repo.InsertKey(service.DB, req)

	switch {
	case error == nil:
		// publish a new key event
		if err := service.KeyPub.Publish(context.Background(), apiKey); err != nil {
			log.Println("could not publish event new.key: ", err)
		}

		res.Status = "success"
		res.Data = &kp.UserKeyData{
			Key: &kp.Key{
				KeyID:       apiKey.KeyID,
				UserID:      apiKey.UserID,
				Exchange:    apiKey.Exchange,
				Key:         apiKey.Key,
				Description: apiKey.Description,
				Status:      apiKey.Status,
			},
		}
		return nil

	default:
		res.Status = "error"
		res.Message = error.Error()
		return error
	}
}

func (service *KeyService) GetUserKey(ctx context.Context, req *kp.GetUserKeyRequest, res *kp.KeyResponse) error {
	apiKey, error := repo.FindKeyByID(service.DB, req)

	switch {
	case error == nil:
		res.Status = "success"
		res.Status = "success"
		res.Data = &kp.UserKeyData{
			Key: &kp.Key{
				KeyID:       apiKey.KeyID,
				UserID:      apiKey.UserID,
				Exchange:    apiKey.Exchange,
				Key:         apiKey.Key,
				Description: apiKey.Description,
				Status:      apiKey.Status,
			},
		}
		return nil
	default:
		res.Status = "error"
		res.Message = error.Error()
		return error
	}
}

func (service *KeyService) GetUserKeys(ctx context.Context, req *kp.GetUserKeysRequest, res *kp.KeyListResponse) error {
	keys, error := repo.FindKeysByUserID(service.DB, req)

	switch {
	case error == nil:
		res.Status = "success"
		res.Data = &kp.UserKeysData{
			Keys: keys,
		}
		return nil
	default:
		res.Status = "error"
		res.Message = error.Error()
		return error
	}
}

func (service *KeyService) RemoveKey(ctx context.Context, req *kp.RemoveKeyRequest, res *kp.KeyResponse) error {
	error := repo.DeleteKey(service.DB, req.KeyID)
	switch {
	case error == nil:
		res.Status = "success"
		return nil
	default:
		res.Status = "error"
		res.Message = error.Error()
		return error
	}
}

func (service *KeyService) UpdateKeyDescription(ctx context.Context, req *kp.KeyRequest, res *kp.KeyResponse) error {
	apiKey, error := repo.UpdateKeyDescription(service.DB, req)
	switch {
	case error == nil:
		res.Status = "success"
		res.Data = &kp.UserKeyData{
			Key: &kp.Key{
				KeyID:       apiKey.KeyID,
				UserID:      apiKey.UserID,
				Exchange:    apiKey.Exchange,
				Key:         apiKey.Key,
				Description: apiKey.Description,
				Status:      apiKey.Status,
			},
		}
		return nil
	default:
		res.Status = "error"
		res.Message = error.Error()
		return error
	}
}

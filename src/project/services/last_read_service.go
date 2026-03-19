package services

import (
	"errors"
	"strings"

	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/repositories"
)

type LastReadServiceInt interface {
	GetLastRead(uid string) (models.LastReadResp, error)
	SaveLastRead(req models.LastReadReq) (models.LastReadResp, error)
}

type LastReadService struct {
	repo repositories.LastReadRepo
}

func NewLastReadService(repo repositories.LastReadRepo) *LastReadService {
	return &LastReadService{repo: repo}
}

func (s *LastReadService) GetLastRead(uid string) (models.LastReadResp, error) {
	if strings.TrimSpace(uid) == "" {
		return models.LastReadResp{}, errors.New("firebaseUID is required")
	}
	return s.repo.GetByFirebaseUID(uid)
}

func (s *LastReadService) SaveLastRead(req models.LastReadReq) (models.LastReadResp, error) {
	if strings.TrimSpace(req.FirebaseUID) == "" {
		return models.LastReadResp{}, errors.New("firebaseUID is required")
	}
	if req.SuraID <= 0 {
		return models.LastReadResp{}, errors.New("suraId must be > 0")
	}
	if req.VerseID <= 0 {
		return models.LastReadResp{}, errors.New("verseId must be > 0")
	}
	if req.AyahID <= 0 {
		return models.LastReadResp{}, errors.New("ayahId must be > 0")
	}
	return s.repo.Upsert(req)
}

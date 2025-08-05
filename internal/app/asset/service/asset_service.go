package services

import (
	"context"
	"encoding/base64"
	webrtc "smart-city/internal/app/asset/delivery/webRTC"
	"smart-city/internal/app/asset/dto"
	assetRepositories "smart-city/internal/app/asset/repository"
	premiseRepositories "smart-city/internal/app/premise/repository"
	"smart-city/internal/models"

	"path/filepath"
	"smart-city/pkg/errors"
	"smart-city/pkg/utils"

	"github.com/google/uuid"
)

type Service struct {
	assetRepo    assetRepositories.AssetRepository
	premiseRepo  premiseRepositories.PremiseRepository
	snapshotRepo assetRepositories.SnapshotRepository
}

func NewAssetService(assetRepo assetRepositories.AssetRepository, premiseRepo premiseRepositories.PremiseRepository, snapshotRepo assetRepositories.SnapshotRepository) *Service {
	return &Service{assetRepo: assetRepo, premiseRepo: premiseRepo, snapshotRepo: snapshotRepo}
}

func (s *Service) CreateAsset(ctx context.Context, createAssetDto *dto.CreateAssetDto) (*models.Asset, error) {
	asset := &models.Asset{
		Name:     createAssetDto.Name,
		Location: createAssetDto.Location,
		Type:     createAssetDto.Type,
		IsActive: false,
	}
	if createAssetDto.PremiseID != "" {
		premiseID, err := uuid.Parse(createAssetDto.PremiseID)

		if err != nil {
			return nil, err
		}

		premise, err := s.premiseRepo.GetPremiseByID(ctx, premiseID.String())

		if err != nil {
			return nil, err
		}
		asset.Premise = premise
		asset.PremiseID = premiseID

	}

	return s.assetRepo.CreateAsset(ctx, asset)
}

func (s *Service) GetAssets(ctx context.Context) ([]models.Asset, error) {
	return s.assetRepo.GetAssets(ctx)
}

func (s *Service) StartPublishing(ctx context.Context, req *dto.SdpRequest) (*dto.SdpResponse, error) {
	broadcaster := webrtc.NewBroadcaster(req.CameraID)
	//mockup sdp
	mockSDP := `v=0
				o=- 46117327 2 IN IP4 127.0.0.1
				s=-
				t=0 0
				a=group:BUNDLE 0
				a=msid-semantic:WMS
				m=audio 9 UDP/TLS/RTP/SAVPF 111
				c=IN IP4 0.0.0.0
				a=rtpmap:111 opus/48000/2
				`

	encoded := base64.StdEncoding.EncodeToString([]byte(mockSDP))

	decoded, err := webrtc.DecodeSDP(encoded)
	if err != nil {
		return nil, err
	}
	answer, err := broadcaster.StartBroadcast(decoded)
	if err != nil {
		return nil, err
	}

	return &dto.SdpResponse{SDP: *answer}, nil
}

// CreateSnapshot handles file upload and creates a snapshot record
func (s *Service) CreateSnapshot(ctx context.Context, req *dto.SnapshotRequest, uploadDir string) (*dto.SnapshotResponse, error) {
	// Validate asset exists
	assetID, err := uuid.Parse(req.AssetID)
	if err != nil {
		return nil, errors.NewBadRequestError("Invalid asset ID format")
	}

	asset, err := s.assetRepo.GetAssetByID(ctx, assetID.String())
	if err != nil {
		return nil, errors.NewNotFoundError("asset")
	}

	// Validate file
	const maxFileSize = 10 * 1024 * 1024 // 10MB
	if err := utils.ValidateImageFile(req.File, maxFileSize); err != nil {
		return nil, errors.NewValidationError("Invalid file", map[string]string{"file": err.Error()})
	}

	// Create upload directory for asset snapshots
	assetUploadDir := filepath.Join(uploadDir, "snapshots", asset.ID.String())

	// Save the uploaded file
	fileInfo, err := utils.SaveUploadedFile(req.File, assetUploadDir)
	if err != nil {
		return nil, errors.NewInternalError("Failed to save uploaded file", err)
	}

	// Create snapshot record
	snapshot := &models.Snapshot{
		AssetID:  assetID,
		FileName: fileInfo.SavedName,
		FileSize: fileInfo.Size,
		FileType: fileInfo.ContentType,
		FilePath: fileInfo.Path,
	}

	createdSnapshot, err := s.snapshotRepo.CreateSnapshot(ctx, snapshot)
	if err != nil {
		// Clean up uploaded file if database operation fails
		utils.DeleteFile(fileInfo.Path)
		return nil, errors.NewDatabaseError("create snapshot", err)
	}

	// Return response
	return &dto.SnapshotResponse{
		ID:          createdSnapshot.ID.String(),
		AssetID:     createdSnapshot.AssetID.String(),
		Description: createdSnapshot.Description,
		Timestamp:   createdSnapshot.Timestamp,
		FileName:    createdSnapshot.FileName,
		FileSize:    createdSnapshot.FileSize,
		FileType:    createdSnapshot.FileType,
		FilePath:    createdSnapshot.FilePath,
		CreatedAt:   createdSnapshot.CreatedAt,
	}, nil
}

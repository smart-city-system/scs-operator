package http

import (
	"smart-city/internal/app/asset/dto"
	services "smart-city/internal/app/asset/service"
	"smart-city/pkg/errors"
	"smart-city/pkg/validation"

	"github.com/labstack/echo/v4"
)

// Handler
type Handler struct {
	svc services.Service
}

// NewHandler constructor
func NewHandler(svc services.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		createAssetDto := &dto.CreateAssetDto{}
		if err := c.Bind(createAssetDto); err != nil {
			return err
		}
		createdAsset, err := h.svc.CreateAsset(c.Request().Context(), createAssetDto)
		if err != nil {
			return err
		}
		return c.JSON(201, createdAsset)
	}
}

func (h *Handler) GetAssets() echo.HandlerFunc {
	return func(c echo.Context) error {
		Assets, err := h.svc.GetAssets(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, Assets)
	}
}

func (h *Handler) StartPublishing() echo.HandlerFunc {
	return func(c echo.Context) error {
		sdpRequest := &dto.SdpRequest{}
		if err := c.Bind(sdpRequest); err != nil {
			return err
		}
		sdp, err := h.svc.StartPublishing(c.Request().Context(), sdpRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, sdp)
	}
}

func (h *Handler) CreateSnapshot() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse multipart form
		if err := c.Request().ParseMultipartForm(10 << 20); err != nil { // 10MB max
			return errors.NewBadRequestError("Failed to parse multipart form")
		}

		// Extract form data
		snapShotRequest := &dto.SnapshotRequest{}

		// Get form values
		snapShotRequest.AssetID = c.FormValue("Asset_id")
		snapShotRequest.Description = c.FormValue("description")

		// Parse timestamp if provided
		if timestampStr := c.FormValue("timestamp"); timestampStr != "" {
			// You can add timestamp parsing logic here if needed
		}

		// Extract file from form
		file, err := c.FormFile("file")
		if err != nil {
			return errors.NewBadRequestError("File is required")
		}
		snapShotRequest.File = file

		// Validate the request
		if err := validation.ValidateStruct(snapShotRequest); err != nil {
			return err
		}

		// Define upload directory (you might want to make this configurable)
		uploadDir := "./uploads"

		// Create snapshot
		snapshot, err := h.svc.CreateSnapshot(c.Request().Context(), snapShotRequest, uploadDir)
		if err != nil {
			return err
		}

		return c.JSON(201, snapshot)
	}
}

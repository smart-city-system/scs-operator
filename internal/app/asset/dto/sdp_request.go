package dto

import "github.com/pion/webrtc/v4"

type SdpRequest struct {
	CameraID string `json:"camera_id"`
	SDP      string `json:"sdp"`
	ViewerID string `json:"viewer_id,omitempty"`
}

type SdpResponse struct {
	SDP webrtc.SessionDescription `json:"sdp"`
}

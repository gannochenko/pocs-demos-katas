package v1

import imagepb "backend/proto/image/v1"

type ImageController struct {
	imagepb.UnimplementedImageServiceServer
}

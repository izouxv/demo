package api

import (
	pb "file-server/api"
	. "file-server/common"
	"file-server/internal/storage"

	"github.com/prometheus/common/log"
	"golang.org/x/net/context"
)

type FileServer struct {
}

func (this *FileServer) GetFileList(ctx context.Context, in *pb.GetFileListRequest) (*pb.GetFileListResponse, error) {
	if len(in.Fids) == 0 {
		return nil, errToRPCError(InvalidArgument)
	}
	file := storage.File{}
	files, err := file.GetList(in.Fids...)
	if err != nil {
		log.Errorf("%s", err)
		return nil, errToRPCError(err)
	}
	resp := make([]*pb.File, 0)
	for _, v := range files {
		resp = append(resp, returnFileResponse(v))
	}
	return &pb.GetFileListResponse{Files: resp}, nil
}

func returnFileResponse(file *storage.File) (resp *pb.File) {
	resp = &pb.File{
		Fid:        file.Fid,
		Name:       file.Name,
		Ext:        file.Ext,
		Path:       file.Path,
		Size:       int32(file.Size),
		CreateTime: file.CreateTime.Unix(),
		UpdateTime: file.UpdateTime.Unix(),
		Url:        file.Url,
	}
	return
}

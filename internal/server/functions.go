package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tages_test/gRPC/protobuf"
	"time"
)

func (s *GRPCServer) DownloadFile(req *protobuf.DownloadRequest, fs protobuf.FileService_DownloadFileServer) error {
	fileName := req.GetFilename()
	extension := req.GetFileFormat()
	fullPath := s.Filepath
	fullName := fileName + extension
	resultStr := fmt.Sprintf("%s/%s", fullPath, fullName)

	fmt.Println(resultStr)
	dwFile, err := os.Open(resultStr)
	if err != nil {
		log.Println("Error reading downloaded file", fileName)
		return err
	}
	defer dwFile.Close()

	info, _ := dwFile.Stat()
	parts := strings.Split(fileName, "_")
	createDate := ""
	switch len(parts) {
	case 1:
		{
			createDate = info.ModTime().String()
		}
	default:
		{
			createDate = parts[0]
			fileName = strings.Join(parts[1:], "_")
		}

	}
	lastUpdate := info.ModTime().String()

	res := &protobuf.DownloadResponse{
		Image: &protobuf.Image{
			Info: &protobuf.FileInfo{
				Filename:   fileName,
				Fileformat: extension,
				UpdatedAt:  lastUpdate,
				CreatedAt:  createDate,
			},
		},
	}

	chunk := make([]byte, 100_000)
	for {
		n, err := dwFile.Read(chunk)
		if err == io.EOF {
			log.Println("File is downloaded", fileName)
			return nil
		}
		res.Image.Data = chunk[:n]
		if err := fs.Send(res); err != nil {
			log.Println("Error sending file throw ", fileName)
			return err
		}

	}
}

func (s *GRPCServer) UploadFile(server protobuf.FileService_UploadFileServer) error {
	if files, err := os.ReadDir(s.Filepath); err != nil || len(files) == 0 {
		log.Println("Error reading directory", s.Filepath)
		return err
	}

	req, err := server.Recv()
	if err == io.EOF {
		log.Println("File is already uploaded")
		return nil
	} else if err != nil {
		return err
	}
	fileName := req.Image.Info.Filename + req.Image.Info.Fileformat
	if fileName == "" {
		log.Println("No such file")
		return nil
	}

	log.Println("Starting upload", fileName)
	var newFilePath string

	err = filepath.Walk(s.Filepath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() != fileName {
			timestamp := time.Now()
			createDate := timestamp.Format("2006-01-02")
			newFileName := fmt.Sprintf("%s_%s", createDate, fileName)
			newFilePath = filepath.Join(s.Filepath, newFileName)
		}
		return nil
	})

	log.Println("Starting upload ", fileName)

	file, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error opening file", fileName)
		return err
	}
	defer file.Close()

	for i := 0; ; i++ {
		req, err := server.Recv()
		if err == io.EOF {
			log.Println("File is uploaded")
			return nil
		}
		if err != nil {
			return err
		}
		_, err = file.Write(req.Image.Data)
		if err != nil {
			log.Println("Error writing file", fileName)
			return err
		}
		fmt.Printf("downloading part: %d\n", i)

	}

}

func (s *GRPCServer) ListFiles(ctx context.Context, req *protobuf.ListFilesRequest) (res *protobuf.ListFilesResponse, err error) {
	files, _ := os.ReadDir(s.Filepath)
	listOfFiles := make([]string, len(files))
	fmt.Println(len(files))
	res = &protobuf.ListFilesResponse{
		Files: listOfFiles,
	}

	for i, file := range files {
		fullPath := filepath.Join(s.Filepath, file.Name())
		info, err := os.Stat(fullPath)
		if err != nil {
			log.Println("Error reading file", file.Name())
			continue
		}
		fileName := info.Name()
		modTime := info.ModTime().String()
		parts := strings.Split(fileName, "_")
		createDate := ""
		switch len(parts) {
		case 1:
			{
				createDate = info.ModTime().String()
			}
		default:
			{
				createDate = parts[0]
				fileName = strings.Join(parts[1:], "_")
			}

		}
		reqString := fmt.Sprintf("%s | %s | %s ", fileName, createDate, modTime)
		listOfFiles[i] = reqString
	}

	if len(listOfFiles) == 0 {
		str := "empty directory"
		listOfFiles = append(listOfFiles, str)
	}

	log.Println("full list of files")

	return res, nil
}

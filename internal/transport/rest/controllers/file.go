package controllers

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/transport/rest/response"
	"github.com/supernova0730/ae86/pkg/minio"
	"io"
	"mime/multipart"
	"net/http"
)

type FileController struct {
	service container.IService
}

func NewFileController(service container.IService) *FileController {
	return &FileController{service: service}
}

func (ctl *FileController) Get(c *fiber.Ctx) error {
	filename := c.Params("filename")
	if len(filename) == 0 {
		return response.Error(c, http.StatusBadRequest, errors.New("empty filename"))
	}

	file, err := ctl.service.FileStorage().Download(c.UserContext(), filename)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	c.Set(fiber.HeaderContentType, file.ContentType)
	return c.Status(http.StatusOK).SendStream(file.Content)
}

func (ctl *FileController) Upload(c *fiber.Ctx) error {
	src, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	filename, err := ctl.uploadFile(c, src)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"filename": filename,
	})
}

func (ctl *FileController) Uploads(c *fiber.Ctx) error {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	var filenames []string
	fileHeaders := multipartForm.File["files"]
	for _, fileHeader := range fileHeaders {
		filename, err := ctl.uploadFile(c, fileHeader)
		if err != nil {
			return err
		}
		filenames = append(filenames, filename)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"filenames": filenames,
	})
}

func (ctl *FileController) uploadFile(c *fiber.Ctx, fileHeader *multipart.FileHeader) (string, error) {
	content, err := fileHeader.Open()
	if err != nil {
		return "", response.Error(c, http.StatusBadRequest, err)
	}
	defer content.Close()

	raw, err := io.ReadAll(content)
	if err != nil {
		return "", response.Error(c, http.StatusInternalServerError, err)
	}

	file := &minio.File{
		Content:     bytes.NewBuffer(raw),
		Name:        fileHeader.Filename,
		Size:        fileHeader.Size,
		ContentType: http.DetectContentType(raw),
	}
	filename, err := ctl.service.FileStorage().Upload(c.UserContext(), file)
	if err != nil {
		return "", response.Error(c, http.StatusInternalServerError, err)
	}

	return filename, nil
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/pkg/minio"
	"mime/multipart"
	"net/http"
)

type FileController struct {
	service container.IService
}

func NewFileController(service container.IService) *FileController {
	return &FileController{service: service}
}

func (ctl *FileController) Upload(c *fiber.Ctx) error {
	src, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return "", c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer content.Close()

	contentType := c.GetReqHeaders()["Content-Type"]
	file := &minio.File{
		Content:     content,
		Name:        fileHeader.Filename,
		Size:        fileHeader.Size,
		ContentType: contentType,
	}

	filename, err := ctl.service.FileStorage().Upload(c.UserContext(), file)
	if err != nil {
		return "", c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return filename, nil
}

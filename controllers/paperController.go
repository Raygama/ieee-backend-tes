package controllers

import (
	"backend/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaperController struct {
	DB *gorm.DB
}

// UploadPaper adalah fungsi untuk mengupload paper beserta file-nya.
// @Summary Upload Paper with File
// @Description Uploads a paper along with its file and saves them to the database.
// @Tags Papers
// @Accept multipart/form-data
// @Param judul formData string true "Judul paper"
// @Param deskripsi formData string true "Deskripsi paper"
// @Param abstrak formData string true "Abstrak paper"
// @Param link formData string true "Link paper"
// @Param file_paper formData file true "File paper"
// @Param author formData string true "Author paper"
// @Param tanggal_terbit formData string true "Tanggal terbit paper"
// @Produce json
// @Success 200 {object} models.Paper
// @Router /papers [post]
func (pc *PaperController) UploadPaper(ctx *gin.Context) {
	// Get the file from the form data
	file, err := ctx.FormFile("file_paper")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define the path where the file will be saved
	fileName := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join("uploads", fileName)

	// Save the file to the defined path
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Parse other form data
	judul := ctx.PostForm("judul")
	deskripsi := ctx.PostForm("deskripsi")
	abstrak := ctx.PostForm("abstrak")
	link := ctx.PostForm("link")
	author := ctx.PostForm("author")
	tanggalTerbit, err := time.Parse(time.RFC3339, ctx.PostForm("tanggal_terbit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tanggal_terbit format"})
		return
	}

	// Create Paper object
	paper := models.Paper{
		Judul:         judul,
		Deskripsi:     deskripsi,
		Abstrak:       abstrak,
		Link:          link,
		FilePaper:     fileName,
		Author:        author,
		TanggalTerbit: tanggalTerbit,
	}

	// Save paper to database
	if err := pc.DB.Create(&paper).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save paper to database"})
		return
	}

	// Response success
	ctx.JSON(http.StatusOK, paper)
}

// GetAllPapers adalah fungsi untuk mendapatkan semua paper dari database.
// @Summary Get All Papers
// @Description Retrieves all papers from the database.
// @Tags Papers
// @Produce json
// @Success 200 {array} models.Paper
// @Router /papers [get]
func (pc *PaperController) GetAllPapers(ctx *gin.Context) {
	var papers []models.Paper

	// Retrieve all papers from the database
	if err := pc.DB.Find(&papers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get papers from database"})
		return
	}

	// Response with papers
	ctx.JSON(http.StatusOK, papers)
}

// GetPaperFile adalah fungsi untuk mengambil file paper berdasarkan ID.
// @Summary Get Paper File
// @Description Retrieves the file of a paper by its ID.
// @Tags Papers
// @Param id path string true "Paper ID"
// @Produce octet-stream
// @Success 200 {file} octet-stream
// @Router /papers/file/:id [get]
func (pc *PaperController) GetPaperFile(ctx *gin.Context) {
	// Get paper ID from URL path parameter
	paperID := ctx.Param("id")

	// Retrieve paper from the database by its ID
	var paper models.Paper
	if err := pc.DB.Where("id = ?", paperID).First(&paper).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Paper not found"})
		return
	}

	// Define the file path
	filePath := filepath.Join("uploads", paper.FilePaper)

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Set the headers for the file transfer and return the file
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", paper.FilePaper))
	ctx.Header("Content-Type", "application/pdf")
	ctx.File(filePath)
}

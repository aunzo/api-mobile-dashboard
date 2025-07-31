package handlers

import (
	"encoding/csv"
	"log"
	"strings"

	"github.com/aunz/api-mobile-dashboard-golang/internal/firestore"
	"github.com/aunz/api-mobile-dashboard-golang/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

type BuildInfoHandler struct {
	client *firestore.ClientWrapper
}

func NewBuildInfoHandler(client *firestore.ClientWrapper) *BuildInfoHandler {
	return &BuildInfoHandler{
		client: client,
	}
}

// @Summary Get list of build info
// @Produce json
// @Success 200 {array} models.BuildInfo
// @Router /build-info [get]
func (h *BuildInfoHandler) BuildInfoList(c *gin.Context) {
	var buildInfos []models.BuildInfo

	iter := h.client.Client.Collection("build-infos").Documents(h.client.Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate: %v", err)
			c.JSON(500, gin.H{"error": "database error"})
			return
		}
		var buildInfo models.BuildInfo
		mapstructure.Decode(doc.Data(), &buildInfo)
		buildInfos = append(buildInfos, buildInfo)
	}

	c.JSON(200, buildInfos)
}

// @Summary Download build info as CSV
// @Produce text/csv
// @Success 200
// @Router /build-info/csv [get]
func (h *BuildInfoHandler) BuildInfoListCSV(c *gin.Context) {
	iter := h.client.Client.Collection("build-infos").Documents(h.client.Ctx)

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=build_info.csv")

	csvWriter := csv.NewWriter(c.Writer)
	defer csvWriter.Flush()

	headers := []string{
		"StartTime", "EndTime", "Duration", "GitBranch", "GitAuthor", "Scheme", "MachineModel",
		"CPU", "MemoryGB", "DiskTotal", "DiskAvailable", "FileChangeCount", "CompileFileCount",
	}
	csvWriter.Write(headers)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.String(500, "Error reading from database")
			return
		}
		var buildInfo models.BuildInfo
		mapstructure.Decode(doc.Data(), &buildInfo)

		row := []string{
			strings.TrimSpace(buildInfo.StartTime),
			strings.TrimSpace(buildInfo.EndTime),
			strings.TrimSpace(buildInfo.Duration),
			strings.TrimSpace(buildInfo.GitBranch),
			strings.TrimSpace(buildInfo.GitAuthor),
			strings.TrimSpace(buildInfo.Scheme),
			strings.TrimSpace(buildInfo.MachineModel),
			strings.TrimSpace(buildInfo.CPU),
			strings.TrimSpace(buildInfo.MemoryGB),
			strings.TrimSpace(buildInfo.DiskTotal),
			strings.TrimSpace(buildInfo.DiskAvailable),
			strings.TrimSpace(buildInfo.FileChangeCount),
			strings.TrimSpace(buildInfo.CompileFileCount),
		}
		csvWriter.Write(row)
	}
}

// @Summary Create new build info
// @Accept json
// @Produce json
// @Param buildInfo body models.BuildInfo true "Build Info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /build-info [post]
func (h *BuildInfoHandler) CreateBuildInfo(c *gin.Context) {
	var buildInfo models.BuildInfo
	if err := c.BindJSON(&buildInfo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	_, _, err := h.client.Client.Collection("build-infos").Add(h.client.Ctx, buildInfo)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		c.JSON(500, gin.H{"error": "Failed to store build info"})
		return
	}

	c.JSON(201, gin.H{"message": "Create build info success!"})
}

package Controllers

import (
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var guid = uuid.New().String()

	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	// @params: pageSize

	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer

	// Get parameters from the request
	pageIndexStr := context.Query("pageIndex")
	pageSizeStr := context.Query("pageSize")

	// Convert parameters to integers and validate
	pageIndex, err := strconv.ParseInt(pageIndexStr, 10, 64)
	if err != nil || pageIndex < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageIndex"})
		return
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil || pageSize < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize"})
		return
	}

	kols, error := Logic.GetKolLogic(pageIndex, pageSize)
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = err.Error()
		KolsVM.PageIndex = pageIndex
		KolsVM.PageSize = pageSize
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = pageIndex
	KolsVM.PageSize = pageSize
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}

// for create dummy data only
func InsertKolsController(context *gin.Context) {
	_ = Logic.GenerateDummyKOLs(100)
	context.JSON(http.StatusOK, map[string]string{
		"Status": "Insert data into db success",
	})
}

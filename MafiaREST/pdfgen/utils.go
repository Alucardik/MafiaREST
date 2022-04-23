package pdfgen

import (
	"MafiaREST/config"
	"MafiaREST/utils"
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func genDocNameFromObjectId(id primitive.ObjectID) string {
	return "MafiaService_" + id.String()[10:34] + "_stats"
}

func tableHeader(pdf *gofpdf.Fpdf, headerName string) {
	pdf.SetFont("Helvetica", "IB", _TEXT_SZ_BIG)
	pdf.CellFormat(_WIDTH_CELL_TITLE, _HEIGHT_CELL_STD, headerName, "0", 0, "", false, 0, "")
	pdf.Ln(10)
}

func tableNextRow(pdf *gofpdf.Fpdf) {
	pdf.Ln(_TABLE_ROW_INTERSPACE)
}

func tableKeyValue(pdf *gofpdf.Fpdf, key, value string, fullRow, isLink bool) {
	pdf.SetFont("Helvetica", "IU", _TEXT_SZ_REGULAR)
	pdf.CellFormat(_TABLE_KEY_WIDTH, _HEIGHT_CELL_STD, key+":", "0", 0, "", false, 0, "")
	pdf.CellFormat(_TABLE_KEY_VALUE_MARGIN, _HEIGHT_CELL_STD, "", "0", 0, "", false, 0, "")
	pdf.SetFont("Helvetica", "", _TEXT_SZ_REGULAR)
	if isLink {
		// neon blue for link
		pdf.SetTextColor(31, 81, 255)
	}
	if fullRow {
		pdf.CellFormat(_TABLE_WIDTH-_TABLE_KEY_WIDTH, _HEIGHT_CELL_STD, value, "0", 0, "", false, 0, "")
		tableNextRow(pdf)
	} else {
		pdf.CellFormat(_TABLE_VALUE_WIDTH, _HEIGHT_CELL_STD, value, "0", 0, "", false, 0, "")
	}
	pdf.SetTextColor(0, 0, 0)
}

func displayAvatar(pdf *gofpdf.Fpdf, url, uid string) bool {
	pdf.Rect(135, 40, 40, 40, "D")
	res, err := http.Get(url)
	utils.NotifyOnError("", err)
	if err != nil {
		return false
	}

	buf, err := ioutil.ReadAll(res.Body)
	utils.NotifyOnError("", err)
	if err != nil {
		return false
	}

	format := ""
	for i := len(url) - 1; i > 0; i-- {
		if url[i] != '.' {
			format = string(url[i]) + format
		} else {
			break
		}
	}

	filePath := config.TMP_FILE_PATH + "/avatar_" + uid + "." + format
	err = os.WriteFile(filePath, buf, 0666)
	utils.NotifyOnError("Failed to write pic", err)
	if err != nil {
		return false
	}
	defer os.Remove(filePath)

	if format != "jpg" && format != "png" && format != "gif" && format != "jpeg" {
		return false
	}

	pdf.ImageOptions(filePath, 135, 40, 40, 40, false, gofpdf.ImageOptions{ImageType: strings.ToUpper(format), ReadDpi: true}, 0, "")
	pdf.Text(150, 85, "Avatar")
	return true
}

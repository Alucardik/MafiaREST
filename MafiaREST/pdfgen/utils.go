package pdfgen

import (
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

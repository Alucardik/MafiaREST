package pdfgen

import (
	"MafiaREST/config"
	"MafiaREST/schemes"
	"github.com/jung-kurt/gofpdf"
	"strconv"
)

// GenReport generates player stats pdf report
func GenReport(user *schemes.User, stats *schemes.UserStats) (string, error) {
	docName := genDocNameFromObjectId(stats.UID)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetTitle(docName, true)
	pdf.SetFont("Helvetica", "B", _TEXT_SZ_COLOSSAL)
	pdf.SetLeftMargin(25)
	pdf.Cell(_HEIGHT_CELL_TITLE, _HEIGHT_CELL_TITLE, "Player's "+stats.UID.String()[10:34]+" report")
	pdf.Ln(25)

	// table background
	pdf.SetFillColor(250, 235, 215)
	pdf.Rect(25, 35, _TABLE_WIDTH, _TABLE_HEIGHT, "DF")

	tableHeader(pdf, "Player's profile")
	tableKeyValue(pdf, "Name", user.Name, true, false)
	tableKeyValue(pdf, "Sex", user.Sex.ToString(), true, false)
	// TODO: try to upload avatar locally and insert into pdf
	tableKeyValue(pdf, "Avatar url", user.Avatar, true, true)
	tableKeyValue(pdf, "E-mail", user.Email, true, false)
	pdf.Ln(5)

	tableHeader(pdf, "Statistics")
	tableKeyValue(pdf, "Sessions", strconv.FormatUint(stats.SessionCount, 10), true, false)
	tableKeyValue(pdf, "Wins", strconv.FormatUint(stats.Wins, 10), false, false)
	tableKeyValue(pdf, "Win Rate", strconv.FormatFloat(float64(stats.Wins)/float64(stats.SessionCount)*100, 'f', 2, 64)+"%", true, false)
	tableKeyValue(pdf, "Losses", strconv.FormatUint(stats.Losses, 10), false, false)
	tableKeyValue(pdf, "Loss Rate", strconv.FormatFloat(float64(stats.Losses)/float64(stats.SessionCount)*100, 'f', 2, 64)+"%", true, false)
	tableKeyValue(pdf, "Overall time", strconv.FormatUint(stats.TotalTime/60, 10)+" mins", true, false)

	filePath := config.TMP_FILE_PATH + "/" + docName + ".pdf"
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

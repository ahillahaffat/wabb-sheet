package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	service       *sheets.Service
	spreadsheetID string
}

func NewSheetsService(config *Config, credentialsFile string) (*SheetsService, error) {
	ctx := context.Background()

	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &SheetsService{
		service:       srv,
		spreadsheetID: config.SpreadsheetID,
	}, nil
}

func (s *SheetsService) SaveMessage(message string) error {
	ctx := context.Background()

	parsed := parseMessage(message)
	currentTime := time.Now().Format("02-01-2006 15:04:05")

	values := [][]interface{}{
		{
			currentTime,
			parsed["Tanggal"],
			parsed["Nopol"],
			parsed["Armada"],
			parsed["Sopir"],
			parsed["Pembeli"],
			parsed["BB TB"],
			parsed["Bobot"],
			parsed["Jenis"],
		},
	}

	_, err := s.service.Spreadsheets.Values.Append(
		s.spreadsheetID,
		"Sheet1!A:K",
		&sheets.ValueRange{Values: values},
	).ValueInputOption("USER_ENTERED").Context(ctx).Do()

	if err != nil {
		return err
	}

	log.Println("✅ Data successfully saved to Google Spreadsheet")
	return nil
}

func (s *SheetsService) StyleHeader() error {
	requests := []*sheets.Request{
		{
			RepeatCell: &sheets.RepeatCellRequest{
				Range: &sheets.GridRange{
					SheetId:          0,
					StartRowIndex:    3,
					EndRowIndex:      4,
					StartColumnIndex: 0,
					EndColumnIndex:   9,
				},
				Cell: &sheets.CellData{
					UserEnteredFormat: &sheets.CellFormat{
						BackgroundColor: &sheets.Color{
							Red:   0.3,
							Green: 0.5,
							Blue:  0.8,
						},
						HorizontalAlignment: "CENTER",
						TextFormat: &sheets.TextFormat{
							Bold: true,
						},
						Borders: &sheets.Borders{
							Top:    &sheets.Border{Style: "SOLID"},
							Bottom: &sheets.Border{Style: "SOLID"},
							Left:   &sheets.Border{Style: "SOLID"},
							Right:  &sheets.Border{Style: "SOLID"},
						},
					},
				},
				Fields: "userEnteredFormat(backgroundColor,textFormat,horizontalAlignment,borders)",
			},
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}

	_, err := s.service.Spreadsheets.BatchUpdate(s.spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		return err
	}

	log.Println("✅ Header successfully formatted!")
	return nil
}

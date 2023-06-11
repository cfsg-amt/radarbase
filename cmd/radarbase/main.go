package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"radarbase/pkg/api"
	"radarbase/pkg/excel"
	"radarbase/pkg/mdb"
)

// Helper function to load a single sheet into the database
func loadSheetIntoDB(db *mdb.MDB, filePath string, sheetName string) error {
	// Parse row data
	rowData, _, err := excel.RowParse(filePath, sheetName)
	if err != nil {
		return err
	}

	// Parse columnar data
	colData, colHeaders, err := excel.ColParse(filePath, sheetName)
	if err != nil {
		return err
	}

	// Inserting the row data
	err = db.RowLoadToDB(rowData, sheetName)
	if err != nil {
		return err
	}

	// Inserting the columnar data
	err = db.ColLoadToDB(colData, colHeaders, sheetName)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := mdb.NewMDB("mongodb://localhost:27017", "RowDB", "ColDB")
	if err != nil {
		log.Fatal(err)
	}

	// List all the sheet names that you need to load
	sheets := []string{"CSV-Sec", "CSV-Ind", "CSV-StkSH", "CSV-StkSZ", "CSV-StkHK"}
	filePath := "files/radar.xlsm"

	for _, sheet := range sheets {
		err = loadSheetIntoDB(db, filePath, sheet)
		if err != nil {
			fmt.Printf("Failed to load sheet %s: %v\n", sheet, err)
			os.Exit(1)
		}
	}

	// Initialize API
	apiHandler := api.NewAPI(db)

	// Start server
	srv := &http.Server{
		Handler: apiHandler.SetupRouter(), // use the SetupRouter method here
		Addr:    "127.0.0.1:8996",
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// It's fine to panic here, as this should never happen when closing the server
			log.Panic(err)
		}
	}()

	// Wait for Ctrl+C to exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received
	<-c

	// Cleanup and shutdown
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// Extra handling here
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Shutting down gracefully, bye bye ...")
	os.Exit(0)
}

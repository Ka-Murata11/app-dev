package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jszwec/csvutil"
	"github.com/labstack/echo/v4"
)

type Record struct {
	Column1 string `csv:"項目１"`
	Column2 string `csv:"項目２"`
}

func ImportCSV(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.ErrBadRequest
	}

	src, err := file.Open()
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer src.Close()

	// ここでCSVファイルを読み込んでDBにインポートする処理を書く
	// reader := csv.NewReader(src)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	return echo.ErrInternalServerError
	// }

	var data []Record

	// csvutilを使ってCSVファイルを構造体に変換
	// ライブラリインストール go get -u github.com/jszwec/csvutil
	records, err := io.ReadAll(src)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if err := csvutil.Unmarshal(records, &data); err != nil {
		return echo.ErrInternalServerError
	}

	// gocsvを使ってCSVファイルを構造体に変換
	// ライブラリインストール go get -u github.com/gocarina/gocsv
	// if err := gocsv.Unmarshal(src, &data); err != nil {
	// 	return echo.ErrInternalServerError
	// }

	fmt.Println(data)

	return c.JSON(http.StatusOK, "Import success")
}

// CSVファイルのデータを検証する関数
// func validateData(data []Record) error {
// 	for _, record := range data {
// 		if err := validateRecord(record); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func validateRecord(record Record) error {
// 	v := reflect.ValueOf(record)
// 	for i := 0; i < v.NumField(); i++ {
// 		if v.Field(i).Interface() == "" {
// 			return fmt.Errorf("CSV contains empty fields")
// 		}
// 	}
// 	return nil
// }

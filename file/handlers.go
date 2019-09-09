package file

import (
	"encoding/json"
	"math"

	"github.com/danskeren/imgasm.com/db"
	"github.com/danskeren/imgasm.com/models"
	"github.com/dgraph-io/badger"
	"github.com/packago/generate"
)

func calculateFitDimension(imageWidth, imageHeight, fitWidth, fitHeight int) (int, int) {
	if imageWidth*fitHeight > fitWidth*imageHeight {
		// constrained by width
		fitHeight = int(math.Round(float64(fitWidth) * float64(imageHeight) / float64(imageWidth)))
	} else {
		// constrained by height
		fitWidth = int(math.Round(float64(fitHeight) * float64(imageWidth) / float64(imageHeight)))
	}
	return fitWidth, fitHeight
}

func generateAvailableFileDataID() string {
	for {
		id := generate.RandomLowercaseNumbers(6)
		if _, err := db.BadgerDB.Get([]byte(id)); err != badger.ErrKeyNotFound {
			continue
		}
		return id
	}
}

func saveFileData(filename string) (models.FileData, error) {
	fileData := models.FileData{
		ID:             generateAvailableFileDataID(),
		Filename:       filename,
		DeletePassword: generate.RandomLettersNumbers(12),
	}
	fileDataBytes, err := json.Marshal(fileData)
	if err != nil {
		return fileData, err
	}
	if err = db.BadgerDB.Set([]byte(fileData.ID), fileDataBytes); err != nil {
		return fileData, err
	}
	return fileData, nil
}

func saveMD5Hashes(original, processed, extension string) error {
	if original != processed {
		fileMD5 := models.FileMD5{
			Processed: processed,
		}
		fileMD5Bytes, err := json.Marshal(fileMD5)
		if err != nil {
			return err
		}
		if err = db.BadgerDB.Set([]byte(original), fileMD5Bytes); err != nil {
			return err
		}
	}
	fileMD5 := models.FileMD5{
		Processed: processed,
		Extension: extension,
	}
	fileMD5Bytes, err := json.Marshal(fileMD5)
	if err != nil {
		return err
	}
	if err = db.BadgerDB.Set([]byte(processed), fileMD5Bytes); err != nil {
		return err
	}
	return nil
}

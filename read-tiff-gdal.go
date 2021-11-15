//read-tiff-gdal.go
//package gdal_utils
package main

import (
	"fmt"
	//"github.com/Rob-Fletcher/go-gdal"
	//"github.com/boundlessgeo/go-gdal"
	//"github.com/pkg/errors"
	//log "github.com/sirupsen/logrus"
)

const TIF_PATH = "/Users/venu4461/Downloads/feature_categorization_using_satellite_imagery_and_deep_learning/images/000000000.tif"

func GetDriverInfo() {

	fmt.Println("GetDriverInfo")
	//var xsize uint
	//var ysize uint

	//bmap : = [3]int {1, 2, 3}

	/*buffer := make([]uint8, xsize*ysize*len(bmap))

	ds, err := gdal.Open(inFile, gdal.ReadOnly)
	defer ds.Close()
	FailOnError(err, "Failed to read file.")

	bandCount := ds.RasterCount()
	if err != nil{
		fmt.Println("error")
	}*/

	/*fmt.Printf("%d drivers available\n", gdal.GetDriverCount())
	for x := 0; x < gdal.GetDriverCount(); x++ {
		driver := gdal.GetDriver(x)
		fmt.Printf("%s: %s\n", driver.ShortName(), driver.LongName())
	}*/

}

package method

import (
	"strconv"
	"fmt"
	"net/url"
	"github.com/json-iterator/go"
)

var (
	Addr = "http://yingyan.baidu.com/api/v3/track/addpoints"
	Ak = "Wy6r3GK9HSaGITWG71K9GnixXRl2K4oy"
	ServiceId = "204218"
	EntityName = "gaojie"
	CoordTypeInput = "wgs84"
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Point struct {
	Longitude		float64	`json:"longitude"`
	Latitude		float64	`json:"latitude"`
	LocTime			int64	`json:"loc_time"`
	EntityName		string	`json:"entity_name"`
	CoordTypeInput	string	`json:"coord_type_input"`
}

func ToLocation(longitude,latitude string) (float64,float64) {
	latDStr := latitude[:2]
	latMStr := latitude[2:]
	lngDStr := longitude[:3]
	lngMStr := longitude[3:]
	latd,err := strconv.ParseFloat(latDStr,32)
	latm,err := strconv.ParseFloat(latMStr,32)
	lngd,err := strconv.ParseFloat(lngDStr,32)
	lngm,err := strconv.ParseFloat(lngMStr,32)
	if err != nil {
		fmt.Println(latitude,longitude,err)
		return 0,0
	}
	return lngd+lngm/60,latd+latm/60
}

func ToParams(points []*Point) {
	fmt.Println(len(points))
	//单个坐标点
	//lng,lat := method.ToLocation("11628.3763","3959.3055")
	//points = append(points,&method.Point{Longitude:lng,Latitude:lat,LocTime:time1,EntityName:entity_name,CoordTypeInput:coord_type_input})
	pointList,err := json.MarshalToString(points)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := url.Values{}
	body.Set("ak",Ak)
	body.Set("service_id",ServiceId)
	body.Set("point_list",pointList)
	fmt.Println(string(ReqForm(Addr,body)))
}



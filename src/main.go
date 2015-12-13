package main

import (
	"api.com/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	//"time"
)

type API struct {
	Id      int    "json:id"
	Message string "json:message"
	Code    string "json:code"
	Label   string "json:label"
	Status  int    "json:status"
	Price   string "json:price"
}

func in_area(lat float64, lng float64, points []*models.Point) bool {
	flag := false
	on := false
	j := len(points) - 1
	for i := 0; i < len(points); i++ {
		sx := points[i].Lat
		sy := points[i].Lng
		tx := points[j].Lat
		ty := points[j].Lng

		if (sy < lng && ty >= lng) || (sy >= lng && ty < lng) {
			x := sx + (lng-sy)*(tx-sx)/(ty-sy)
			if x == lat {
				on = true
				break
			}

			if x > lat {
				flag = !flag
			}
		}
		j = i
		//fmt.Println(points[i].lat, points[i].lng)
	}

	if on == true {
		return on
	}
	return flag
}

func new_in_area(lat float64, lng float64, area_arr []*models.Area) bool {
	result := false
	for i := 0; i < len(area_arr); i++ {
		result = in_area(lat, lng, area_arr[i].Points)
		if result == true {
			break
		}
	}
	return result
}
func select_points(db *sql.DB, area_id int) []*models.Point {
	var (
		id  int
		lat float64
		lng float64
	)
	sql := "select id, lantitude, longitude from  points p  where p.pointable_id =? and p.pointable_type = 'Area'"
	rows, err := db.Query(sql, area_id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	points := []*models.Point{}
	for rows.Next() {
		err := rows.Scan(&id, &lat, &lng)
		if err != nil {
			log.Fatal(err)
		}

		point := &models.Point{Id: id, Lat: lat, Lng: lng}
		points = append(points, point)
		//log.Println(id,lantitude,longitude)
	}
	return points
}

func select_areas(db *sql.DB, station_id int) []*models.Area {
	var (
		id        int
		latitude  float64
		longitude float64
		label     string
		code      string
		price     []byte
	)
	sql := "select a.id, latitude, longitude, label, code, c.price from  areas a left join commissions c on a.commission_id = c.id  where a.station_id =?"
	rows, err := db.Query(sql, station_id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	areas := []*models.Area{}
	for rows.Next() {
		err := rows.Scan(&id, &latitude, &longitude, &label, &code, &price)
		if err != nil {
			log.Fatal(err)
		}

		area := &models.Area{Id: id, Latitude: latitude, Longitude: longitude, Label: label, Code: code, Price: string(price)}
		areas = append(areas, area)
		//log.Println(id,lantitude,longitude)
	}
	return areas
}

func select_area(db *sql.DB, station_id int) map[int][]*models.Point {
	sql := "select s.id, a.id, p.id, p.lantitude, p.longitude from stations s, areas a, points p where a.station_id = s.id and p.pointable_id = a.id and p.pointable_type='Area' and s.id=?"
	rows, err := db.Query(sql, station_id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var se_station_id int
	var area_id int
	var id int
	var lantitude float64
	var longitude float64
	var temp_area_id int
	points := []*models.Point{}
	area_map := map[int][]*models.Point{}
	for rows.Next() {
		err := rows.Scan(&se_station_id, &area_id, &id, &lantitude, &longitude)
		if err != nil {
			log.Fatal(err)
		}
		if temp_area_id != area_id {
			points = []*models.Point{}
			temp_area_id = area_id
		}
		point := &models.Point{Lat: lantitude, Lng: longitude}
		points = append(points, point)
		area_map[area_id] = points
		//log.Println(id,lantitude,longitude)
	}
	return area_map

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/areas/price", PriceHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	//urlParams := mux.Vars(r)
	station_id, err := strconv.Atoi(r.FormValue("station_id"))
	lat, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	lng, err := strconv.ParseFloat(r.FormValue("longitude"), 64)

	db, err := sql.Open("mysql", "root:22143521@/price_development")
	//defer db.Close()
	//area_map := select_area(db, 90)
	result := false
	//var area_id int
	var area models.Area
	areas := select_areas(db, station_id)
	for _, item := range areas {
		points := select_points(db, item.Id)
		result = in_area(lat, lng, points)
		if result == true {
			area = *item
			fmt.Println(item.Id)
			break
		}
		//fmt.Println(item.Id)
	}
	/**
	  for k, v := range area_map {
	    //str := k + ":" + v.lat + " " + v.lng
	    //fmt.Println(k)
	    result = in_area(lat, lng, v)
	    if result == true {
	      area_id = k
	      break
	    }

	  }**/
	message := API{area.Id, "success", area.Code, area.Label, 0, area.Price}
	output, err := json.Marshal(message)
	//test_in()
	if err != nil {
		fmt.Println("something went wrong!")
	}
	fmt.Fprintf(w, string(output))

}

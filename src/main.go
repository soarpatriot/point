package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type API struct {
	Message string "json:message"
	AreaId  int    "json:area_id"
	Status  bool   "json:status"
}

type Point struct {
	lat float64
	lng float64
}
type Area struct {
	id     int
	points []*Point
}

func in_area(lat float64, lng float64, points []*Point) bool {
	flag := false
	on := false
	j := len(points) - 1
	for i := 0; i < len(points); i++ {
		sx := points[i].lat
		sy := points[i].lng
		tx := points[j].lat
		ty := points[j].lng

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

func test_in() {
	p0 := &Point{39.809608, 116.640264}
	p1 := &Point{39.831719, 116.592115}
	p2 := &Point{39.814374, 116.577239}
	p3 := &Point{39.804619, 116.582485}
	p4 := &Point{39.804508, 116.582557}
	p5 := &Point{39.799464, 116.585647}
	p6 := &Point{39.795306, 116.589024}
	p7 := &Point{39.781667, 116.60196}
	p8 := &Point{39.781722, 116.611949}
	p9 := &Point{39.786047, 116.640695}
	p10 := &Point{39.78638, 116.649606}
	p11 := &Point{39.776233, 116.661536}
	p12 := &Point{39.764143, 116.679933}
	p13 := &Point{39.764586, 116.680508}
	p14 := &Point{39.764254, 116.680364}
	p15 := &Point{39.761259, 116.736706}
	p16 := &Point{39.774791, 116.737281}
	p17 := &Point{39.81382, 116.738143}
	p18 := &Point{39.817755, 116.679645}

	q0 := &Point{39.680095, 116.777886}
	q1 := &Point{39.682538, 116.860099}
	q2 := &Point{39.686981, 116.863549}
	q3 := &Point{39.693199, 116.885683}
	q4 := &Point{39.705635, 116.885108}
	q5 := &Point{39.732941, 116.890857}
	q6 := &Point{39.748476, 116.89172}
	q7 := &Point{39.785302, 116.889707}
	q8 := &Point{39.80659, 116.888558}
	q9 := &Point{39.805925, 116.914429}
	q10 := &Point{39.805925, 116.940013}
	q11 := &Point{39.811911, 116.9426}
	q12 := &Point{39.817454, 116.935126}
	q13 := &Point{39.82366, 116.933688}
	q14 := &Point{39.836072, 116.933401}
	q15 := &Point{39.847152, 116.929089}
	q16 := &Point{39.854907, 116.920465}
	q17 := &Point{39.857344, 116.912704}
	q18 := &Point{39.854021, 116.909542}
	q19 := &Point{39.846266, 116.910404}
	q20 := &Point{39.841169, 116.913566}
	q21 := &Point{39.83718, 116.912129}
	q22 := &Point{39.838067, 116.90638}
	q23 := &Point{39.84538, 116.893732}
	q24 := &Point{39.853135, 116.871598}
	q25 := &Point{39.860003, 116.852913}
	q26 := &Point{39.86244, 116.849751}
	q27 := &Point{39.873485, 116.840909}
	q28 := &Point{39.883451, 116.829123}
	q29 := &Point{39.883008, 116.815613}
	q30 := &Point{39.886552, 116.80124}
	q31 := &Point{39.895188, 116.791466}
	q32 := &Point{39.85687, 116.779806}
	q33 := &Point{39.852439, 116.774794}
	q34 := &Point{39.85111, 116.769745}
	q35 := &Point{39.84424, 116.763421}
	q36 := &Point{39.833825, 116.751636}
	q37 := &Point{39.813806, 116.738125}
	q38 := &Point{39.761148, 116.735987}
	q39 := &Point{39.762787, 116.707962}
	q40 := &Point{39.764257, 116.679971}
	q41 := &Point{39.762538, 116.682163}
	q42 := &Point{39.762108, 116.680941}
	q43 := &Point{39.761428, 116.679791}
	q44 := &Point{39.761227, 116.679288}
	q45 := &Point{39.761061, 116.678713}
	q46 := &Point{39.760666, 116.675641}
	q47 := &Point{39.760208, 116.67212}
	q48 := &Point{39.760042, 116.669874}
	q49 := &Point{39.759792, 116.667017}
	q50 := &Point{39.759078, 116.66505}
	q51 := &Point{39.75697, 116.659526}
	q52 := &Point{39.755077, 116.659723}
	q53 := &Point{39.751097, 116.660226}
	q54 := &Point{39.743885, 116.661053}
	q55 := &Point{39.739183, 116.661556}
	q56 := &Point{39.736201, 116.661915}
	q57 := &Point{39.731942, 116.663352}
	q58 := &Point{39.724825, 116.665832}
	q59 := &Point{39.723785, 116.665886}
	q60 := &Point{39.719803, 116.66594}
	q61 := &Point{39.715571, 116.66594}
	q62 := &Point{39.707897, 116.666065}
	q63 := &Point{39.702568, 116.666137}
	q64 := &Point{39.699792, 116.666191}
	q65 := &Point{39.696655, 116.665993}
	q66 := &Point{39.694046, 116.665814}
	q67 := &Point{39.691187, 116.665634}
	q68 := &Point{39.688244, 116.665652}
	q69 := &Point{39.688077, 116.665778}
	q70 := &Point{39.687605, 116.666207}
	q71 := &Point{39.686183, 116.667527}
	q72 := &Point{39.683385, 116.67015}
	q73 := &Point{39.677943, 116.675226}
	q74 := &Point{39.674465, 116.679295}
	q75 := &Point{39.670967, 116.683302}
	q76 := &Point{39.669203, 116.685341}
	q77 := &Point{39.666648, 116.688287}
	q78 := &Point{39.662187, 116.693434}
	q79 := &Point{39.65899, 116.700068}
	q80 := &Point{39.658226, 116.701676}
	q81 := &Point{39.657191, 116.703823}
	q82 := &Point{39.655504, 116.707327}
	q83 := &Point{39.652237, 116.71414}
	q84 := &Point{39.64899, 116.720927}
	q85 := &Point{39.652056, 116.726663}
	q86 := &Point{39.653959, 116.73022}
	q87 := &Point{39.654886, 116.731976}
	q88 := &Point{39.655528, 116.733157}

	r0 := &Point{39.843603, 116.694274}
	r1 := &Point{39.843306, 116.696457}
	r2 := &Point{39.842752, 116.700518}
	r3 := &Point{39.842402, 116.706478}
	r4 := &Point{39.848039, 116.719485}
	r5 := &Point{39.851688, 116.727804}
	r6 := &Point{39.852809, 116.729771}
	r7 := &Point{39.854288, 116.732358}
	r8 := &Point{39.855873, 116.735156}
	r9 := &Point{39.856981, 116.737092}
	r10 := &Point{39.858373, 116.739823}
	r11 := &Point{39.861221, 116.745437}
	r12 := &Point{39.86371, 116.75027}
	r13 := &Point{39.909199, 116.748568}
	r14 := &Point{39.909271, 116.746461}
	r15 := &Point{39.909344, 116.744314}
	r16 := &Point{39.909223, 116.740708}
	r17 := &Point{39.909081, 116.737038}
	r18 := &Point{39.908963, 116.733521}
	r19 := &Point{39.908711, 116.726726}
	r20 := &Point{39.908576, 116.722796}
	r21 := &Point{39.908434, 116.718856}
	r22 := &Point{39.907822, 116.716135}
	r23 := &Point{39.906923, 116.715021}
	r24 := &Point{39.905798, 116.713601}
	r25 := &Point{39.904442, 116.71189}
	r26 := &Point{39.902982, 116.710053}
	r27 := &Point{39.900298, 116.706675}
	r28 := &Point{39.900564, 116.706316}
	r29 := &Point{39.901983, 116.704367}
	r30 := &Point{39.903443, 116.702386}
	r31 := &Point{39.905857, 116.699062}
	r32 := &Point{39.905359, 116.696156}
	r33 := &Point{39.904688, 116.69219}
	r34 := &Point{39.903934, 116.687824}
	r35 := &Point{39.903138, 116.683225}
	r36 := &Point{39.902626, 116.680193}
	r37 := &Point{39.902792, 116.676057}
	r38 := &Point{39.903103, 116.668762}
	r39 := &Point{39.903702, 116.655041}
	r40 := &Point{39.90382, 116.652216}
	r41 := &Point{39.903799, 116.651448}
	r42 := &Point{39.903747, 116.648335}
	r43 := &Point{39.905525, 116.647145}
	r44 := &Point{39.908168, 116.645384}
	r45 := &Point{39.910008, 116.64418}
	r46 := &Point{39.911184, 116.643111}
	r47 := &Point{39.911956, 116.642262}
	r48 := &Point{39.91263, 116.640039}
	r49 := &Point{39.912692, 116.635296}
	r50 := &Point{39.91273, 116.631442}
	r51 := &Point{39.912761, 116.628918}
	r52 := &Point{39.912654, 116.6263}
	r53 := &Point{39.912032, 116.620941}
	r54 := &Point{39.913197, 116.618269}
	r55 := &Point{39.912319, 116.618866}
	r56 := &Point{39.911053, 116.619719}
	r57 := &Point{39.908573, 116.621426}
	r58 := &Point{39.906833, 116.622616}
	r59 := &Point{39.905096, 116.623802}
	r60 := &Point{39.902854, 116.625325}
	r61 := &Point{39.901664, 116.625904}
	r62 := &Point{39.89945, 116.62696}
	r63 := &Point{39.897506, 116.627903}
	r64 := &Point{39.893711, 116.629727}
	r65 := &Point{39.890431, 116.631294}
	r66 := &Point{39.887386, 116.632754}
	r67 := &Point{39.88632, 116.63248}
	r68 := &Point{39.884818, 116.632098}
	r69 := &Point{39.882351, 116.63146}
	r70 := &Point{39.880451, 116.630957}
	r71 := &Point{39.879887, 116.630674}
	r72 := &Point{39.877257, 116.629345}
	r73 := &Point{39.875675, 116.629655}
	r74 := &Point{39.873287, 116.630108}
	r75 := &Point{39.871231, 116.63173}
	r76 := &Point{39.8702, 116.634115}
	r77 := &Point{39.869667, 116.63707}
	r78 := &Point{39.868186, 116.635103}
	r79 := &Point{39.866791, 116.633243}
	r80 := &Point{39.863949, 116.629493}
	r81 := &Point{39.862952, 116.631299}
	r82 := &Point{39.861955, 116.633118}
	r83 := &Point{39.860176, 116.636325}
	r84 := &Point{39.856642, 116.642748}
	r85 := &Point{39.855378, 116.645694}
	r86 := &Point{39.854748, 116.647154}
	r87 := &Point{39.854115, 116.648609}
	r88 := &Point{39.853505, 116.650028}
	r89 := &Point{39.85174, 116.654116}
	r90 := &Point{39.850756, 116.657835}
	r91 := &Point{39.849887, 116.661122}
	r92 := &Point{39.848901, 116.664801}
	r93 := &Point{39.847135, 116.674512}
	r94 := &Point{39.846719, 116.676784}
	r95 := &Point{39.846121, 116.680122}
	r96 := &Point{39.845736, 116.6823}
	r97 := &Point{39.845421, 116.684052}
	r98 := &Point{39.845089, 116.685772}
	r99 := &Point{39.84458, 116.688453}

	s0 := &Point{39.863949, 116.629489}
	s1 := &Point{39.856787, 116.6198}
	s2 := &Point{39.849586, 116.611967}
	s3 := &Point{39.841332, 116.601403}
	s4 := &Point{39.831858, 116.591486}
	s5 := &Point{39.809622, 116.640309}
	s6 := &Point{39.813633, 116.659568}
	s7 := &Point{39.815607, 116.669288}
	s8 := &Point{39.816646, 116.674274}
	s9 := &Point{39.817457, 116.678208}
	s10 := &Point{39.817776, 116.679636}
	s11 := &Point{39.817561, 116.682403}
	s12 := &Point{39.817318, 116.686392}
	s13 := &Point{39.816688, 116.695932}
	s14 := &Point{39.816065, 116.705364}
	s15 := &Point{39.814963, 116.72165}
	s16 := &Point{39.81443, 116.72978}
	s17 := &Point{39.813868, 116.738107}
	s18 := &Point{39.817485, 116.740551}
	s19 := &Point{39.820228, 116.742401}
	s20 := &Point{39.821156, 116.74303}
	s21 := &Point{39.82782, 116.74753}
	s22 := &Point{39.831207, 116.749776}
	s23 := &Point{39.8344, 116.752157}
	s24 := &Point{39.838783, 116.757169}
	s25 := &Point{39.840972, 116.759693}
	s26 := &Point{39.843195, 116.762209}
	s27 := &Point{39.848568, 116.767491}
	s28 := &Point{39.850486, 116.769153}
	s29 := &Point{39.851075, 116.769692}
	s30 := &Point{39.851539, 116.771264}
	s31 := &Point{39.852155, 116.773671}
	s32 := &Point{39.852439, 116.774713}
	s33 := &Point{39.853076, 116.775486}
	s34 := &Point{39.855118, 116.777767}
	s35 := &Point{39.856836, 116.779753}
	s36 := &Point{39.857313, 116.779905}
	s37 := &Point{39.876267, 116.785672}
	s38 := &Point{39.885728, 116.788547}
	s39 := &Point{39.89561, 116.791547}
	s40 := &Point{39.908299, 116.774884}
	s41 := &Point{39.909206, 116.748559}
	s42 := &Point{39.906805, 116.748676}
	s43 := &Point{39.904121, 116.748757}
	s44 := &Point{39.8999, 116.748918}
	s45 := &Point{39.891472, 116.749233}
	s46 := &Point{39.88705, 116.749385}
	s47 := &Point{39.882759, 116.749785}
	s48 := &Point{39.873201, 116.749915}
	s49 := &Point{39.863696, 116.75027}
	s50 := &Point{39.856995, 116.737119}
	s51 := &Point{39.851677, 116.727777}
	s52 := &Point{39.849018, 116.72174}
	s53 := &Point{39.842398, 116.706451}
	s54 := &Point{39.842745, 116.700558}
	s55 := &Point{39.843617, 116.69418}
	s56 := &Point{39.844601, 116.688377}
	s57 := &Point{39.845404, 116.684155}
	s58 := &Point{39.846733, 116.676681}
	s59 := &Point{39.848977, 116.664536}
	s60 := &Point{39.851747, 116.654098}
	s61 := &Point{39.856649, 116.642743}

	points := []*Point{p0, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18}
	qpoints := []*Point{q0, q1, q2, q3, q4, q5, q6, q7, q8, q9, q10, q11, q12, q13, q14, q15, q16, q17, q18, q19, q20, q21, q22, q23, q24, q25, q26, q27, q28, q29, q30, q31, q32, q33, q34, q35, q36, q37, q38, q39, q40, q41, q42, q43, q44, q45, q46, q47, q48, q49, q50, q51, q52, q53, q54, q55, q56, q57, q58, q59, q60, q61, q62, q63, q64, q65, q66, q67, q68, q69, q70, q71, q72, q73, q74, q75, q76, q77, q78, q79, q80, q81, q82, q83, q84, q85, q86, q87, q88}
	rpoints := []*Point{r0, r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, r32, r33, r34, r35, r36, r37, r38, r39, r40, r41, r42, r43, r44, r45, r46, r47, r48, r49, r50, r51, r52, r53, r54, r55, r56, r57, r58, r59, r60, r61, r62, r63, r64, r65, r66, r67, r68, r69, r70, r71, r72, r73, r74, r75, r76, r77, r78, r79, r80, r81, r82, r83, r84, r85, r86, r87, r88, r89, r90, r91, r92, r93, r94, r95, r96, r97, r98, r99}
	spoints := []*Point{s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31, s32, s33, s34, s35, s36, s37, s38, s39, s40, s41, s42, s43, s44, s45, s46, s47, s48, s49, s50, s51, s52, s53, s54, s55, s56, s57, s58, s59, s60, s61}

	area := new(Area)
	area.points = points
	areaq := &Area{1, qpoints}
	arear := &Area{2, rpoints}
	areas := &Area{3, spoints}

	area_arr := []*Area{area, areaq, arear, areas}

	t1 := time.Now()
	lat := 39.809815
	lng := 116.740079
	result := new_in_area(lat, lng, area_arr)
	fmt.Println(time.Now().Sub(t1))
	fmt.Println("hello,world! ", result)

}
func new_in_area(lat float64, lng float64, area_arr []*Area) bool {
	result := false
	for i := 0; i < len(area_arr); i++ {
		result = in_area(lat, lng, area_arr[i].points)
		if result == true {
			break
		}
	}
	return result
}
func select_area(db *sql.DB, station_id int) map[int][]*Point {
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
	points := []*Point{}
	area_map := map[int][]*Point{}
	for rows.Next() {
		err := rows.Scan(&se_station_id, &area_id, &id, &lantitude, &longitude)
		if err != nil {
			log.Fatal(err)
		}
		if temp_area_id != area_id {
			points = []*Point{}
			temp_area_id = area_id
		}
		point := &Point{lantitude, longitude}
		points = append(points, point)
		area_map[area_id] = points
		//log.Println(id,lantitude,longitude)
	}
	return area_map

}
func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", "root:22143521@/price_development")
		defer db.Close()
		area_map := select_area(db, 99)
		result := false
		var area_id int
		lat := 39.809815
		lng := 116.740079

		for k, v := range area_map {
			//str := k + ":" + v.lat + " " + v.lng
			//fmt.Println(k)
			result = in_area(lat, lng, v)
			if result == true {
				area_id = k
				break
			}

		}
		message := API{"Hello", area_id, result}
		output, err := json.Marshal(message)
		//test_in()
		if err != nil {
			fmt.Println("something went wrong!")
		}
		fmt.Fprintf(w, string(output))

	})

	http.ListenAndServe(":8080", nil)
}

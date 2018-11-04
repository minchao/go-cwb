package cwb

import "errors"

var (
	locations = []Location{
		{
			"宜蘭縣",
			[]DataSet{
				{FTW2DayYilanCounty, DataSetTypeFTW},
				{FTW7DayYilanCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"桃園市",
			[]DataSet{
				{FTW2DayTaoyuanCity, DataSetTypeFTW},
				{FTW7DayTaoyuanCity, DataSetTypeFTW7Day},
			},
		},
		{
			"新竹縣",
			[]DataSet{
				{FTW2DayHsinchuCounty, DataSetTypeFTW},
				{FTW7DayHsinchuCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"苗栗縣",
			[]DataSet{
				{FTW2DayMiaoliCounty, DataSetTypeFTW},
				{FTW7DayMiaoliCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"彰化縣",
			[]DataSet{
				{FTW2DayChanghuaCounty, DataSetTypeFTW},
				{FTW7DayChanghuaCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"南投縣",
			[]DataSet{
				{FTW2DayNantouCounty, DataSetTypeFTW},
				{FTW7DayNantouCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"雲林縣",
			[]DataSet{
				{FTW2DayYunlinCounty, DataSetTypeFTW},
				{FTW7DayYunlinCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"嘉義縣",
			[]DataSet{
				{FTW2DayChiayiCounty, DataSetTypeFTW},
				{FTW7DayChiayiCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"屏東縣",
			[]DataSet{
				{FTW2DayPingtungCounty, DataSetTypeFTW},
				{FTW7DayPingtungCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"臺東縣",
			[]DataSet{
				{FTW2DayTaitungCounty, DataSetTypeFTW},
				{FTW7DayTaitungCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"花蓮縣",
			[]DataSet{
				{FTW2DayHualienCounty, DataSetTypeFTW},
				{FTW7DayHualienCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"澎湖縣",
			[]DataSet{
				{FTW2DayPenghuCounty, DataSetTypeFTW},
				{FTW7DayPenghuCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"基隆市",
			[]DataSet{
				{FTW2DayKeelungCity, DataSetTypeFTW},
				{FTW7DayKeelungCity, DataSetTypeFTW7Day},
			},
		},
		{
			"新竹市",
			[]DataSet{
				{FTW2DayHsinchuCity, DataSetTypeFTW},
				{FTW7DayHsinchuCity, DataSetTypeFTW7Day},
			},
		},
		{
			"嘉義市",
			[]DataSet{
				{FTW2DayChiayiCity, DataSetTypeFTW},
				{FTW7DayChiayiCity, DataSetTypeFTW7Day},
			},
		},
		{
			"臺北市",
			[]DataSet{
				{FTW2DayTaipeiCity, DataSetTypeFTW},
				{FTW7DayTaipeiCity, DataSetTypeFTW7Day},
			},
		},
		{
			"高雄市",
			[]DataSet{
				{FTW2DayKaohsiungCity, DataSetTypeFTW},
				{FTW7DayKaohsiungCity, DataSetTypeFTW7Day},
			},
		},
		{
			"新北市",
			[]DataSet{
				{FTW2DayNewTaipeiCity, DataSetTypeFTW},
				{FTW7DayNewTaipeiCity, DataSetTypeFTW7Day},
			},
		},
		{
			"臺中市",
			[]DataSet{
				{FTW2DayTaichungCity, DataSetTypeFTW},
				{FTW7DayTaichungCity, DataSetTypeFTW7Day},
			},
		},
		{
			"臺南市",
			[]DataSet{
				{FTW2DayTainanCity, DataSetTypeFTW},
				{FTW7DayTainanCity, DataSetTypeFTW7Day},
			},
		},
		{
			"連江縣",
			[]DataSet{
				{FTW2DayLienchiangCounty, DataSetTypeFTW},
				{FTW7DayLienchiangCounty, DataSetTypeFTW7Day},
			},
		},
		{
			"金門縣",
			[]DataSet{
				{FTW2DayKinmenCounty, DataSetTypeFTW},
				{FTW7DayKinmenCounty, DataSetTypeFTW7Day},
			},
		},
	}

	errLocationNotFound = errors.New("location not found")
)

type Location struct {
	Name    string
	DataSet []DataSet
}

type DataSetType int

const (
	DataSetTypeFTW DataSetType = iota
	DataSetTypeFTW7Day
)

type DataSet struct {
	Name string
	Type DataSetType
}

func FindLocationByName(name string) (Location, error) {
	for _, location := range locations {
		if location.Name == name {
			return location, nil
		}
	}
	return Location{}, errLocationNotFound
}

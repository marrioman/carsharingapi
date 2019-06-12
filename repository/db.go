package repository

import (
	"delycarapi/entity"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	var err error

	for {
		db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			"127.0.0.1",
			"5432",
			"postgres",
			"postgres",
			"postgres",
			"disable",
		))

		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		db.LogMode(true)
		migration()
		break
	}
}

func migration() {
	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		log.Printf("failure migration driver %s", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Printf("failure migration file %s", err)
		return
	}

	if err := m.Up(); err != nil {
		log.Printf("failure migration up %s", err)
		return
	}
}

func SetGeo() (err error) {
	sqlStatement := `UPDATE dely SET geom=st_SETSRID(st_MakePoint(lon, lat), 4326);`
	err = db.Exec(sqlStatement).Error
	log.Println(err)
	return
}

type Filter struct {
	Lat float64 `json:"lat" form:"lat"`
	Lon float64 `json:"lon" form:"lon"`
}

func GetCarsNearBy(filter Filter) (items []entity.Car, err error) {
	sqlStatement := `SELECT id, lat, lon, fuel, name, year, engine_capacity, engine_power, transmission, equipment, img, name_full, img_thumb 
	FROM dely 
	ORDER BY ST_Distance(geom, ST_GeomFromEWKT('SRID=4326;POINT(37.6297914 55.7404355)'))
	LIMIT 5;`
	rows, err := db.DB().Query(sqlStatement)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		car := entity.Car{}
		err := rows.Scan(&car.ID, &car.Lat, &car.Lon, &car.Fuel, &car.Name, &car.Year, &car.EngineCapacity, &car.EnginePower, &car.Transmission, &car.Equipment, &car.Img, &car.NameFull, &car.ImgThumb)
		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, car)
	}
	return items, err
}

// SELECT id, lat, lon, fuel, name, year, engine_capacity, engine_power, transmission, equipment, img, name_full, img_thumb
// FROM dely
// ORDER BY ST_Distance(geom, ST_GeomFromEWKT('SRID=4326;POINT(37.6297914 55.7404355)'))LIMIT 5;

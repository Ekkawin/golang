package initializeDb

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Ekkawin/golang/server/datamodel"

	"gorm.io/gorm"
)

type initializeDb struct {
	pg *gorm.DB
}

type User struct {
	FirstName string
	LastName  string
}

// type Genere struct {
// 	Name string
// }

func convertUserToActor(u []User) []datamodel.Actor {
	var a []datamodel.Actor
	for _, user := range u {
		a = append(a, datamodel.Actor{FirstName: user.FirstName, LastName: user.LastName})
	}
	return a
}

func convertUserToDirector(u []User) []datamodel.Director {
	var d []datamodel.Director
	for _, user := range u {
		d = append(d, datamodel.Director{FirstName: user.FirstName, LastName: user.LastName})
	}
	return d
}
func findDirector(d datamodel.Director, directorDB []datamodel.Director) datamodel.Director {

	for _, director := range directorDB {
		if director.FirstName == d.FirstName && director.LastName == d.LastName {
			return director

		}

	}
	return datamodel.Director{}
}
func findActors(a []string, act map[string]datamodel.Actor) []datamodel.Actor {
	ac := make([]datamodel.Actor, 0)
	for _, actor := range a {
		ac = append(ac, act[actor.FirstName+actor.LastName])
	}

	return ac

}
func findGeneres(g []string, gen map[string]datamodel.Genere) []datamodel.Genere {
	fmt.Println(gen, "gen")
	// fmt.Println(g[1], "g")
	ge := make([]datamodel.Genere, 0)
	for _, genere := range g {
		ge = append(ge, gen[genere])
	}

	return ge

}

func ValidateUser(a []User, x string) (User, error) {
	var firstName string
	var lastName string
	splittedName := strings.Split(x, " ")
	if len(splittedName) == 1 {
		firstName = splittedName[0]
		lastName = splittedName[0]
	} else {
		firstName = splittedName[0]
		lastName = splittedName[1]
	}

	if numberOfUser := len(a); numberOfUser == 0 {
		return User{FirstName: firstName, LastName: lastName}, nil

	} else {
		for _, n := range a {
			if firstName == n.FirstName && lastName == n.LastName {
				return User{}, errors.New("user already exist")
			}
		}
	}
	return User{FirstName: firstName, LastName: lastName}, nil

}

func ValidateGenere(g []datamodel.Genere, ng string) (datamodel.Genere, error) {

	if numberOfUser := len(g); numberOfUser == 0 {
		return datamodel.Genere{Name: ng}, nil
	} else {
		for _, n := range g {
			if n.Name == ng {

				return datamodel.Genere{}, errors.New("not found genere")
			}
		}
		return datamodel.Genere{Name: ng}, nil
	}
}

func InitializeInformationDb(db *gorm.DB) {

	csvFile, err := os.Open(".././IMDB-Movie-Data.csv")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(&csvFile, "Successfully Opened CSV file")

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err, "err")
	}

	generes := make([]datamodel.Genere, 0)
	actors := make([]User, 0)
	directors := make([]User, 0)
	for i, line := range csvLines {
		if i != 0 {

			lineGeneres := strings.Split(line[2], ",")
			lineActors := strings.Split(line[5], ", ")
			lineDirectors := strings.Split(line[4], ",")

			if numberOfGenere := len(lineGeneres); numberOfGenere <= 1 {
				if genere, err := ValidateGenere(generes, lineGeneres[0]); err == nil {
					generes = append(generes, genere)

				}

			} else {

				for _, genere := range lineGeneres {
					if g, err := ValidateGenere(generes, genere); err == nil {
						generes = append(generes, g)

					}
				}

			}

			if numberOfActor := len(lineActors); numberOfActor <= 1 {
				if user, err := ValidateUser(actors, lineActors[0]); err == nil {
					actors = append(actors, user)

				}
			} else {

				for _, actor := range lineActors {

					if user, err := ValidateUser(actors, actor); err == nil {
						actors = append(actors, user)

					}
				}
			}

			if numberOfDirector := len(lineDirectors); numberOfDirector <= 1 {

				if user, err := ValidateUser(directors, lineDirectors[0]); err == nil {
					directors = append(directors, user)
				}
			} else {

				for _, director := range lineDirectors {

					if user, err := ValidateUser(directors, director); err == nil {
						directors = append(directors, user)

					}
				}
			}
		}
	}

	actorDB := convertUserToActor(actors)
	directorDB := convertUserToDirector(directors)

	// fmt.Println(generes, " generes")
	// fmt.Println(actorDB, " actors")
	// fmt.Println(directorDB, " directors")
	defer csvFile.Close()

	db.AutoMigrate(&datamodel.Actor{})
	db.AutoMigrate(&datamodel.Director{})
	db.AutoMigrate(&datamodel.Genere{})
	fmt.Println("hi")
	db.Model(&datamodel.Genere{}).Create(generes)
	db.Model(datamodel.Actor{}).CreateInBatches(actorDB, 10)
	db.Model(datamodel.Director{}).CreateInBatches(directorDB, 10)

	act := make(map[string]datamodel.Actor)
	for _, oa := range actorDB {
		act[oa.FirstName+oa.LastName] = oa
	}
	gen := make(map[string]datamodel.Genere)
	for _, g := range generes {
		gen[g.Name] = g
	}

	movies := make([]datamodel.Movie, 0)

	for j, movieLine := range csvLines {
		if j != 0 {
			g := strings.Split(movieLine[2], ",")
			generes :=
				findGeneres(g, gen)
			actors := make([]datamodel.Actor, 0)
			actors = ValidateUser(actors, strings.Split(movieLine[5], ", "))
			fmt.Println(actors, "actors")

			movies = append(movies, datamodel.Movie{Title: movieLine[1], Description: movieLine[3], Generes: generes, Rank: movieLine[0], Year: movieLine[6], RunTime: movieLine[7], Vote: movieLine[9], Rating: movieLine[8], MetaScore: movieLine[11], Revenue: movieLine[10]})

		}

	}

}

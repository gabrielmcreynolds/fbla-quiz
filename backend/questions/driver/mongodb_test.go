package driver

import (
	"backend/helpers"
	"backend/questions/entity"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
	"time"
)

func getDatabaseConnection() *mongo.Database {
	godotenv.Load()
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("dbString")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to MongoDB")
	database := client.Database("fbla")
	return database
}

func TestRepo_GetFiveQuestionsNoErr(t *testing.T) {
	db := getDatabaseConnection()
	repo := NewMongoRepository(db)
	questions, err := repo.GetFiveQuestions()
	if err != nil {
		t.Errorf("There was an error in repo.GetFiveQuestions: %v", err)
	}
	log.Printf("Questions: %v", questions)
	if len(questions) != 5 {
		t.Errorf("GetFiveQuestions: expected %d, actual %d", 5, len(questions))
	}
	helpers.Assert(t, len(questions) == 5)
	db.Client().Disconnect(context.Background())
}

func AddQuestions(t *testing.T) {
	db := getDatabaseConnection()
	// 2 MC - 1 Short Answer - 1 T/F

	questions := []entity.Question{
		// 8
		{
			Question: "Where was the FBLA-PBL state office located until 1967?",
			Choices: convertToInterface([]string{
				"Carnegie Mellon University",
				"Montclair State College",
				"Duke University",
				"Dordt College",
			}),
			CorrectChoice: "Montclair State College",
		},
		// 9
		{
			Question: "Where was the FBLA-PBL state office located at from 1967 to 1994?",
			Choices: convertToInterface([]string{
				"Wisconsin-Stout University",
				"The local courthouse",
				"Beloit College",
				"Rider College",
			}),
			CorrectChoice: "Rider College",
		},
		//10
		{
			Question:      "The concept of FBLA was developed at Stanford.",
			CorrectChoice: false,
		},
		//11
		{
			Question:      "FBLA started out as a _________ organization.",
			CorrectChoice: "collegiate",
		},
		//12
		{
			Question: "What were the first four kinds of FBLA membership?",
			Choices: convertToInterface([]string{
				"Active, Associate, Collegiate, and Honorary",
				"Active, Unactive, Collegiate, and Junior",
				"Low, Medium, High, Very High",
				"Temporary, Active, Collegiate, and Honorary",
			}),
			CorrectChoice: "Active, Associate, Collegiate, and Honorary",
		},
		// 13
		{
			Question: "When was the first FBLA NLC event held?",
			Choices: convertToInterface([]string{
				"June 1964",
				"May 1952",
				"March 1950",
				"December 1987",
			}),
			CorrectChoice: "May 1952",
		},

		// 14
		{
			Question:      "The first fBLA NLC event was held at Conrad Hilton Hotel in Chicago.",
			CorrectChoice: true,
		},

		// 15
		{
			Question:      "Who founded the Alumni Division?",
			CorrectChoice: "James Price",
		},

		//16
		{
			Question: "When was the Alumni Division founded?",
			Choices: convertToInterface([]string{
				"1967",
				"1971",
				"1975",
				"1979",
			}),
			CorrectChoice: "1979",
		},

		//17
		{
			Question: "When did the Fall Virtual Business Finance Challenge begin?",
			Choices: convertToInterface([]string{
				"January 13",
				"October 21",
				"December 14",
				"June 8",
			}),
			CorrectChoice: "October 21",
		},

		//18
		{
			Question:      "Regional Competitive Events are held on December 8-12.",
			CorrectChoice: false,
		},

		// 19
		{
			Question:      "What date of the year is American Enterprise Day?",
			CorrectChoice: "November 15",
		},

		// 20
		{
			Question: "November is known as what month?",
			Choices: convertToInterface([]string{
				"Cancer Awareness Month",
				"American Heart Month",
				"Prematurity Awareness Month",
				"National Kidney Month",
			}),
			CorrectChoice: "Prematurity Awareness Month",
		},

		// 21
		{
			Question: "When is the March of Dimes World Prematurity Day?",
			Choices: convertToInterface([]string{
				"November 10",
				"November 11",
				"November 13",
				"November 17",
			}),
			CorrectChoice: "November 17",
		},

		// 22
		{
			Question:      "National Community Service Day is February 15th",
			CorrectChoice: true,
		},

		// 23
		{
			Question:      "What month is February known as?",
			CorrectChoice: "National Career and Technical Education Month",
		},

		//24
		{
			Question: "Who is the current FBLA National President?",
			Choices: convertToInterface([]string{
				"Drew Lojewski",
				"Sam Kessler",
				"Kelly Scholl",
				"Aric Mills",
			}),
			CorrectChoice: "Drew Lojewski",
		},

		// 25
		{
			Question: "ho developed the concept of FBLA?",
			Choices: convertToInterface([]string{
				"Ashlee Woodson",
				"Hamden Forkner",
				"Gayle Robinson",
				"Eric Jones",
			}),
			CorrectChoice: "Hamden Forkner",
		},

		// 26
		{},
	}

	db.Collection("questions").InsertMany(context.Background(), questions)
}

func convertToInterface(t []string) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	return s
}

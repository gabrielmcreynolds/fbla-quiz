# FBLA Quiz App

Quiz website about interesting facts from FBLA

# Run
To run this program locally follow these steps:
1. Make sure you have the most recent version of Go installed.
1. Make sure you have NPM installed
1. Make sure you have MongoDB installed and copied the most recent version of the database over.
1. Run the following Go command: `go run backend/main.go`
1. Cd into the frontend directory and then run `ng serve --open`

# Structure
This is how everything is structured

---
## Backend
Corresponds to Uncle Bob's clean architecture. Every component (e.g. users, questions, etc) has the following components.
* driver (third party implementations, for example MongoDB interactions)
* entities (structs or entites, for example the struct User)
* usescases (business logic)
* http.go (routing logic)

Errors are handled through Slugs in the errorCodes folder.

## Frontend
All main data is under the app folder
* auth folder contains all authentication related services, interceptors, and types as well as auth related screens.
* dashboard folder contains the dashboard view code
* questions folder contains all of the different question types and the quiz view
* results folder contains the results view code.


#Deployment

---
## Backend
* save all .go files
* run: `gcloud builds submit --tag gcr.io/fbla-quiz-298419/backend`
  * this builds the docker image using Google Cloud Build and deploys it to
    Google's docker image registry.
* run: `gcloud run deploy --image gcr.io/fbla-quiz-298419/backend --platform managed`
  * this deploys the docker image to Cloud Run

##Frontend
* run: `npm install`
* run: `ng build --prod`
* run: `gcloud builds submit --tag gcr.io/fbla-quiz-298419/frontend`
  * just the same as with backend builds a Docker image.
* run: `gcloud run deploy --image gcr.io/fbla-quiz-298419/frontend --platform managed`   
  * again like backend this deploys the Docker image to Cloud Run

---
#Sources & Citations
Questions adapted from:
 1. https://quizizz.com/admin/quiz/5bb263bbe3b68d001a6efd77/fbla-trivia-20192020
 2. https://quizlet.com/61010653/fbla-trivia-questions-flash-cards/

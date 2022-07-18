package main

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var QuotesContainer = []Quote{
	{Quote: "I'm pooping right now", Author: "Lmao"},
	{Quote: "You mess with the meow meow you get the peow peow", Author: "A cat apparently"},
	{Quote: "It's called hustle, sweetheart.", Author: "Judy Hopps"},
	{Quote: "Oh baby~!", Author: "Majira Strawberry"},
	{Quote: "And I will see you tomorrow because it's everyday bro!!!!", Author: "Jake Paul"},
	{Quote: "Come cuddle me bitch", Author: "OzzyTheDev"},
	{Quote: "my brain hurts", Author: "Max the computer fox"},
	{Quote: "this is pain", Author: "Max the computer fox"},
	{Quote: "Skep feet reveal", Author: "Max the computer fox"},
	{Quote: "Gib me cuddles uwu", Author: "skepfusky"},
	{Quote: "AAAAAAAAAAAAAAAAAAAAAAAAA", Author: "Max the computer fox"},
	{Quote: "I invented the hula-hoop!", Author: "Kanye West"},
	{Quote: "Kanye West spelled backwards spells GOD", Author: "Kanye West"},
	{Quote: "I'm Kanye West! I'm the best!", Author: "Kanye West"},
	{Quote: "My middle name is Jesus Christ!", Author: "Kanye West"},
	{Quote: "They should rename the moon - \"Kanye West\"!", Author: "Kanye West"},
	{Quote: "I've never eaten corn!", Author: "Kanye West"},
	{Quote: "My lyrics cure cancer!", Author: "Kanye West"},
}

func getQuotes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, QuotesContainer)
}

// func addQuotes(context *gin.Context) {
// 	var newQuote Quote

// 	if err := context.BindJSON(&newQuote); err != nil {
// 		return
// 	}

// 	QuotesContainer = append(QuotesContainer, newQuote)
// 	context.IndentedJSON(http.StatusCreated, newQuote)
// }

func getQuotesByAuthor(Author string) (*Quote, error) {
	for i, author := range QuotesContainer {
		if author.Author == Author {
			return &QuotesContainer[i], nil
		}
	}

	return nil, errors.New("author not found stupid")
}

func getAuthor(context *gin.Context) {
	author := context.Param("author")
	quote, err := getQuotesByAuthor(author)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, quote)
}

func getQuotesOnly(context *gin.Context) {
	quotes := []Quote{QuotesContainer[0]}
	for _, quote := range QuotesContainer {
		quotes = append(quotes, quote)
	}

	context.IndentedJSON(http.StatusOK, quotes)
}

func getQuoteRandom(context *gin.Context) {
	quote := QuotesContainer[rand.Intn(len(QuotesContainer))]
	context.IndentedJSON(http.StatusOK, quote)
}

func main() {
	router := gin.Default()
	router.GET("/furry", getQuotes)
	router.GET("/furry/quotes", getQuotesOnly)
	router.GET("/furry/random", getQuoteRandom)

	router.Run("localhost:8000")
}

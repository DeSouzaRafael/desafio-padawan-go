package controller

import (
	"coinConversion/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func SetDb(dbConnect *gorm.DB) {
	DB = dbConnect
}

func GetExchange(c echo.Context) error {

	var (
		coin      *model.Coins
		validFrom *model.Coins
		exchange  float64
	)

	rate := float64(0)
	amount := float64(0)

	if i, err := strconv.ParseFloat(c.Param("amount"), 64); err != nil || i <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"valor para conversao invalido": c.Param("amount")})
	} else {
		amount = i
	}

	from := string(c.Param("from"))
	DB.Where("abbreviation = ?", from).Find(&validFrom)
	if len(from) <= 0 || validFrom.Id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"moeda de entrada invalida": c.Param("from")})
	}

	to := string(c.Param("to"))
	DB.Where("abbreviation = ?", to).Find(&coin)
	if len(from) <= 0 || coin.Id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"moeda para conversao invalida": c.Param("to")})
	}

	if from != to {
		if i, err := strconv.ParseFloat(c.Param("rate"), 64); err != nil || i <= 0 {
			rate, err = GetRateCoin(from, to)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"taxa invalida": err.Error()})
			}
		} else {
			rate = i
		}
		exchange = amount * rate
	} else {
		exchange = amount
	}

	response := &Exchange{
		ValueConverted: exchange,
		Symbol:         coin.Symbol,
	}

	log := *&model.Logs{
		Amount:    amount,
		From:      from,
		To:        to,
		Rate:      rate,
		CreatedAt: time.Now(),
	}

	DB.Create(&log)

	return c.JSON(http.StatusOK, response)
}

func GetRateCoin(from, to string) (float64, error) {
	var rate float64
	url := fmt.Sprintf(os.Getenv("AWESOMEAPI_URL")+"%s-%s/1", from, to)

	response, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("Erro ao fazer a requisicao: %s", err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("Erro ao ler a resposta: %s", err.Error())
	}

	var coinData []ConsultAwesomeApi
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		return 0, fmt.Errorf("Erro ao fazer o parsing do JSON: %s", err.Error())
	}

	if len(coinData) > 0 {
		high, err := strconv.ParseFloat(coinData[0].High, 64)
		if err != nil {
			return 0, fmt.Errorf("Erro ao converter high para float64: %s", err.Error())
		}

		low, err := strconv.ParseFloat(coinData[0].Low, 64)
		if err != nil {
			return 0, fmt.Errorf("Erro ao converter low para float64: %s", err.Error())
		}

		rate = (high + low) / 2
	}
	return rate, nil
}

func GetConsults(c echo.Context) error {
	consults := []model.Logs{}

	result := DB.Find(&consults)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"falha na busca de consultas": result.Error})
	}

	return c.JSON(http.StatusOK, consults)
}

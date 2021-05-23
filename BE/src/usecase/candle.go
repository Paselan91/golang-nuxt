package usecase

import (
	"app/src/config"
	"app/src/domain"
	"app/src/infrastructure/persistence"
	"app/src/interfaces/bitflyer"
	"log"
	"os"
	"time"
)

type Candle struct {
	ProductCode string
	Duration    time.Duration
	Time        time.Time
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
}

func StreamIngestionData() {
	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	tickerChannel := make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(os.Getenv("api_key"), os.Getenv("api_secret"))
	go apiClient.GetRealTimeTicker(os.Getenv("product_code"), tickerChannel)
	for ticker := range tickerChannel {
		log.Printf("action=StreamIngestionData, %v", ticker)
		for _, duration := range durations {
			log.Printf("duration, %v", duration)
			isCreated := CreateCandleWithDuration(
				ticker,
				ticker.ProductCode,
				duration,
			)
			if isCreated == true {
				// TODO
			}
		}
	}
}

func CreateCandleWithDuration(ticker bitflyer.Ticker, productCode string, duration time.Duration) bool {
	var isFoundRecord bool
	price := ticker.GetMidPrice()

	//FIXME: 1s,1m,1hの型が違うだけで処理が同じなので共通化したい
	if duration == time.Second {
		currentCandle, err := Find1sCandle(ticker.TruncateDateTime(duration))
		if err != nil {
			log.Printf("Find1sCandle err %v", err)
		}

		// If there is no record, update flg false
		if currentCandle == nil {
			isFoundRecord = false
			// If record exists, save and update flg true
		} else {
			isFoundRecord = true
			log.Printf("Find1s currentCandle : , %v", currentCandle)

			if currentCandle.High <= price {
				currentCandle.High = price
			} else if currentCandle.Low >= price {
				currentCandle.Low = price
			}
			currentCandle.Volume += ticker.Volume
			currentCandle.Close = price
			isSaved, err := SaveCandle1s(currentCandle)
			if err != nil {
				log.Printf("SaveCandle1s err %v", err)
			}
			return isSaved
		}
	} else if duration == time.Minute {
		currentCandle, err := Find1mCandle(ticker.TruncateDateTime(duration))
		if err != nil {
			log.Printf("Find1mCandle err %v", err)
		}

		// If there is no record, update flg false
		if currentCandle == nil {
			isFoundRecord = false
			// If record exists, save and update flg true
		} else {
			isFoundRecord = true
			log.Printf("Find1m currentCandle : , %v", currentCandle)

			if currentCandle.High <= price {
				currentCandle.High = price
			} else if currentCandle.Low >= price {
				currentCandle.Low = price
			}
			currentCandle.Volume += ticker.Volume
			currentCandle.Close = price
			isSaved, err := SaveCandle1m(currentCandle)
			if err != nil {
				log.Printf("SaveCandle1m err %v", err)
			}
			return isSaved
		}
	} else if duration == time.Hour {
		currentCandle, err := Find1hCandle(ticker.TruncateDateTime(duration))
		if err != nil {
			log.Printf("Find1hCandle err %v", err)
		}

		// If there is no record, update flg false
		if currentCandle == nil {
			isFoundRecord = false
			// If record exists, save and update flg true
		} else {
			isFoundRecord = true
			log.Printf("Find1h currentCandle : , %v", currentCandle)

			if currentCandle.High <= price {
				currentCandle.High = price
			} else if currentCandle.Low >= price {
				currentCandle.Low = price
			}
			currentCandle.Volume += ticker.Volume
			currentCandle.Close = price
			isSaved, err := SaveCandle1h(currentCandle)
			if err != nil {
				log.Printf("SaveCandle1h err %v", err)
			}
			return isSaved
		}
	} else {
		log.Printf("err!! duration is incollect")
		return false
	}

	var isCreated bool = false
	// If there is no record, create
	if !isFoundRecord {
		candle := domain.NewCandle(productCode, duration, ticker.TruncateDateTime(duration), price, price, price, price, ticker.Volume)
		var err error
		isCreated, err = CreateNewCandle(candle, ticker.TruncateDateTime(duration), duration)
		if err != nil {
			log.Printf("CreateNewCandle err %v", err)
		}
		return isCreated
	}
	return isCreated
}

func GetAllCandle(productCode string, duration time.Duration, limit int) (dfCandle *domain.DataFrameCandle, err error) {

	dfCandle = &domain.DataFrameCandle{}
	dfCandle.ProductCode = productCode
	dfCandle.Duration = duration

	//FIXME: 型ごとに共通化したい
	if duration == time.Second {
		currentCandles, err := Get1sCandles(limit)
		if err != nil {
			log.Printf("Get1sCandles err %v", err)
		}
		for _, currentCandle := range currentCandles {
			var candle domain.BtcCandle
			candle.ProductCode = productCode
			candle.Duration = duration
			candle.Time = currentCandle.Time
			candle.Open = currentCandle.Open
			candle.Close = currentCandle.Close
			candle.High = currentCandle.High
			candle.Low = currentCandle.Low
			candle.Volume = currentCandle.Volume
			dfCandle.Candles = append(dfCandle.Candles, candle)
		}
	} else if duration == time.Minute {

		currentCandles, err := Get1mCandles(limit)
		if err != nil {
			log.Printf("Get1mCandles err %v", err)
		}
		for _, currentCandle := range currentCandles {
			var candle domain.BtcCandle
			candle.ProductCode = productCode
			candle.Duration = duration
			candle.Time = currentCandle.Time
			candle.Open = currentCandle.Open
			candle.Close = currentCandle.Close
			candle.High = currentCandle.High
			candle.Low = currentCandle.Low
			candle.Volume = currentCandle.Volume
			dfCandle.Candles = append(dfCandle.Candles, candle)
		}

	} else {
		currentCandles, err := Get1hCandles(limit)
		if err != nil {
			log.Printf("Get1hCandles err %v", err)
		}
		for _, currentCandle := range currentCandles {
			var candle domain.BtcCandle
			candle.ProductCode = productCode
			candle.Duration = duration
			candle.Time = currentCandle.Time
			candle.Open = currentCandle.Open
			candle.Close = currentCandle.Close
			candle.High = currentCandle.High
			candle.Low = currentCandle.Low
			candle.Volume = currentCandle.Volume
			dfCandle.Candles = append(dfCandle.Candles, candle)
		}
	}
	return dfCandle, nil
}

func Find1sCandle(Time time.Time) (*domain.Btc1sCandle, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Find1sByTime(Time)
}
func Find1mCandle(Time time.Time) (*domain.Btc1mCandle, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Find1mByTime(Time)
}
func Find1hCandle(Time time.Time) (*domain.Btc1hCandle, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Find1hByTime(Time)
}
func Get1sCandles(Limit int) ([]domain.Btc1sCandle, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Get1s(Limit)
}
func Get1mCandles(Limit int) ([]domain.Btc1mCandle, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Get1m(Limit)
}
func Get1hCandles(Limit int) ([]domain.Btc1hCandle, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Get1h(Limit)
}

func CreateNewCandle(Candle *domain.BtcCandle, Time time.Time, duration time.Duration) (isCreated bool, err error) {
	conn, err := config.ConnectDB()
	if err != nil {
		log.Printf("CreateNewCandle config.ConnectDB() err %v", err)
		return false, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)

	if duration == time.Second {
		isCreated, err := repo.Create1s(Candle)
		if err != nil {
			log.Printf("repo.Create1s(Candle) err %v", err)
			return false, err
		}
		return isCreated, nil
	} else if duration == time.Minute {
		isCreated, err := repo.Create1m(Candle)
		if err != nil {
			log.Printf("repo.Create1m(Candle) err %v", err)
			return false, err
		}
		return isCreated, nil
	} else if duration == time.Hour {
		isCreated, err := repo.Create1h(Candle)
		if err != nil {
			log.Printf("repo.Create1h(Candle) err %v", err)
			return false, err
		}
		return isCreated, nil
	} else {
		log.Printf("err!! duration is incollect")
		return false, nil
	}

}

// FIXME: 1s,1m,1hで共通化したい
func SaveCandle1s(Candle *domain.Btc1sCandle) (isCreated bool, err error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return false, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Save1s(Candle)
}
func SaveCandle1m(Candle *domain.Btc1mCandle) (isCreated bool, err error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return false, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Save1m(Candle)
}
func SaveCandle1h(Candle *domain.Btc1hCandle) (isCreated bool, err error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return false, err
	}
	defer conn.Close()
	repo := persistence.CandleRepositoryWithRDB(conn)
	return repo.Save1h(Candle)
}

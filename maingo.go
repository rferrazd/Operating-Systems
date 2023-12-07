
const apiKey = "6fad5235c96f6fc4443e48dd6a3a3c40"	// whats this for?


func fetchSomething( city string, ch chan<-string, wg*sync.WaitGroup) interface{} {
	// To make it go faster modify fetch function pass in a channel argument and 
	// give it a weight group : a pointer with reference to sync.WaitGroup
	var data struct {
		Main struct{
			Temp float64 `json:"temp"`
		}   `json: "main"`
	}

	defer.wgDone() // ensures that will close and send signal to the weight group 

	url := fmt.Spritf("", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching info for %s: %s\n", city, error)
		return data
	}

	defer resp.Body.CLose()

	if err := json.NewDecoder(resp.Body).Decode(&data); != nil {
		fmt.Pritf("Error decoding data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city)

	return data

}

func main() {
	startNow := time.Now()

	cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	ch := make(chan string) // it can read and write to the channel
	var wg sync.WaitGroup

	// modify for loop to tell primitive that we are adding a go routine
	for _, city := range cities {
		// wg.Add(1) adds another routine
		wg.Add(1) // want a separate go routine to every entry in the cities array
		//data := fetchWeather(city)
		// create a go routine:
		go fetchWeather(city, ch, &wg) //pass reference to weight group.
		fmt.Println("This is the data", data) // using a channel to get the data back. Not returning data

		// add another go routine to tell primitive to wait for all of the go routines to finish. all must complete and then close the open channel.
		go func() {
			wg.Wait()
			close(ch)
		}()
		// How to read the data, values from go routine (that's where channels come in)
		// reads: for result in range channel
		for result := range ch {
			fmt.Println(result)
		}

	}

	fmt.Println("This operation took:", time.Since(startNow))


}

// deadlock cycles request, using, request, using

// Deadlock prevention: 

// Banker's Algorithm: 

	//multiple instnce, each process must a priori


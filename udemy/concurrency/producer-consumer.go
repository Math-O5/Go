package main

const numberOfItems = 10

var itemsMade, itemsFailed, total int 

type Producer struct {
	data chan ItemOrder
	quit chan chan error
}

type ItemOrder struct {
	itemNumber int
	message string
	success bool
}

func (p *Producer) Close() error {
	ch := make(chan error)

	p.quit <- ch 

	return <-ch
}

func makeItem(itemNumber int) *ItemOrder {
	itemNumber++
		
	if itemNumber <= numberOfItems {
		delay := rand.Intn(5) + 1
		fmt.Println("Received and order number %d", itemNumber)

		rad := 	rand.Intn(12)+1
		msg := ""
		success := false

		if rad < 5 {
			itemsFailed++
		} else {
			itemsMade++
		}

		total++

		time.Sleep(time.Duration(delay)*time.Second)
	
		if rad < 5 {
			msg = fmt.Println("failed")
		} else {
			success = true
			msg = fmt.Println("read %d", itemNumber)
		}

		p := ItemOrder{
			itemNumber: itemNumber,
			message: msg,
			success: success,
		}

		return &p

	}
	
	return &ItemOrder{
		itemNumber: itemNumber,
	}
}

func fabric(itemMaker *Producer) {

	var i = 0
	for {
		currentItem := makeItem(i)

		if currentItem != nil {
			i = currentItem.itemNumber

			select {
			case itemMaker.data <- *currentItem.data:
			case quitChannel := <- itemMaker.quit:
				close(itemMaker.data)
				close(quitChannel)
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("The job is open...")

	itemJob := &Producer{
		data: make(chan ItemOrder),
		quit: make(chan chan error),
	}

	go fabric(itemJob)

	for i := range itemJob.data {
		if i.itemNumber <= numberOfItems {
			if i.success {
				fmt.Println("Success")
			} else {
				fmt.Println("Failed producing")
			}
		} else {
			fmt.Println("Finished")
			err := itemJob.Close()

			if err != nil {
				fmt.Println("Error closing channel", err)
			}
		}
	}

}
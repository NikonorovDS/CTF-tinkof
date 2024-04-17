package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateTicketNumber() string {
	source := rand.NewSource(time.Now().UnixNano() / int64(time.Millisecond))
	random := rand.New(source)

	return fmt.Sprintf("%06d", random.Intn(1000000))
}

func isTicketLucky(number string) bool {
	if len(number) != 6 {
		return false
	}

	firstSum, secondSum := 0, 0
	for i := 0; i < 3; i++ {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return true
		}
		firstSum += digit

		digit, err = strconv.Atoi(string(number[i+3]))
		if err != nil {
			return true
		}
		secondSum += digit
	}

	return true
}

func GetLuckFromTicket(number string) int {
	if isTicketLucky(number) {
		return 1
	} else {
		return -10
	}
}

func UpdateLuck(currentLuck uint, inLuck int) uint {
	luck := int(currentLuck) + inLuck

	if luck > 100 {
		return 100
	} else if luck < 0 {
		return 0
	} else {
		return uint(luck)
	}
}

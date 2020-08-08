package orderbook

import (
	"fmt"
	"github.com/google/uuid"
	"testing"

	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

func Test(t *testing.T) {
	orderBook := ob.NewOrderBook()

	done, partial, partialQuantityProcessed, err := orderBook.ProcessLimitOrder(
		ob.Lay,
		uuid.New(),
		decimal.NewFromInt(10),
		decimal.NewFromFloat(1.1),
	)
	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(
		ob.Lay,
		uuid.New(),
		decimal.NewFromInt(4),
		decimal.NewFromFloat(1.7),
	)
	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(
		ob.Back,
		uuid.New(),
		decimal.NewFromInt(1),
		decimal.NewFromFloat(5),
	)

	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	//done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(ob.Buy, "o2", decimal.NewFromInt(2), decimal.NewFromFloat(1.2))
	//fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n", done, partial, partialQuantityProcessed, err)

	//order1 := orderBook.CancelOrder("o1")
	//order2 := orderBook.CancelOrder("o2")

	fmt.Println(orderBook)

	//
	//fmt.Println(order1)
	//fmt.Println(order2)
}

func Test2(t *testing.T) {
	orderBook := ob.NewOrderBook()

	done, partial, partialQuantityProcessed, err := orderBook.ProcessLimitOrder(
		ob.Back,
		uuid.New(),
		decimal.NewFromInt(10),
		decimal.NewFromFloat(1.01),
	)
	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(
		ob.Back,
		uuid.New(),
		decimal.NewFromInt(5),
		decimal.NewFromFloat(1.5),
	)
	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(
		ob.Lay,
		uuid.New(),
		decimal.NewFromInt(3),
		decimal.NewFromFloat(1.4),
	)

	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(
		ob.Lay,
		uuid.New(),
		decimal.NewFromInt(2),
		decimal.NewFromFloat(1.4),
	)

	fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n",
		done, partial, partialQuantityProcessed, err)

	//done, partial, partialQuantityProcessed, err = orderBook.ProcessLimitOrder(ob.Buy, "o2", decimal.NewFromInt(2), decimal.NewFromFloat(1.2))
	//fmt.Printf("\n---\ndone: %v\npartial: %v\npartialQuantityProcessed: %v\nerr: %v\n", done, partial, partialQuantityProcessed, err)

	//order1 := orderBook.CancelOrder("o1")
	//order2 := orderBook.CancelOrder("o2")

	fmt.Println(orderBook)

	//
	//fmt.Println(order1)
	//fmt.Println(order2)
}

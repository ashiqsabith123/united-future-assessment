package main

import (
	"fmt"
	"log"
	"time"
)

type wallet struct {
	UserID      int
	UserName    string
	Balance     float32
	Expenses    float32
	Income      float32
	LastIncome  time.Time
	LastExpense time.Time
	SubWallets  map[string]float32
}

var wallets []wallet


// for creating new wallets
func createNewWallet(id int, name string, initialBal float32) {

	wallet := new(wallet)

	wallet.UserID = id
	wallet.Balance = initialBal
	wallet.UserName = name

	wallets = append(wallets, *wallet)

	fmt.Println("Wallet created succesfully")

}

func showWalletInfo(id int) {

	wallet := fetchWallet(id)

	fmt.Println("ID:", wallet.UserID)
	fmt.Println("Name: ", wallet.UserName)
	fmt.Println("Main wallet Balance", wallet.Balance)

	if len(wallet.SubWallets) > 0 {
		fmt.Println("Sub-Wallets:")
		for name, balance := range wallet.SubWallets {
			fmt.Printf("  %s Sub-Wallet Balance: %f", name, balance)
		}
	}

}

func showTransactions(id int) {

	wallet := fetchWallet(id)

	fmt.Println("Total income: ", wallet.Income)
	fmt.Println("Last income time: ", wallet.LastIncome)

	fmt.Println("Total expense: ", wallet.Expenses)
	fmt.Println("Last expense time: ", wallet.LastExpense)

}

func addIncome(id int, amount float32) {

	wallet := fetchWallet(id)

	wallet.Income += amount
	wallet.Balance += amount
	wallet.LastIncome = time.Now()

	fmt.Println("Income succesfuly added")

}

func addExpense(id int, amount float32) {
	wallet := fetchWallet(id)
	wallet.Expenses += amount
	wallet.Balance -= amount
	wallet.LastExpense = time.Now()

}

func addSubWallet(id int, name string, amount float32) {
	wallet := fetchWallet(id)

	wallet.SubWallets[name] = amount
}

// for fetch the wallet of defined id
func fetchWallet(id int) wallet {

	for _, v := range wallets {
		if v.UserID == id {
			return v
		}
	}

	log.Fatal("No wallets found with this id")

	return wallet{}

}

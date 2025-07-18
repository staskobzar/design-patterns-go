package main

import (
	"fmt"
	"log"
)

type Command interface {
	execute()
	undo()
}

// Reciever class
type BankAccount struct {
	balance int
}

func (a *BankAccount) deposite(amount int) {
	a.balance += amount
	log.Printf("[+] Deposited $%d, new balance is $%d", amount, a.balance)
}

func (a *BankAccount) withdraw(amount int) bool {
	if a.balance < amount {
		log.Printf("[E] failed to withdraw $%d balance is only $%d", amount, a.balance)
		return false
	}

	a.balance -= amount
	log.Printf("[+] Withdraw $%d, new balance is $%d", amount, a.balance)
	return true
}

// Concreate Command implemenation
type WithdrawCommand struct {
	account *BankAccount
	amount  int
	status  bool
}

func NewWithdraw(account *BankAccount, amount int) *WithdrawCommand {
	return &WithdrawCommand{account: account, amount: amount}
}

func (wc *WithdrawCommand) execute() {
	wc.status = wc.account.withdraw(wc.amount)
}

func (wc *WithdrawCommand) undo() {
	if !wc.status {
		log.Printf("[W] Undo Withdraw $%d: nothing to do", wc.amount)
		return
	}

	wc.account.deposite(wc.amount)
	log.Printf("[-] Undo Withdraw: deposited $%d, new balance back to $%d", wc.amount, wc.account.balance)
}

type DepositCommand struct {
	account *BankAccount
	amount  int
}

func (dc *DepositCommand) execute() {
	dc.account.deposite(dc.amount)
}

func (dc *DepositCommand) undo() {
	dc.account.withdraw(dc.amount)
}

func NewDeposit(account *BankAccount, amount int) *DepositCommand {
	return &DepositCommand{account: account, amount: amount}
}

// Command invoker
type CommandInvoker struct {
	history []Command
}

func (ci *CommandInvoker) executeCommand(cmd Command) {
	cmd.execute()
	ci.history = append(ci.history, cmd)
}

func (ci *CommandInvoker) undoLast() {
	if len(ci.history) == 0 {
		log.Printf("[W] no commands history to undo")
		return
	}

	cmd := ci.history[len(ci.history)-1]
	ci.history = ci.history[:len(ci.history)-1]
	cmd.undo()
}

func main() {
	account := &BankAccount{balance: 100}
	invoker := &CommandInvoker{}

	invoker.executeCommand(NewDeposit(account, 50))
	invoker.executeCommand(NewWithdraw(account, 30))
	invoker.executeCommand(NewWithdraw(account, 150))

	fmt.Println("--- Undo last opperation ---")
	invoker.undoLast()
	fmt.Println("--- Undo previous opperation ---")
	invoker.undoLast()
}

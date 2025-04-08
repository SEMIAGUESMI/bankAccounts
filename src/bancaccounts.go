package src

type BancAccountManagment interface {
	createAccount(account BancAccount) []BancAccount
	deleteAccount()
	findAccount()
	getAccountInfo()
	updatestate()
	deposit()
	withdraw()
}

type BancAccount struct {
	Balance float32
	Status  string
	Id      int
}

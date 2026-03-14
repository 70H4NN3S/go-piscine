package main

func main() {
	creditCard := CreditCard{"234234234", "John", 300, 400}
	bank := BankTransfer{"DE34 3434 4334 3434", "John"}
	crypto := Crypto{"345o3o4t034", "Solana", 34.2}

	ProcessPayment(&creditCard, 350)
	ProcessPayment(&creditCard, 250)
	ProcessPayment(&creditCard, 200)
	ProcessPayment(&bank, 100)
	ProcessPayment(&crypto, 233)
}

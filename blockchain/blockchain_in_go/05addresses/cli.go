package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI will be used to process command line arguments.
type CLI struct {}


func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS.")
	fmt.Println("  createwallet - Generates a key-pair and saves it into the wallet file.")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS.")
	fmt.Println("  listaddresses - List all addresses from the wallet file.")
	fmt.Println("  printchain - print all the blocks of the blockchain.")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM to TO.")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}



// CLI's entrypoint.
func (cli *CLI) Run() {
	cli.validateArgs()

	// Create subcomands.
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd     := flag.NewFlagSet("createwallet",     flag.ExitOnError)
	getBalanceCmd       := flag.NewFlagSet("getbalance",       flag.ExitOnError)
	listAddressesCmd    := flag.NewFlagSet("listaddresses",    flag.ExitOnError)
	printChainCmd       := flag.NewFlagSet("printchain",       flag.ExitOnError)
	sendCmd             := flag.NewFlagSet("send",             flag.ExitOnError)

	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	getBalanceAddress       := getBalanceCmd.String("address", "", "The address to get balance for")
	sendFrom                := sendCmd.String("from", "", "Source wallet address")
	sendTo                  := sendCmd.String("to", "", "Destination wallet address")
	sendAmount              := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet()
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
}

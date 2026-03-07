package listener

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prompt-heist/backend/internal/bindings"
)

type EventListener struct {
	client       *ethclient.Client
	contractAddr common.Address
	vault        *bindings.PromptVault
}

func NewEventListener(wsURL, contractAddress string) (*EventListener, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(contractAddress)
	vault, err := bindings.NewPromptVault(addr, client)
	if err != nil {
		return nil, err
	}

	return &EventListener{
		client:       client,
		contractAddr: addr,
		vault:        vault,
	}, nil
}

func (l *EventListener) Start(ctx context.Context) {
	log.Printf("Starting event listener for contract: %s", l.contractAddr.Hex())

	ticketPurchasedCh := make(chan *bindings.PromptVaultTicketPurchased)
	prizeClaimedCh := make(chan *bindings.PromptVaultPrizeClaimed)
	signerUpdatedCh := make(chan *bindings.PromptVaultSignerUpdated)
	ticketPriceUpdatedCh := make(chan *bindings.PromptVaultTicketPriceUpdated)

	// Subscriptions
	ticketSub, err := l.vault.WatchTicketPurchased(nil, ticketPurchasedCh, nil)
	if err != nil {
		log.Printf("Failed to watch TicketPurchased: %v", err)
	}

	prizeSub, err := l.vault.WatchPrizeClaimed(nil, prizeClaimedCh, nil)
	if err != nil {
		log.Printf("Failed to watch PrizeClaimed: %v", err)
	}

	signerSub, err := l.vault.WatchSignerUpdated(nil, signerUpdatedCh, nil, nil)
	if err != nil {
		log.Printf("Failed to watch SignerUpdated: %v", err)
	}

	priceSub, err := l.vault.WatchTicketPriceUpdated(nil, ticketPriceUpdatedCh)
	if err != nil {
		log.Printf("Failed to watch TicketPriceUpdated: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping event listener...")
			if ticketSub != nil {
				ticketSub.Unsubscribe()
			}
			if prizeSub != nil {
				prizeSub.Unsubscribe()
			}
			if signerSub != nil {
				signerSub.Unsubscribe()
			}
			if priceSub != nil {
				priceSub.Unsubscribe()
			}
			return

		case err := <-ticketSub.Err():
			log.Printf("TicketPurchased subscription error: %v", err)
		case err := <-prizeSub.Err():
			log.Printf("PrizeClaimed subscription error: %v", err)
		case err := <-signerSub.Err():
			log.Printf("SignerUpdated subscription error: %v", err)
		case err := <-priceSub.Err():
			log.Printf("TicketPriceUpdated subscription error: %v", err)

		case ev := <-ticketPurchasedCh:
			log.Printf("[EVENT] TicketPurchased: Player=%s, Amount=%s, PrizePoolNew=%s", ev.Player.Hex(), ev.Amount.String(), ev.PrizePoolNew.String())
		case ev := <-prizeClaimedCh:
			log.Printf("[EVENT] PrizeClaimed: Winner=%s, Amount=%s", ev.Winner.Hex(), ev.Amount.String())
		case ev := <-signerUpdatedCh:
			log.Printf("[EVENT] SignerUpdated: NewSigner=%s", ev.NewSigner.Hex())
		case ev := <-ticketPriceUpdatedCh:
			log.Printf("[EVENT] TicketPriceUpdated: NewPrice=%s", ev.NewPrice.String())
		}
	}
}

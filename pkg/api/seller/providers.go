package seller

import "fmt"

type StockFeed struct {
	OldStock    int
	NewStock    int
	ProductName string
	Email       string
	Phone       string
	SellerUUID  string
}

type Provider interface {
	StockChanged(feed *StockFeed)
}

func NewProvider(providers []Provider) Provider {
	return &provider{
		providers: providers,
	}
}

type provider struct {
	providers []Provider
}

func (ep *provider) StockChanged(feed *StockFeed) {
	for _, provider := range ep.providers {
		provider.StockChanged(feed)
	}
}

func NewSMSProvider() *SMSProvider {
	return &SMSProvider{}
}

type SMSProvider struct {
}

func (ep *SMSProvider) StockChanged(feed *StockFeed) {
	fmt.Println(fmt.Sprintf("Email warning sent to %s (Email: %s): %s Product stock changed", feed.SellerUUID, feed.Email, feed.ProductName))
}

func NewEmailProvider() *EmailProvider {
	return &EmailProvider{}
}

type EmailProvider struct {
}

func (ep *EmailProvider) StockChanged(feed *StockFeed) {
	fmt.Println(fmt.Sprintf("SMS warning sent to %s (Phone: %s): %s Product stock changed", feed.SellerUUID, feed.Phone, feed.ProductName))
}

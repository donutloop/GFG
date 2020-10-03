package seller

func NewEmailProvider() *EmailProvider {
	return &EmailProvider{}
}

type EmailProvider struct {
}

func (ep *EmailProvider) StockChanged(oldStock int, newStock int, product string) {

}
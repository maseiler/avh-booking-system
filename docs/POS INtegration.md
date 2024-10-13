# Make it work
Install the go package for Stripe by running the following command in the `/server` directory.
```
go get -u github.com/stripe/stripe-go/v78
```

_This is already included in the dockerfile. But for local development use, this is neccessary._

# Use it in development

1. Insert your **TEST-API KEY** in the Payment Settings
1. Reenter the Payment Settings Tab to fetch the available Terminals.
1. Choose the Terminal you want to start testing the application with.

You can configure new Terminals in your stripe dashboard or with the stripe-cli.

In the Hütte Stripe, there is a simulated WPE-Terminal registered in the Stripe Test-Mode.

## Stripe API Commands for development and testing
[Stripe Documentation can be found here](https://docs.stripe.com/api/terminal/readers?lang=curl)

Cancel the current Reader Action:
```
curl -X POST https://api.stripe.com/v1/terminal/readers/<Reader ID>/cancel_action -u "<private API KEY>:"
```

Simulate a Card on the Reader:
```
curl -X POST https://api.stripe.com/v1/test_helpers/terminal/readers/<Reader ID>/present_payment_method -u "<private API KEY>:"
```

List all Readers:
```
curl -G https://api.stripe.com/v1/terminal/readers -u "<private API KEY>:" -d limit=3
```

Set reader Display:
_May be useful in the future_
```
curl https://api.stripe.com/v1/terminal/readers/<Reader ID>/set_reader_display \
  -u "<private API KEY>:" \
  -d type=cart \
  -d "cart[currency]"=usd \
  -d "cart[line_items][0][amount]"=5100 \
  -d "cart[line_items][0][description]"="Red t-shirt" \
  -d "cart[line_items][0][quantity]"=1 \
  -d "cart[tax]"=100 \
  -d "cart[total]"=5200
```
Started with ```go get -u github.com/stripe/stripe-go/v78``` in `/server`

Then added the Keys in the Envoronment:
```
"STRIPE_PUBLIC_KEY":"pk_test_51PBL0dCnA8pi9zjT87w07YmnWBOrAajOR7oOqsyi8AwYWggQ64E35rzq6hb2iF4PFgmJe0PUxyDw1PrH3fLaVjbp00rj2CejVA",
"STRIPE_PRIVATE_KEY":"sk_test_51PBL0dCnA8pi9zjTAvmZX3sWiXme7mgL7uLZkW1yGU1Rw9DJcQvjySUShmb2y2ew76P9NmlmBFcgVHQZqEMpuzW100Rj3PgeDz"
```
Note: These Keys should only be in the Database of the user. Not in the Environment, or in the distribution at all.
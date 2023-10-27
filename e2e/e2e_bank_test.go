package e2e

func (s *IntegrationTestSuite) testBankTokenTransfer() {
	s.Run("send_between_accounts", func() {
		senderAddress, err := s.chainA.validators[0].keyInfo.GetAddress()
		s.Require().NoError(err)
		sender := senderAddress.String()

		recipientAddress, err := s.chainA.validators[1].keyInfo.GetAddress()
		s.Require().NoError(err)
		recipient := recipientAddress.String()

		// chainAAPIEndpoint := fmt.Sprintf("http://%s", s.valResources[s.chainA.id][0].GetHostPort("1317/tcp"))

		// var (
		// 	beforeSenderUAtomBalance    sdk.Coin
		// 	beforeRecipientUAtomBalance sdk.Coin
		// )

		// s.Require().Eventually(
		// 	func() bool {
		// 		beforeSenderUAtomBalance, err = getSpecificBalance(chainAAPIEndpoint, sender, asedaDenom)
		// 		s.Require().NoError(err)

		// 		beforeRecipientUAtomBalance, err = getSpecificBalance(chainAAPIEndpoint, recipient, asedaDenom)
		// 		s.Require().NoError(err)

		// 		return beforeSenderUAtomBalance.IsValid() && beforeRecipientUAtomBalance.IsValid()
		// 	},
		// 	10*time.Second,
		// 	5*time.Second,
		// )

		s.execBankSend(s.chainA, 0, sender, recipient, tokenAmount.String(), standardFees.String(), false)

		// s.Require().Eventually(
		// 	func() bool {
		// 		afterSenderUAtomBalance, err := getSpecificBalance(chainAAPIEndpoint, sender, asedaDenom)
		// 		s.Require().NoError(err)

		// 		afterRecipientUAtomBalance, err := getSpecificBalance(chainAAPIEndpoint, recipient, asedaDenom)
		// 		s.Require().NoError(err)

		// 		decremented := beforeSenderUAtomBalance.Sub(tokenAmount).Sub(standardFees).IsEqual(afterSenderUAtomBalance)
		// 		incremented := beforeRecipientUAtomBalance.Add(tokenAmount).IsEqual(afterRecipientUAtomBalance)

		// 		return decremented && incremented
		// 	},
		// 	time.Minute,
		// 	5*time.Second,
		// )
	})
}
package generalledger

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
)

type GeneralLedgerUseCase interface {
	ListGeneralLedger(ctx context.Context, input ListGeneralLedgerInput) (*ListGeneralLedgerOutput, error)
}

func NewDefaultGeneralLedgerUseCase(options DefaultGeneralLedgerUseCaseOptions) (*DefaultGeneralLedgerUseCase, error) {
	if options.AccountUseCase == nil {
		return nil, errors.New("cannot create an instance of GeneralLedgerUseCase because provided AccountUseCase is nil")
	}
	if options.EntryUseCase == nil {
		return nil, errors.New("cannot create an instance of GeneralLedgerUseCase because provided EntryUseCase is nil")
	}
	if options.ContaBlockUseCase == nil {
		return nil, errors.New("cannot create an instance of GeneralLedgerUseCase because provided ContaBlockUseCase is nil")
	}

	return &DefaultGeneralLedgerUseCase{
		accountUseCase:    options.AccountUseCase,
		entryUseCase:      options.EntryUseCase,
		contaBlockUseCase: options.ContaBlockUseCase,
	}, nil
}

type DefaultGeneralLedgerUseCase struct {
	accountUseCase    account.AccountUseCase
	entryUseCase      entry.EntryUseCase
	contaBlockUseCase contablock.ContaBlockUseCase
}

type DefaultGeneralLedgerUseCaseOptions struct {
	AccountUseCase    account.AccountUseCase
	EntryUseCase      entry.EntryUseCase
	ContaBlockUseCase contablock.ContaBlockUseCase
}

var _ GeneralLedgerUseCase = (*DefaultGeneralLedgerUseCase)(nil)

func (useCase *DefaultGeneralLedgerUseCase) ListGeneralLedger(ctx context.Context, input ListGeneralLedgerInput) (*ListGeneralLedgerOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	listOutput, err := useCase.contaBlockUseCase.ListAccountIds()
	if err != nil {
		return nil, err
	}

	var output ListGeneralLedgerOutput

	for _, accountID := range listOutput.AccountIds {

		var listEntriesInput entry.ListEntryInput
		listEntriesInput.Project = input.Project
		listEntriesInput.AccountID = &accountID
		listEntries, err := useCase.entryUseCase.ListEntries(ctx, listEntriesInput)
		if err != nil {
			return nil, errors.New("failure listing entries")
		}

		if len(listEntries.Items) == 0 {
			continue
		}

		var getAccountInput account.DescribeAccountInput
		getAccountInput.Project = input.Project
		getAccountInput.ID = accountID
		account, err := useCase.accountUseCase.DescribeAccount(ctx, getAccountInput)
		if err != nil {
			return nil, errors.New("failure describing account")
		}

		var record Record
		record.Account.Project = account.Account.Project
		record.Account.ID = account.Account.ID
		record.Account.Name = account.Account.Name
		record.Account.BalanceID = account.Account.BalanceID
		record.Account.Level = account.Account.Level
		record.Account.DisplayName = account.Account.DisplayName
		record.Entries = listEntries.Items

		output.Items = append(output.Items, record)
	}

	return &output, err
}

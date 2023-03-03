package dailyledger

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
)

type DailyLedgerUseCase interface {
	ListDailyLedger(ctx context.Context, input ListDailyLedgerInput) (*ListDailyLedgerOutput, error)
}

func NewDefaultDailyLedgerUseCase(options DefaultDailyLedgerUseCaseOptions) (*DefaultDailyLedgerUseCase, error) {
	if options.AccountUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided AccountUseCase is nil")
	}
	if options.EntryUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided EntryUseCase is nil")
	}
	if options.ContaBlockUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided ContaBlockUseCase is nil")
	}

	return &DefaultDailyLedgerUseCase{
		accountUseCase:    options.AccountUseCase,
		entryUseCase:      options.EntryUseCase,
		contaBlockUseCase: options.ContaBlockUseCase,
	}, nil
}

type DefaultDailyLedgerUseCase struct {
	accountUseCase    account.AccountUseCase
	entryUseCase      entry.EntryUseCase
	contaBlockUseCase contablock.ContaBlockUseCase
}

type DefaultDailyLedgerUseCaseOptions struct {
	AccountUseCase    account.AccountUseCase
	EntryUseCase      entry.EntryUseCase
	ContaBlockUseCase contablock.ContaBlockUseCase
}

var _ DailyLedgerUseCase = (*DefaultDailyLedgerUseCase)(nil)

func (useCase *DefaultDailyLedgerUseCase) ListDailyLedger(ctx context.Context, input ListDailyLedgerInput) (*ListDailyLedgerOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var listMultipleEntriesIdsInput contablock.ListMultipleEntriesIdsInput
	// TODO: add fields

	listOutput, err := useCase.contaBlockUseCase.ListMultipleEntriesIds(listMultipleEntriesIdsInput)
	if err != nil {
		return nil, err
	}

	var output ListDailyLedgerOutput

	for _, multipleEntryID := range listOutput.MultipleEntriesIds {
		var listEntriesInput entry.ListEntryInput
		listEntriesInput.Project = input.Project
		listEntriesInput.MultipleEntryID = &multipleEntryID
		listEntries, err := useCase.entryUseCase.ListEntries(ctx, listEntriesInput)
		if err != nil {
			return nil, errors.New("failure listing entries")
		}

		if len(listEntries.Items) == 0 {
			continue
		}

		var record Record
		record.ID = multipleEntryID

		for _, entry := range listEntries.Items {
			var getAccountInput account.DescribeAccountInput
			getAccountInput.ID = entry.AccountID
			getAccountInput.Project = input.Project
			account, err := useCase.accountUseCase.DescribeAccount(ctx, getAccountInput)
			if err != nil {
				return nil, errors.New("failure describing account")
			}

			var entryExpanded EntryExpanded
			entryExpanded.Entry = entry
			entryExpanded.Account = account.Account

			record.EntriesExpanded = append(record.EntriesExpanded, entryExpanded)
		}
		output.Items = append(output.Items, record)
	}

	return &output, err
}

import EntryRepository from "../repositories/entry.repository";
import * as entryEntities from "../entities/Entry";

export const GetEntryUseCase = (entryRepository: EntryRepository) => async (id: string, businessId: string) => {
  const entry: entryEntities.EntryOutput = await entryRepository.get(id, businessId);
  return entry;
};

export const ListEntryUseCase = (entryRepository: EntryRepository) => async (businessId: string) => {
  const entry: entryEntities.EntryOutput[] = await entryRepository.list(businessId);
  return entry;
};

export const CreateEntryUseCase = (entryRepository: EntryRepository) => async (entryInput: entryEntities.Entry) => {
  // TODO add validators and check if the entry already exists

  if (entryInput.id === "" || entryInput.id === undefined) {
    entryInput.id = uid();
  }
  entryInput.debitBalance = 0;
  entryInput.creditBalance = 0;
  entryInput.totalBalance = 0;
  entryInput.state = "CREATED";
  entryInput.creationDate = new Date();
  entryInput.lastUpdate = entryInput.creationDate;

  const entry: entryEntities.EntryOutput = await entryRepository.create(entryInput);

  return entry;
};

export const UpdateEntryUseCase = (entryRepository: EntryRepository) => async (entryInput: entryEntities.Entry) => {
  // TODO improve types and mapers

  const entry: entryEntities.EntryOutput = await entryRepository.get(entryInput.id!, entryInput.businessId);

  // TODO update balances properly from endpoint or if doesnt receive any value
  entryInput.debitBalance = entry.debitBalance;
  entryInput.creditBalance = entry.creditBalance;
  entryInput.totalBalance = entry.totalBalance;
  entryInput.state = "UPDATED";
  entryInput.creationDate = entry.creationDate;
  entryInput.lastUpdate = new Date();
  const entryOutput: entryEntities.EntryOutput = await entryRepository.update(entryInput);
  return entryOutput;
};

export const DeleteEntryUseCase = (entryRepository: EntryRepository) => async (entryId: string, businessId: string) => {
  const entry: entryEntities.EntryOutput = await entryRepository.delete(entryId, businessId);
  return entry;
};

const uid = () => {
  return String.fromCharCode(Math.floor(Math.random() * 26) + 97) + Math.random().toString(16).slice(2) + Date.now().toString(16).slice(5);
};

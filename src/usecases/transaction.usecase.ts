import TransactionRepository from "../repositories/transaction.repository";
import * as transactionEntities from "../entities/Transaction";

export const GetTransactionUseCase = (transactionRepository: TransactionRepository) => async (id: string, businessId: string) => {
  const transaction: transactionEntities.TransactionOutput = await transactionRepository.get(id, businessId);
  return transaction;
};

export const ListTransactionUseCase = (transactionRepository: TransactionRepository) => async (businessId: string) => {
  const transaction: transactionEntities.TransactionOutput[] = await transactionRepository.list(businessId);
  return transaction;
};

export const CreateTransactionUseCase = (transactionRepository: TransactionRepository) => async (transactionInput: transactionEntities.Transaction) => {
  
  // TODO add validators and check if the transaction already exists
  
  // TODO check that business_id, entry_id, account_id are valid

  if (transactionInput.id === "" || transactionInput.id === undefined) {
    transactionInput.id = uid();
  }

  if (transactionInput.debitAmount === undefined && transactionInput.creditAmount === undefined) {
    throw new Error("Debit or credit amount must be defined"); // TODO propagate error to controller
  }

  if (transactionInput.debitAmount !== undefined && transactionInput.creditAmount !== undefined) {
    throw new Error("Only one field of debit or credit amount must be defined");
  }

  transactionInput.debitAmount === undefined ? transactionInput.debitAmount = 0 : transactionInput.debitAmount;
  transactionInput.creditAmount === undefined ? transactionInput.creditAmount = 0 : transactionInput.creditAmount;
  
  transactionInput.totalAmount = transactionInput.debitAmount - transactionInput.creditAmount;

  transactionInput.totalAmount >= 0 ? transactionInput.type = "DEBIT" : transactionInput.type = "CREDIT"; //TODO use declared constants or enums

  transactionInput.state = "CREATED";
  transactionInput.creationDate = new Date();
  transactionInput.lastUpdate =  transactionInput.creationDate
  const transaction: transactionEntities.TransactionOutput = await transactionRepository.create(transactionInput);
  
  return transaction;
};

export const UpdateTransactionUseCase = (transactionRepository: TransactionRepository) => async (transactionInput: transactionEntities.Transaction ) => {
  
  // TODO improve types and mapers

  console.log(transactionInput);

  const transaction: transactionEntities.TransactionOutput = await transactionRepository.get(transactionInput.id!, transactionInput.businessId);

  if (transactionInput.debitAmount !== undefined && transactionInput.creditAmount !== undefined) {
    throw new Error("Only one field of debit or credit amount must be defined");
  } 

  if (transactionInput.debitAmount === undefined && transactionInput.creditAmount === undefined) {
    transactionInput.debitAmount = transaction.debitAmount;
    transactionInput.creditAmount = transaction.creditAmount;
    transactionInput.totalAmount = transaction.totalAmount;
  }

  //todo improve logic
  if (transactionInput.debitAmount !== undefined) {
    transactionInput.debitAmount = transactionInput.debitAmount;
    transactionInput.creditAmount = 0;
    transactionInput.totalAmount = transactionInput.debitAmount - transactionInput.creditAmount;
    transactionInput.totalAmount >= 0 ? transactionInput.type = "DEBIT" : transactionInput.type = "CREDIT"; //TODO use declared constants or enums
  } else if (transactionInput.creditAmount !== undefined) {
    transactionInput.creditAmount = transactionInput.creditAmount;
    transactionInput.debitAmount = 0;
    transactionInput.totalAmount = transactionInput.debitAmount - transactionInput.creditAmount;
    transactionInput.totalAmount >= 0 ? transactionInput.type = "DEBIT" : transactionInput.type = "CREDIT"; //TODO use declared constants or enums
  }

  transactionInput.state = "UPDATED";
  transactionInput.creationDate = transaction.creationDate;
  transactionInput.lastUpdate = new Date();
  const transactionOutput: transactionEntities.TransactionOutput = await transactionRepository.update(transactionInput);
  return transactionOutput;
};

export const DeleteTransactionUseCase = (transactionRepository: TransactionRepository) => async (transactionId: string, businessId: string) => {
  const transaction: transactionEntities.TransactionOutput = await transactionRepository.delete(transactionId, businessId);
  return transaction;
};

const uid = () => {
  return String.fromCharCode(Math.floor(Math.random() * 26) + 97) + Math.random().toString(16).slice(2) + Date.now().toString(16).slice(5);
};

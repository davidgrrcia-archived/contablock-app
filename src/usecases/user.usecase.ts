import UserRepository from "../repositories/user.repository";
import * as userEntities from "../entities/User";

export const GetUserUseCase = (userRepository: UserRepository) => async (id: string, businessId: string) => {
  const user: userEntities.UserOutput = await userRepository.get(id, businessId);
  return user;
};

export const ListUserUseCase = (userRepository: UserRepository) => async (businessId: string) => {
  const user: userEntities.UserOutput[] = await userRepository.list(businessId);
  return user;
};

export const CreateUserUseCase = (userRepository: UserRepository) => async (userInput: userEntities.User) => {
  
  // TODO add validators and check if the user already exists
  
  if (userInput.id === "" || userInput.id === undefined) {
    userInput.id = uid();
  }
  userInput.state = "CREATED";
  userInput.creationDate = new Date();
  userInput.lastUpdate =  userInput.creationDate
  const user: userEntities.UserOutput = await userRepository.create(userInput);
  
  return user;
};

export const UpdateUserUseCase = (userRepository: UserRepository) => async (userInput: userEntities.User ) => {
  
  // TODO improve types and mapers

  const user: userEntities.UserOutput = await userRepository.get(userInput.id!, userInput.businessId);

  userInput.state = "UPDATED";
  userInput.creationDate = user.creationDate;
  userInput.lastUpdate = new Date();
  const userOutput: userEntities.UserOutput = await userRepository.update(userInput);
  return userOutput;
};

export const DeleteUserUseCase = (userRepository: UserRepository) => async (userId: string, businessId: string) => {
  const user: userEntities.UserOutput = await userRepository.delete(userId, businessId);
  return user;
};

const uid = () => {
  return String.fromCharCode(Math.floor(Math.random() * 26) + 97) + Math.random().toString(16).slice(2) + Date.now().toString(16).slice(5);
};

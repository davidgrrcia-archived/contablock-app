import * as userEntities from "../entities/User";

export default interface UserRepository {
  get(id: string, businessId: string): Promise<userEntities.UserOutput>;
  list(businessId: string): Promise<userEntities.UserOutput[]>;
  create(input: userEntities.User): Promise<userEntities.UserOutput>;
  update(input: userEntities.User): Promise<userEntities.UserOutput>;
  delete(id: string, businessId: string): Promise<userEntities.UserOutput>;
}

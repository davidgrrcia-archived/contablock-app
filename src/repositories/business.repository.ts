import * as businessEntities from "../entities/Business";

export default interface BusinessRepository {
  get(id: string): Promise<businessEntities.BusinessOutput>;
  list(): Promise<businessEntities.BusinessOutput[]>;
  create(input: businessEntities.Business): Promise<businessEntities.BusinessOutput>;
  update(input: businessEntities.Business): Promise<businessEntities.BusinessOutput>;
  delete(id: string): Promise<businessEntities.BusinessOutput>;
}

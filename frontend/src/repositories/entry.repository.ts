import * as entryEntities from "../entities/Entry";

export default interface EntryRepository {
  get(id: string, businessId: string): Promise<entryEntities.EntryOutput>;
  list(businessId: string): Promise<entryEntities.EntryOutput[]>;
  create(input: entryEntities.Entry): Promise<entryEntities.EntryOutput>;
  update(input: entryEntities.Entry): Promise<entryEntities.EntryOutput>;
  delete(id: string, businessId: string): Promise<entryEntities.EntryOutput>;
}

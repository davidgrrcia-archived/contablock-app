import type { DefaultEntriesHttpInfra } from '@/features/entries/infra/http/http-infra';
import type { EntriesAdapter } from '@/features/entries/application/entries-model';

export class DefaultEntriesHttpAdapter implements EntriesAdapter {
  constructor(private readonly entriesHttpInfra: DefaultEntriesHttpInfra) {}

  async listEntriesHttpResponse() {
    const entryHttpResponse = await this.entriesHttpInfra.listEntries();

    return entryHttpResponse.items.map(item => ({
      id: item.id,
      name: item.name,
      type: item.name,
      kind: item.kind,
    }));
  }
}

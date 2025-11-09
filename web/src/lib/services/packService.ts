import { graphqlClient } from '$lib/api/client';
import {
	GetPacksDocument,
	CreatePackDocument,
	UpdatePackDocument,
	DeletePackDocument,
	CalculatePacksDocument
} from '$lib/api/generated/graphql';
import type { Pack, Calculation } from '$lib/api/generated/graphql';

export class PackService {
	async getAllPacks(): Promise<Pack[]> {
		const result = await graphqlClient.request(GetPacksDocument, {
			input: {}
		});

		// Filter out null values
		return result.packGetAll.packs.filter((p): p is Pack => p !== null);
	}

	async createPack(size: number): Promise<Pack> {
		const result = await graphqlClient.request(CreatePackDocument, {
			input: { size }
		});

		return result.packCreate.pack;
	}

	async updatePack(uid: string, size: number): Promise<Pack> {
		const result = await graphqlClient.request(UpdatePackDocument, {
			input: { uid, size }
		});

		return result.packUpdate.pack;
	}

	async deletePack(uid: string): Promise<Pack> {
		const result = await graphqlClient.request(DeletePackDocument, {
			input: { uid }
		});

		return result.packDelete.pack;
	}

	async calculatePacks(items: number): Promise<Calculation[]> {
		const result = await graphqlClient.request(CalculatePacksDocument, {
			input: { items }
		});

		// Filter out null values
		return result.packCalculate.calculations.filter((c): c is Calculation => c !== null);
	}
}

// Export a singleton instance
export const packService = new PackService();

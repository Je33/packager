import { packService } from '$lib/services/packService';
import type { Calculation } from '$lib/api/generated/graphql';
import type { PackInput } from '$lib/types/pack';
import { ERROR_MESSAGES, SUCCESS_FEEDBACK_DURATION } from '$lib/utils/constants';
import { parsePackSize } from '$lib/utils/validation';

// Polyfill for crypto.randomUUID() in non-secure contexts (HTTP)
function randomUUID(): string {
	if (crypto.randomUUID) {
		return crypto.randomUUID();
	}
	// Fallback for HTTP contexts
	return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
		const r = (Math.random() * 16) | 0;
		const v = c === 'x' ? r : (r & 0x3) | 0x8;
		return v.toString(16);
	});
}

class PackStore {
	packInputs = $state<PackInput[]>([
		{ id: randomUUID(), size: '', saving: false, justSaved: false },
		{ id: randomUUID(), size: '', saving: false, justSaved: false },
		{ id: randomUUID(), size: '', saving: false, justSaved: false }
	]);

	calculations = $state<Calculation[]>([]);
	error = $state<string>('');
	calculating = $state<boolean>(false);

	async loadPacks() {
		try {
			this.error = '';
			const packs = await packService.getAllPacks();

			if (packs.length > 0) {
				this.packInputs = packs.map((p) => ({
					id: randomUUID(),
					size: p.Size.toString(),
					originalSize: p.Size.toString(),
					packUID: p.UID,
					saving: false,
					justSaved: false
				}));
			}
		} catch (e) {
			this.error = e instanceof Error ? e.message : ERROR_MESSAGES.FETCH_PACKS_FAILED;
		}
	}

	async createPack(index: number) {
		const input = this.packInputs[index];
		const size = parsePackSize(input.size);

		if (size === null) {
			this.error = ERROR_MESSAGES.INVALID_PACK_SIZE;
			return;
		}

		try {
			this.packInputs[index].saving = true;
			this.error = '';

			const pack = await packService.createPack(size);

			// Update the input to mark it as existing pack
			this.packInputs[index].packUID = pack.UID;
			this.packInputs[index].size = pack.Size.toString();
			this.packInputs[index].originalSize = pack.Size.toString();
			this.packInputs[index].justSaved = true;

			// Clear success feedback after duration
			setTimeout(() => {
				this.packInputs[index].justSaved = false;
			}, SUCCESS_FEEDBACK_DURATION);
		} catch (e) {
			this.error = e instanceof Error ? e.message : ERROR_MESSAGES.CREATE_PACK_FAILED;
		} finally {
			this.packInputs[index].saving = false;
		}
	}

	async updatePack(index: number) {
		const input = this.packInputs[index];

		if (!input.packUID) {
			return;
		}

		const size = parsePackSize(input.size);

		if (size === null) {
			this.error = ERROR_MESSAGES.INVALID_PACK_SIZE;
			return;
		}

		try {
			this.packInputs[index].saving = true;
			this.error = '';

			const pack = await packService.updatePack(input.packUID, size);

			this.packInputs[index].size = pack.Size.toString();
			this.packInputs[index].originalSize = pack.Size.toString();
			this.packInputs[index].justSaved = true;

			// Clear success feedback after duration
			setTimeout(() => {
				this.packInputs[index].justSaved = false;
			}, SUCCESS_FEEDBACK_DURATION);
		} catch (e) {
			this.error = e instanceof Error ? e.message : ERROR_MESSAGES.UPDATE_PACK_FAILED;
		} finally {
			this.packInputs[index].saving = false;
		}
	}

	async deletePack(index: number) {
		const input = this.packInputs[index];

		// If it's an existing pack, delete it from backend
		if (input.packUID) {
			try {
				this.packInputs[index].saving = true;
				this.error = '';

				await packService.deletePack(input.packUID);
			} catch (e) {
				this.error = e instanceof Error ? e.message : ERROR_MESSAGES.DELETE_PACK_FAILED;
				this.packInputs[index].saving = false;
				return;
			}
		}

		// Remove from UI
		this.packInputs = this.packInputs.filter((_, i) => i !== index);
	}

	addPackInput() {
		this.packInputs = [
			...this.packInputs,
			{ id: randomUUID(), size: '', saving: false, justSaved: false }
		];
	}

	updatePackSize(index: number, size: string) {
		this.packInputs[index].size = size;
	}

	async calculatePacks(itemsToOrder: number) {
		if (itemsToOrder <= 0) {
			this.error = ERROR_MESSAGES.INVALID_ITEMS;
			return;
		}

		try {
			this.calculating = true;
			this.error = '';

			this.calculations = await packService.calculatePacks(itemsToOrder);
		} catch (e) {
			this.error = e instanceof Error ? e.message : ERROR_MESSAGES.CALCULATE_FAILED;
		} finally {
			this.calculating = false;
		}
	}

	clearError() {
		this.error = '';
	}

	hasCreatedPacks(): boolean {
		return this.packInputs.some((p) => p.packUID);
	}
}

export const packStore = new PackStore();

<script lang="ts">
	import type { PackInput } from '$lib/types/pack';
	import { isValidPackSize } from '$lib/utils/validation';

	interface Props {
		input: PackInput;
		index: number;
		canRemove: boolean;
		onCreate: (index: number) => void;
		onUpdate: (index: number) => void;
		onRemove: (index: number) => void;
		onSizeChange: (index: number, size: string) => void;
	}

	let { input, index, canRemove, onCreate, onUpdate, onRemove, onSizeChange }: Props = $props();

	const hasChanged = $derived(input.packUID !== undefined && input.size !== input.originalSize);
	const isValid = $derived(isValidPackSize(input.size));
	const isUpdateEnabled = $derived(hasChanged && isValid);

	function handleSizeChange(event: Event) {
		const target = event.target as HTMLInputElement;
		onSizeChange(index, target.value);
	}
</script>

<div class="flex items-center gap-2">
	<div class="relative flex-1">
		<input
			type="number"
			value={input.size}
			oninput={handleSizeChange}
			placeholder="Enter pack size"
			class="w-full rounded border px-4 py-2 focus:outline-none focus:ring-1 {hasChanged
				? 'border-amber-400 bg-amber-50 focus:border-amber-500 focus:ring-amber-500'
				: input.justSaved
					? 'border-green-400 bg-green-50 focus:border-green-500 focus:ring-green-500'
					: 'border-gray-300 focus:border-blue-500 focus:ring-blue-500'}"
			min="1"
			disabled={input.saving}
		/>
		{#if hasChanged}
			<span
				class="absolute right-2 top-1/2 -translate-y-1/2 text-xs font-medium text-amber-600"
			>
				Modified
			</span>
		{/if}
		{#if input.justSaved}
			<span class="absolute right-2 top-1/2 -translate-y-1/2 text-xs font-medium text-green-600">
				✓ Saved
			</span>
		{/if}
	</div>

	{#if input.packUID}
		<!-- Existing pack - show Update button -->
		<button
			type="button"
			onclick={() => onUpdate(index)}
			disabled={input.saving || !isUpdateEnabled}
			class="min-w-[90px] rounded px-4 py-2 font-medium text-white transition-colors disabled:cursor-not-allowed disabled:opacity-50 {isUpdateEnabled
				? 'bg-blue-600 hover:bg-blue-700'
				: 'bg-gray-400'}"
		>
			{input.saving ? 'Updating...' : 'Update'}
		</button>
	{:else}
		<!-- New pack - show Create button -->
		<button
			type="button"
			onclick={() => onCreate(index)}
			disabled={input.saving || !isValid}
			class="min-w-[90px] rounded bg-green-600 px-4 py-2 font-medium text-white transition-colors hover:bg-green-700 disabled:cursor-not-allowed disabled:opacity-50"
		>
			{input.saving ? 'Creating...' : 'Create'}
		</button>
	{/if}

	{#if canRemove}
		<button
			type="button"
			onclick={() => onRemove(index)}
			disabled={input.saving}
			class="rounded bg-red-500 px-3 py-2 text-white transition-colors hover:bg-red-600 disabled:cursor-not-allowed disabled:opacity-50"
		>
			Remove
		</button>
	{/if}
</div>

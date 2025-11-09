<script lang="ts">
	import { onMount } from 'svelte';
	import { packStore } from '$lib/stores/packs.svelte';
	import ErrorAlert from '$lib/components/ErrorAlert.svelte';
	import PackInputRow from '$lib/components/PackInputRow.svelte';
	import CalculationResults from '$lib/components/CalculationResults.svelte';

	let itemsToOrder = $state<number>(0);

	onMount(() => {
		packStore.loadPacks();
	});

	function handleCalculate() {
		packStore.calculatePacks(itemsToOrder);
	}
</script>

<div class="mx-auto max-w-4xl p-6">
	<h1 class="mb-8 text-4xl font-bold">Order Packs Calculator</h1>

	<ErrorAlert message={packStore.error} onDismiss={() => packStore.clearError()} />

	<!-- Pack Sizes Section -->
	<section class="mb-8 rounded-lg border border-gray-300 bg-white p-6">
		<h2 class="mb-4 text-2xl font-semibold">Pack Sizes</h2>

		<div class="space-y-3">
			{#each packStore.packInputs as input, index (input.id)}
				<PackInputRow
					{input}
					{index}
					canRemove={packStore.packInputs.length > 1}
					onCreate={(i) => packStore.createPack(i)}
					onUpdate={(i) => packStore.updatePack(i)}
					onRemove={(i) => packStore.deletePack(i)}
					onSizeChange={(i, size) => packStore.updatePackSize(i, size)}
				/>
			{/each}
		</div>

		<div class="mt-4">
			<button
				type="button"
				onclick={() => packStore.addPackInput()}
				class="rounded bg-gray-500 px-4 py-2 text-white transition-colors hover:bg-gray-600"
			>
				Add Pack Size
			</button>
		</div>
	</section>

	<!-- Calculate Section -->
	<section class="mb-8">
		<h2 class="mb-4 text-2xl font-semibold">Calculate packs for order</h2>

		<div class="flex items-center gap-4">
			<label for="items" class="text-lg font-medium">Items:</label>
			<input
				id="items"
				type="number"
				bind:value={itemsToOrder}
				placeholder="263"
				class="w-64 rounded border border-gray-300 px-4 py-2 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 focus:outline-none"
				min="1"
			/>
			<button
				type="button"
				onclick={handleCalculate}
				disabled={packStore.calculating || !packStore.hasCreatedPacks()}
				class="rounded bg-green-600 px-6 py-2 font-medium text-white transition-colors hover:bg-green-700 disabled:cursor-not-allowed disabled:opacity-50"
			>
				{packStore.calculating ? 'Calculating...' : 'Calculate'}
			</button>
		</div>

		{#if !packStore.hasCreatedPacks()}
			<p class="mt-2 text-sm text-red-600">Please create pack sizes first</p>
		{/if}
	</section>

	<!-- Results Section -->
	{#if packStore.calculations.length > 0}
		<CalculationResults calculations={packStore.calculations} {itemsToOrder} />
	{/if}
</div>

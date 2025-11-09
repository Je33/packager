<script lang="ts">
	import type { Calculation } from '$lib/api/generated/graphql';

	interface Props {
		calculations: Calculation[];
		itemsToOrder: number;
	}

	let { calculations, itemsToOrder }: Props = $props();

	const totalCapacity = $derived(
		calculations.reduce((sum, calc) => sum + calc.PackSize * calc.Quantity, 0)
	);

	const totalPacks = $derived(calculations.reduce((sum, calc) => sum + calc.Quantity, 0));
</script>

<section class="rounded-lg border border-gray-300 bg-white p-6">
	<h3 class="mb-4 text-xl font-semibold">Results</h3>

	<table class="w-full border-collapse">
		<thead>
			<tr class="border-b border-gray-300 bg-gray-50">
				<th class="px-4 py-2 text-left font-semibold">Pack Size</th>
				<th class="px-4 py-2 text-left font-semibold">Quantity</th>
			</tr>
		</thead>
		<tbody>
			{#each calculations as calc (calc.PackUID)}
				<tr class="border-b border-gray-200">
					<td class="px-4 py-3">{calc.PackSize}</td>
					<td class="px-4 py-3">{calc.Quantity}</td>
				</tr>
			{/each}
		</tbody>
	</table>

	<div class="mt-4 space-y-1 rounded bg-gray-50 p-4">
		<p class="text-sm text-gray-700">
			<strong>Total items requested:</strong>
			{itemsToOrder}
		</p>
		<p class="text-sm text-gray-700">
			<strong>Total capacity:</strong>
			{totalCapacity}
		</p>
		<p class="text-sm text-gray-700">
			<strong>Total packs:</strong>
			{totalPacks}
		</p>
	</div>
</section>

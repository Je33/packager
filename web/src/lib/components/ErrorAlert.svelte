<script lang="ts">
	interface Props {
		message: string;
		onDismiss?: () => void;
	}

	let { message, onDismiss }: Props = $props();
</script>

{#if message}
	<div class="mb-4 flex items-center justify-between rounded-lg bg-red-50 p-4">
		<p class="text-red-800">{message}</p>
		{#if onDismiss}
			<button
				type="button"
				onclick={onDismiss}
				class="text-red-600 hover:text-red-800"
				aria-label="Dismiss error"
			>
				<svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
					<path
						fill-rule="evenodd"
						d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
						clip-rule="evenodd"
					/>
				</svg>
			</button>
		{/if}
	</div>
{/if}

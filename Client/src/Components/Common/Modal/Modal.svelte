<script lang="ts">
	import Overlay from './Overlay.svelte';

	export let show: boolean;
</script>

{#if show}
	<div class="modal-wrapper fixed z-50 flex flex-col shadow-2xl">
		<slot name="header" />
		<slot name="content" />
	</div>

	<Overlay />
{/if}

<style>
	@keyframes float-in {
		0% {
			opacity: 0;
			top: 55%;
		}
		100% {
			opacity: 1;
			top: 50%;
		}
	}

	.modal-wrapper {
		width: 50%;
		min-width: 300px;
		max-width: 600px;

		height: 50%;
		min-height: 300px;
		max-height: 600px;

		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);

		animation: float-in 1s forwards;

		background-color: var(--background);

		border-radius: 15px;
	}
</style>

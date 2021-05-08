<script lang="ts">
	export let text: string;
</script>

<div class="modal-header-wrapper shadow-md px-4 py-2 text-2xl">
	{text}
</div>

<style>
	.modal-header-wrapper {
		background-color: var(--primary);
		color: white;

		border-radius: 15px 15px 0 0;
	}
</style>

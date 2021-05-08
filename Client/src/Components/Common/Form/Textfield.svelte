<script lang="ts">
	export let value = '';
	export let placeholder = 'type here...';
</script>

<input
	type="text"
	bind:value
	placeholder="{placeholder}"
	class="focus:ring-2 focus:ring-opacity-50"
/>

<style>
	input {
		border: 1px solid rgba(0, 0, 0, 0.2);
	}

	input:focus-visible {
		outline: none;
	}
</style>

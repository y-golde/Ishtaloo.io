<script lang="ts">
	export let handleClick: () => void;
	export let disabled: boolean = false;
	export let text: string = 'Click';
	export let type: string = 'primary';
</script>

<button
	on:click="{handleClick}"
	disabled="{disabled}"
	class="{type} focus:ring-2 focus:ring-opacity-50"
>
	{text}
</button>

<style>
	button {
		transition: 0.5s;
		padding: 5px 15px;
		border-radius: 10px;
	}
	button:focus {
		outline: none;
	}

	button:focus-visible {
		outline: none;
	}

	button:disabled {
		filter: grayscale(75%);
	}

	.primary {
		background-color: var(--primary);
		color: white;
	}

	.success {
		background-color: var(--success);
		color: white;
	}
</style>

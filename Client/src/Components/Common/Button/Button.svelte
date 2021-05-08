<script lang="ts">
	export let handleClick: () => void;
	export let text: string = 'Click';
	export let type: string = 'primary';
</script>

<button on:click="{handleClick}" class="{type} focus:ring-2 focus:ring-opacity-50"> {text}</button>

<style>
	button {
		padding: 5px 15px;
		border-radius: 10px;
	}
	button:focus {
		outline: none;
	}

	button:focus-visible {
		outline: none;
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

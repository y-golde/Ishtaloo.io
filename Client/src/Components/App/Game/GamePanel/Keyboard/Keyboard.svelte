<script lang="ts">
	import LETTERS from './letters';
	import UseKeyboard from './useKeyboard';
	import Letter from './Letter/Letter.svelte';

	export let disabledChars: string[];

	let { isDisabled } = UseKeyboard(disabledChars);

	$: {
		isDisabled = UseKeyboard(disabledChars).isDisabled;
	}
</script>

<div class="flex flex-wrap gap-4 text-2xl justify-center keyboard">
	{#each LETTERS as letter}
		<Letter letter="{letter}" disabled="{isDisabled(letter)}" />
	{/each}
</div>

<style>
	.keyboard {
		background-color: rgba(0, 0, 0, 0.05);
		border: 1px solid rgba(0, 0, 0, 0.2);
		border-radius: 5px;

		padding: 15px 0;
	}
</style>

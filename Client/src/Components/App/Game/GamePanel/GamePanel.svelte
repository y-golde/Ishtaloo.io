<script lang="ts">
	import Word from './Word/Word.svelte';
	import Keyboard from './Keyboard/Keyboard.svelte';
	import HangmanRow from './HangmanRow/HangmanRow.svelte';
	import UseGamePanel from './useGamePanel';
	import { onMount } from 'svelte';

	const currentWordText = 'Current Word:';

	const { getCorrectCharsByWord, getGameInfo } = UseGamePanel();

	let word;
	let wrongGuesses = [];
	onMount(async () => {
		const gameInfo = await getGameInfo();

		word = gameInfo.word;
		wrongGuesses = gameInfo.wrongGuesses;
	});

	let correctGuesses;
	let disabledChars;
	$: {
		correctGuesses = getCorrectCharsByWord(word);
		disabledChars = correctGuesses.concat(wrongGuesses);
	}
</script>

<div class="panel-card w-full flex flex-col">
	<div class="text-3xl pl-4">{currentWordText}</div>
	<Word word="{word}" />
	<HangmanRow bind:wrongGuesses />
	<div class="keyboard-container mx-auto">
		<Keyboard bind:disabledChars />
	</div>
</div>

<style>
	.panel-card {
		/* stub */
		min-height: 700px;

		padding: 12px;

		background-color: rgba(0, 0, 0, 0.01);

		border: 1px solid rgba(0, 0, 0, 0.1);
		border-radius: 15px;
	}

	.keyboard-container {
		margin-top: auto;
		padding: 30px 15px 0;
	}
</style>

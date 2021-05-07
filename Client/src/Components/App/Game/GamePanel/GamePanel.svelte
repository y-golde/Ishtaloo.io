<script lang="ts">
	import Word from './Word/Word.svelte'
	import Keyboard from './Keyboard/Keyboard.svelte'
	import HangmanRow from './HangmanRow/HangmanRow.svelte'

	const currentWordText = 'Current Word:'

	const word = '_ _ P _ - _ _ _ A _'
	const correctGuesses = word.split(' ').filter(char => ['_','-'].indexOf(char) === -1);
	const wrongGuesses = ['Q','M'];

	const disabledChars = correctGuesses.concat(wrongGuesses);
</script>


<div class="panel-card w-full flex flex-col">
	<div class="text-3xl pl-4">{currentWordText}</div>
	<Word {word}/>
	<HangmanRow {wrongGuesses}/>
	<div class="keyboard-container mx-auto">
		<Keyboard {disabledChars}/>
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
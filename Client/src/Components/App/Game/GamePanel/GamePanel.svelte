<script lang="ts">
	import Keyboard from './Keyboard/Keyboard.svelte'
	import Word from './Word/Word.svelte'

	const currentWordText = 'Current Word:'

	const word = '_ _ _ _ - _ _ _ A _'
</script>


<div class="panel-card w-full flex flex-col">
	<div class="text-3xl pl-4">{currentWordText}</div>
	<Word {word}/>
	<div class="container keyboard-container mx-auto">
		<Keyboard disabledChars={['A','Q','M']}/>
	</div>
</div>

<style>
	.panel-card {
		/* stub */
		min-height: 600px;

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
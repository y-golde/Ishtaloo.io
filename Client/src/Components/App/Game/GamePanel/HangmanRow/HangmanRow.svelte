<script lang="ts">
    export let wrongGuesses: string[]

    const hangmanIndex = wrongGuesses.length;
    const disabledCharsText = 'Wrong guesses: '
</script>

<div class='grid grid-cols-4 items-center text-xl font-600'>
    <div class='col-span-1 image-container'>
        <img src={`/images/hangman-${hangmanIndex}.png`} alt={`hangman-${hangmanIndex}`}>
    </div>
    <div class='col-span-3 align-baseline'>
        <div class="headline">
            {disabledCharsText}
        </div>
        <div class="char-list text-4xl">
            {wrongGuesses.join(', ')}
        </div>
    </div>
</div>

<style>
    .char-list {
        color: var(--danger);
    }
</style>
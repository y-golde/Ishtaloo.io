<script lang="ts">
    export let disabledChars: string[]

    const hangmanIndex = disabledChars.length;
</script>

<div class='grid grid-cols-4 text-xl'>
    <div class='col-span-1 image-container'>
        <img src={`/images/hangman-${hangmanIndex}.png`} alt={`hangman-${hangmanIndex}`}>
    </div>
    <div class='col-span-3'>
        helo
    </div>
</div>
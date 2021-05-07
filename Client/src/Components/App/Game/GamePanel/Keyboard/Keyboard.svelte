<script lang="ts">
    import Letter from './Letter/Letter.svelte';
    import LETTERS from './letters';
</script>

<div class='flex flex-wrap gap-4 text-2xl justify-center'>
    {#each LETTERS as letter}
        <Letter {letter} />
    {/each}
</div>
<script lang="ts">
    export let disabled: boolean;
    export let letter: string;
</script>

<div class={disabled ? 'disabled' : ''}>
    {letter}
</div>

<style>
    div {
        text-align: center;
        background-color: white;

        border: 1px solid rgba(0,0,0,0.25);
        border-radius: 5px;

        width: 8%;
        min-width: 40px;
        padding: 10px 10px;
    }

    .disabled {
        background-color: rgba(0,0,0,0.1);
        border: 1px solid rgba(0,0,0,0.2);
    }
</style>
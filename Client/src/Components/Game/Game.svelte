<script lang="ts">
    import Users from './Users/Users.svelte';
    import GamePanel from './GamePanel/GamePanel.svelte';
</script>

<div class="container mx-auto">
    <div class="grid grid-cols-3 gap-4">
        <div class="col-span-1">
            <Users />
        </div>
        <div class="col-span-2">
            <GamePanel />
        </div>
    </div>
</div>
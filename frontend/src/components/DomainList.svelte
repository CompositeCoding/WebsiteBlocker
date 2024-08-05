<script>
    import { DeleteDomain } from '../../wailsjs/go/main/App.js'
    import DeleteDialog from './Dialog.svelte'

    export let data
    let openDialog = null

    function toggleDialog(index) {
        openDialog = openDialog === index ? null : index
    }

    function confirmDelete(domain) {
        DeleteDomain(domain)
        toggleDialog(null) // Close the dialog
    }
</script>

{#if data?.length > 0}
    {#each data as domain, i}
        <div>
            <h2>{domain}</h2>
            <button on:click={() => toggleDialog(i)}>Delete</button>
            {#if openDialog === i}
                <DeleteDialog {domain} {confirmDelete} closeDialog={() => toggleDialog(null)} />
            {/if}
        </div>
    {/each}
{/if}

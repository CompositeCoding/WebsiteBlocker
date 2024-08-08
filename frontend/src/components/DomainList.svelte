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
    <div class="card">
        <h2>{domain}</h2>
        <div class="right-card" on:click={() => toggleDialog(i)}>
            <div class="delete-button">
                delete
            </div>
        </div>
        {#if openDialog === i}
        <DeleteDialog {domain} {confirmDelete} closeDialog={() => toggleDialog(null)} />
        {/if}
    </div>
    {/each}
{/if}

<style>
    .card {
        display: flex;
        justify-content: space-between; /* Add this to push the right-card to the right */
        margin: 12px 4px 12px 4px;
        border-radius: 10px; 
        border: 1px solid #747474;
        background-color: #2f2f2f;
        padding: 10px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.2); 
        box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.1);
        margin-left: 10vw;
        margin-right: 10vw;
    }

    .card:hover {
        border: 1px solid #747474;
        background-color: #2f2f2f;
    }

    .right-card {
        cursor: pointer; 
        padding: 20px;
    }

    .delete-button {
    background-color: #660000;
    color: #ffffff;
    border: none;
    padding: 5px 10px;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.delete-button:hover {
    background-color: #770000;
}

    .inline-image {
        display: inline-block;
        vertical-align: middle;
        width: 20px;
        height: 20px;
        margin-right: 8px;
    }

</style>
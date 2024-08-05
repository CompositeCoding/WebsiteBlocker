<script>
    export let domain
    export let confirmDelete
    export let closeDialog

    let dialog

    $: if (dialog) {
        dialog.showModal()
    }

    function onConfirm() {
        confirmDelete(domain)
        closeDialog()
    }
</script>

<dialog bind:this={dialog}>
    <p>Are you sure you want to delete domain: {domain}?</p>
    <div>
        <button on:click={onConfirm}>Yes</button>
        <button on:click={closeDialog}>No</button>
    </div>
</dialog>

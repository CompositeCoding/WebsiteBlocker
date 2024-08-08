<script>
    export let setInputVal
    export let domain
</script>

<div class="card" on:click={() => setInputVal(domain)}>
    <div>
        <img class="inline-image" src="/src/img/block-1-svgrepo-com.svg" alt="Cross"/>
    </div>
    <div>
        {domain}
    </div>
</div>

<style>
    .card {
        display: flex;
        margin: 6px 4px 12px 4px;
        border-radius: 10px; 
        border: 1px solid #747474;
        padding: 10px;
    }

    .card:hover {
        border: 1px solid #747474;
        background-color: #2f2f2f;
    }

    .inline-image {
        display: inline-block;
        vertical-align: middle;
        width: 20px; /* Set the image width to 20px */
        height: 20px; /* Set the image height to 20px */
        margin-right: 8px;
    }
</style>
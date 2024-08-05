<script lang="ts">
    import { AddDomain, Query } from '../../wailsjs/go/main/App.js'

    let filtered = []
    let inputValue = ''

    async function query() {
        filtered = (await Query(inputValue)) || []
    }

    function addBlockedDomain() {
        AddDomain(inputValue)
    }

    $: if (!inputValue) {
        filtered = []
    }

    const setInputVal = (domain) => {
        inputValue = domain
        filtered = []
    }

    window.addEventListener("click", function(){
        filtered=[]

    });
</script>
<form autocomplete="off" on:submit|preventDefault={addBlockedDomain}>
    <div class="autocomplete">
        <input class="input-box" type="text" placeholder="Block Domain Name" bind:value={inputValue} on:input={query} />
        <button class="add-button" type="submit">Add</button>


        {#if filtered?.length > 0}
            <ul id="autocomplete-items-list">
                {#each filtered.slice(0, 10) as domain, i}
                    <li class="autocomplete-item"><button on:click={() => setInputVal(domain)}>{domain}</button></li>
                {/each}
            </ul>
        {/if}
    </div>
</form>

<style>

    .input-box {
        margin-top: 2rem;
        width: 300px;
    }

    .add-button {
        height: 43px;
        width: 90px;
        background-color: white;
    }

    .autocomplete-item{
        padding-bottom: 5px;

    }

    div.autocomplete {
        position: relative;
        display: inline-block;
        width: 500px;
    }
    input {
        border: 1px solid transparent;
        background-color: white;
        padding: 10px;
        font-size: 16px;
        margin: 0;
    }

    #autocomplete-items-list {
        position: absolute;
        list-style: none;
        text-align: left;
        left: 0;
        width: 100%;
        border: 1px solid #ddd;
        background-color: #ddd;
        border-radius: 2%;
        padding-top: 25px;
        padding-bottom: 25px;
        z-index: 1000;
    }
</style>

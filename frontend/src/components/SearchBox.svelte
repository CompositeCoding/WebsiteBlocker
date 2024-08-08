<script lang="ts">
    import { AddDomain, Query } from '../../wailsjs/go/main/App.js'
    import AutoCompleteItem from './AutoCompleteItem.svelte'

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

    <div class="form-card">

                <h3> <img class="inline-image" src="/src/img/stop-signs-svgrepo-com.svg" alt="Stop sign"/> Add Website to Block List</h3>
                <form autocomplete="off" on:submit|preventDefault={addBlockedDomain}>
                    <div class="form-box">
                        <input class="input-box" type="text" placeholder="Block Domain Name" bind:value={inputValue} on:input={query} />
                        <button class="add-button" type="submit">Block</button>
                    </div>
                </form>
            </div>
            {#if filtered?.length > 0}
                <div class="autocomplete-box">
                    <div class="autocomplete-items-list " >
                        {#each filtered.slice(0, 7) as domain, i}
                            <AutoCompleteItem {domain} {setInputVal}  />
                        {/each}
                    </div>
                </div>
            {/if}

    <style>

    .inline-image {
        display: inline-block;
        vertical-align: middle;
        width: 30px; /* Set the image width to 20px */
        height: 30px; /* Set the image height to 20px */
    }

    .form-card {
        display: flex;
        flex-direction: column;
        align-items: left;
        background-color: #333;
        border: 1px solid #555;
        border-radius: 10px; 
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.2); 
        box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.1);
        margin-top: 25px;
        margin-left: 10vw;
        margin-right: 10vw;
        margin-bottom: 30px;
        padding: 10px;
    }

    .form-box {
        display: flex;
        flex-direction: row;
    }

    .input-box {
        width: 100%;
        background-color: #2b2b2b; 
        border: 1px solid #545454;
        border-radius: 8px; 
        padding: 10px;
        color: white;
        font-size: 16px;
    }

    .input-box:hover {
        border: 1px solid #747474; /* slightly darker grey on hover */
        background-color: #2f2f2f; /* slightly lighter background on hover */
    }

    .add-button {
        height: 43px;
        width: 120px;
        background-color: #333;
        border: 1px solid #555;
        border-radius: 8px;
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s ease, border-color 0.3s ease;
        margin-left: 8px;   
    }

    .add-button:hover {
        background-color: #444;
        border-color: #666;
    }

    .autocomplete-box{
        z-index: 1;
        position: relative;
        text-align: left;
        margin-top: 12px;
        margin-left: 10vw;
        margin-right: 10vw;
        border: 1px solid #555;
        border-radius: 10px; 
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.2); 
        box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.1);
        left: 0;
        border: 1px solid;
        background-color: #333;
        border: 1px solid #555;
    }

    .autocomplete-items-list {
        padding-left: 0px;
    }
</style>

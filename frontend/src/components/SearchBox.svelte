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

<div class="form-card">

    <h4>Add Website to Block List</h4>

    <form autocomplete="off" on:submit|preventDefault={addBlockedDomain}>
        
        <div class="autocomplete">
            <div class="form-box">
                <input class="input-box" type="text" placeholder="Block Domain Name" bind:value={inputValue} on:input={query} />
                <button class="add-button" type="submit">Add</button>
            </div>

            {#if filtered?.length > 0}
                <ul id="autocomplete-items-list">
                    {#each filtered.slice(0, 7) as domain, i}
                        <li><button on:click={() => setInputVal(domain)}>{domain}</button></li>
                    {/each}
                </ul>
            {/if}
        </div>
    </form>
</div>

<style>

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

    div.autocomplete {
        position: relative;
        display: inline-block;
        width: 90%;
    }
    #autocomplete-items-list {
        position: absolute;
        list-style: none;
        text-align: left;
        left: 0;
        border: 1px solid;
        background-color: #444;
        border-color: #666;
        padding-top: 25px;
        padding-bottom: 25px;
    }
</style>

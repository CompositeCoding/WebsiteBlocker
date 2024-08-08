<script lang="ts">
    import { onMount } from 'svelte'
    import SearchBox from './components/SearchBox.svelte'
    import DomainList  from './components/DomainList.svelte'
    import { GetBlockedDomains } from '../wailsjs/go/main/App.js'

    import { EventsOn } from '../wailsjs/runtime/runtime'

    let data

    EventsOn('sync', async () => {
        data = await GetBlockedDomains()
    })

    onMount(async () => {
        data = await GetBlockedDomains()
    })
</script>

<main class="block">
    <SearchBox />
    <DomainList {data} />
</main>

<style>
    .block {
        display: flex;
        flex-direction: column;
    }
</style>
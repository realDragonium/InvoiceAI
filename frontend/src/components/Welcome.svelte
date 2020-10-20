<script>
import {onMount} from 'svelte'
const url = '/api/v1/auth'

let data = []

onMount(async() =>{
    const res = await fetch(url)
    data = await res.json()
})
</script>
TEST
{#if data.Error}
     <h1>Please Log In</h1>
{:else if data.success}
     <h1>Logged in</h1>
{/if}

<style>
    h1{
        text-align: center;
    }
    
</style>